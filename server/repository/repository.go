package repository

import (
	"fmt"
	"path"
	"path/filepath"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appsv1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/reposerver"
	"github.com/argoproj/argo-cd/reposerver/repository"
	"github.com/argoproj/argo-cd/util"
	"github.com/argoproj/argo-cd/util/cache"
	"github.com/argoproj/argo-cd/util/db"
	"github.com/argoproj/argo-cd/util/git"
	"github.com/argoproj/argo-cd/util/grpc"
	"github.com/argoproj/argo-cd/util/rbac"
	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
)

// Server provides a Repository service
type Server struct {
	db            db.ArgoDB
	repoClientset reposerver.Clientset
	enf           *rbac.Enforcer
	cache         cache.Cache
}

const (
	DefaultRepoStatusCacheExpiration = 1 * time.Hour
)

// NewServer returns a new instance of the Repository service
func NewServer(
	repoClientset reposerver.Clientset,
	db db.ArgoDB,
	enf *rbac.Enforcer,
	cache cache.Cache,
) *Server {
	return &Server{
		db:            db,
		repoClientset: repoClientset,
		enf:           enf,
		cache:         cache,
	}
}

func (s *Server) getConnectionState(ctx context.Context, url string) appsv1.ConnectionState {
	cacheKey := fmt.Sprintf("connection-state-%s", url)
	var connectionState appsv1.ConnectionState
	if err := s.cache.Get(cacheKey, &connectionState); err == nil {
		return connectionState
	}
	now := v1.Now()
	connectionState = appsv1.ConnectionState{
		Status:     appsv1.ConnectionStatusSuccessful,
		ModifiedAt: &now,
	}
	repo, err := s.db.GetRepository(ctx, url)
	if err == nil {
		err = git.TestRepo(repo.Repo, repo.Username, repo.Password, repo.SSHPrivateKey)
	}
	if err != nil {
		connectionState.Status = appsv1.ConnectionStatusFailed
		connectionState.Message = fmt.Sprintf("Unable to connect to repository: %v", err)
	}
	err = s.cache.Set(&cache.Item{
		Object:     &connectionState,
		Key:        cacheKey,
		Expiration: DefaultRepoStatusCacheExpiration,
	})
	if err != nil {
		log.Warnf("getConnectionState cache set error %s: %v", cacheKey, err)
	}
	return connectionState
}

// List returns list of repositories
func (s *Server) List(ctx context.Context, q *RepoQuery) (*appsv1.RepositoryList, error) {
	urls, err := s.db.ListRepoURLs(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]appsv1.Repository, 0)
	if urls != nil {
		for _, url := range urls {
			if s.enf.EnforceClaims(ctx.Value("claims"), "repositories", "get", url) {
				items = append(items, appsv1.Repository{Repo: url, ConnectionState: s.getConnectionState(ctx, url)})
			}
		}
	}
	return &appsv1.RepositoryList{Items: items}, nil
}

// ListKsonnetApps returns list of Ksonnet apps in the repo
func (s *Server) ListApps(ctx context.Context, q *RepoAppsQuery) (*RepoAppsResponse, error) {
	if !s.enf.EnforceClaims(ctx.Value("claims"), "repositories", "get", q.Repo) {
		return nil, grpc.ErrPermissionDenied
	}
	repo, err := s.db.GetRepository(ctx, q.Repo)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.NotFound {
			repo = &appsv1.Repository{
				Repo: q.Repo,
			}
		} else {
			return nil, err
		}
	}

	// Test the repo
	conn, repoClient, err := s.repoClientset.NewRepositoryClient()
	if err != nil {
		return nil, err
	}
	defer util.Close(conn)

	revision := q.Revision
	if revision == "" {
		revision = "HEAD"
	}

	ksonnetRes, err := repoClient.ListDir(ctx, &repository.ListDirRequest{Repo: repo, Revision: revision, Path: "*app.yaml"})
	if err != nil {
		return nil, err
	}
	componentRes, err := repoClient.ListDir(ctx, &repository.ListDirRequest{Repo: repo, Revision: revision, Path: "*components/params.libsonnet"})
	if err != nil {
		return nil, err
	}

	helmRes, err := repoClient.ListDir(ctx, &repository.ListDirRequest{Repo: repo, Revision: revision, Path: "*Chart.yaml"})
	if err != nil {
		return nil, err
	}

	kustomizationRes, err := repoClient.ListDir(ctx, &repository.ListDirRequest{Repo: repo, Revision: revision, Path: "*kustomization.yaml"})
	if err != nil {
		return nil, err
	}

	componentDirs := make(map[string]interface{})
	for _, i := range componentRes.Items {
		d := filepath.Dir(filepath.Dir(i))
		componentDirs[d] = struct{}{}
	}

	items := make([]*AppInfo, 0)
	for _, i := range ksonnetRes.Items {
		d := filepath.Dir(i)
		if _, ok := componentDirs[d]; ok {
			items = append(items, &AppInfo{Type: string(repository.AppSourceKsonnet), Path: i})
		}
	}

	for i := range helmRes.Items {
		items = append(items, &AppInfo{Type: string(repository.AppSourceHelm), Path: helmRes.Items[i]})
	}

	for i := range kustomizationRes.Items {
		items = append(items, &AppInfo{Type: string(repository.AppSourceKustomize), Path: kustomizationRes.Items[i]})
	}

	return &RepoAppsResponse{Items: items}, nil
}

func (s *Server) GetAppDetails(ctx context.Context, q *RepoAppDetailsQuery) (*RepoAppDetailsResponse, error) {
	if !s.enf.EnforceClaims(ctx.Value("claims"), "repositories", "get", q.Repo) {
		return nil, grpc.ErrPermissionDenied
	}
	repo, err := s.db.GetRepository(ctx, q.Repo)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.NotFound {
			repo = &appsv1.Repository{
				Repo: q.Repo,
			}
		} else {
			return nil, err
		}
	}

	// Test the repo
	conn, repoClient, err := s.repoClientset.NewRepositoryClient()
	if err != nil {
		return nil, err
	}
	defer util.Close(conn)

	revision := q.Revision
	if revision == "" {
		revision = "HEAD"
	}

	appSpecRes, err := repoClient.GetFile(ctx, &repository.GetFileRequest{
		Repo:     repo,
		Revision: revision,
		Path:     q.Path,
	})
	if err != nil {
		return nil, err
	}

	appSourceType := repository.IdentifyAppSourceTypeByAppPath(q.Path)
	switch appSourceType {
	case repository.AppSourceKsonnet:
		var appSpec KsonnetAppSpec
		appSpec.Path = q.Path
		err = yaml.Unmarshal(appSpecRes.Data, &appSpec)
		if err != nil {
			return nil, err
		}
		return &RepoAppDetailsResponse{
			Type:    string(appSourceType),
			Ksonnet: &appSpec,
		}, nil
	case repository.AppSourceHelm:
		var appSpec HelmAppSpec
		appSpec.Path = q.Path
		err = yaml.Unmarshal(appSpecRes.Data, &appSpec)
		if err != nil {
			return nil, err
		}
		appDir := path.Dir(q.Path)
		valuesFilesRes, err := repoClient.ListDir(ctx, &repository.ListDirRequest{
			Revision: revision,
			Repo:     repo,
			Path:     path.Join(appDir, "*values*.yaml"),
		})
		if err != nil {
			return nil, err
		}
		appSpec.ValueFiles = make([]string, len(valuesFilesRes.Items))
		for i := range valuesFilesRes.Items {
			valueFilePath, err := filepath.Rel(appDir, valuesFilesRes.Items[i])
			if err != nil {
				return nil, err
			}
			appSpec.ValueFiles[i] = valueFilePath
		}
		return &RepoAppDetailsResponse{
			Type: string(appSourceType),
			Helm: &appSpec,
		}, nil
	case repository.AppSourceKustomize:
		appSpec := KustomizeAppSpec{
			Path: q.Path,
		}
		return &RepoAppDetailsResponse{
			Type:      string(appSourceType),
			Kustomize: &appSpec,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "specified application path is not supported")
}

// Create creates a repository
func (s *Server) Create(ctx context.Context, q *RepoCreateRequest) (*appsv1.Repository, error) {
	if !s.enf.EnforceClaims(ctx.Value("claims"), "repositories", "create", q.Repo.Repo) {
		return nil, grpc.ErrPermissionDenied
	}
	r := q.Repo
	err := git.TestRepo(git.NormalizeGitURL(r.Repo), r.Username, r.Password, r.SSHPrivateKey)
	if err != nil {
		return nil, err
	}

	r.ConnectionState = appsv1.ConnectionState{Status: appsv1.ConnectionStatusSuccessful}
	repo, err := s.db.CreateRepository(ctx, r)
	if status.Convert(err).Code() == codes.AlreadyExists {
		// act idempotent if existing spec matches new spec
		existing, getErr := s.db.GetRepository(ctx, r.Repo)
		if getErr != nil {
			return nil, status.Errorf(codes.Internal, "unable to check existing repository details: %v", getErr)
		}

		// repository ConnectionState may differ, so make consistent before testing
		existing.ConnectionState = r.ConnectionState
		if reflect.DeepEqual(existing, r) {
			repo, err = existing, nil
		} else if q.Upsert {
			return s.Update(ctx, &RepoUpdateRequest{Repo: r})
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "existing repository spec is different; use upsert flag to force update")
		}
	}
	return &appsv1.Repository{Repo: repo.Repo}, err
}

// Update updates a repository
func (s *Server) Update(ctx context.Context, q *RepoUpdateRequest) (*appsv1.Repository, error) {
	if !s.enf.EnforceClaims(ctx.Value("claims"), "repositories", "update", q.Repo.Repo) {
		return nil, grpc.ErrPermissionDenied
	}
	_, err := s.db.UpdateRepository(ctx, q.Repo)
	return &appsv1.Repository{Repo: q.Repo.Repo}, err
}

// Delete updates a repository
func (s *Server) Delete(ctx context.Context, q *RepoQuery) (*RepoResponse, error) {
	if !s.enf.EnforceClaims(ctx.Value("claims"), "repositories", "delete", q.Repo) {
		return nil, grpc.ErrPermissionDenied
	}
	err := s.db.DeleteRepository(ctx, q.Repo)
	return &RepoResponse{}, err
}
