package common

// Default service addresses and URLS of Argo CD internal services
const (
	// DefaultAppControllerServerAddr is the gRPC address of the Argo CD app controller server
	DefaultAppControllerServerAddr = "application-controller:8083"
	// DefaultRepoServerAddr is the gRPC address of the Argo CD repo server
	DefaultRepoServerAddr = "argocd-repo-server:8081"
	// DefaultDexServerAddr is the HTTP address of the Dex OIDC server, which we run a reverse proxy against
	DefaultDexServerAddr = "http://dex-server:5556"
)

// Kubernetes ConfigMap and Secret resource names which hold Argo CD settings
const (
	ArgoCDConfigMapName     = "argocd-cm"
	ArgoCDSecretName        = "argocd-secret"
	ArgoCDRBACConfigMapName = "argocd-rbac-cm"
)

// Argo CD application related constants
const (
	// KubernetesInternalAPIServerAddr is address of the k8s API server when accessing internal to the cluster
	KubernetesInternalAPIServerAddr = "https://kubernetes.default.svc"
	// DefaultAppProjectName contains name of 'default' app project, which is available in every Argo CD installation
	DefaultAppProjectName = "default"
	// ArgoCDAdminUsername is the username of the 'admin' user
	ArgoCDAdminUsername = "admin"
	// ArgoCDUserAgentName is the default user-agent name used by the gRPC API client library and grpc-gateway
	ArgoCDUserAgentName = "argocd-client"
	// AuthCookieName is the HTTP cookie name where we store our auth token
	AuthCookieName = "argocd.token"
)

// Dex related constants
const (
	// DexAPIEndpoint is the endpoint where we serve the Dex API server
	DexAPIEndpoint = "/api/dex"
	// LoginEndpoint is Argo CD's shorthand login endpoint which redirects to dex's OAuth 2.0 provider's consent page
	LoginEndpoint = "/auth/login"
	// CallbackEndpoint is Argo CD's final callback endpoint we reach after OAuth 2.0 login flow has been completed
	CallbackEndpoint = "/auth/callback"
	// ArgoCDClientAppName is name of the Oauth client app used when registering our web app to dex
	ArgoCDClientAppName = "Argo CD"
	// ArgoCDClientAppID is the Oauth client ID we will use when registering our app to dex
	ArgoCDClientAppID = "argo-cd"
	// ArgoCDCLIClientAppName is name of the Oauth client app used when registering our CLI to dex
	ArgoCDCLIClientAppName = "Argo CD CLI"
	// ArgoCDCLIClientAppID is the Oauth client ID we will use when registering our CLI to dex
	ArgoCDCLIClientAppID = "argo-cd-cli"
)

// Resource metadata labels and annotations (keys and values) used by Argo CD components
const (
	// LabelKeyAppInstance is the label key to use to uniquely identify the instance of an application
	// The Argo CD application name is used as the instance name
	LabelKeyAppInstance = "app.kubernetes.io/instance"
	// LegacyLabelApplicationName is the legacy label (v0.10 and below) and is superceded by 'app.kubernetes.io/instance'
	LabelKeyLegacyApplicationName = "applications.argoproj.io/app-name"
	// LabelKeySecretType contains the type of argocd secret (currently: 'cluster')
	LabelKeySecretType = "argocd.argoproj.io/secret-type"
	// LabelValueSecretTypeCluster indicates a secret type of cluster
	LabelValueSecretTypeCluster = "cluster"

	// AnnotationKeyHook contains the hook type of a resource
	AnnotationKeyHook = "argocd.argoproj.io/hook"
	// AnnotationKeyHookDeletePolicy is the policy of deleting a hook
	AnnotationKeyHookDeletePolicy = "argocd.argoproj.io/hook-delete-policy"
	// AnnotationKeyRefresh is the annotation key in the application which is updated with an
	// arbitrary value (i.e. timestamp) on a git event, to  force the controller to wake up and
	// re-evaluate the application
	AnnotationKeyRefresh = "argocd.argoproj.io/refresh"
	// AnnotationKeyConnectionStatus contains connection state status
	AnnotationKeyConnectionStatus = "argocd.argoproj.io/connection-status"
	// AnnotationKeyConnectionMessage contains additional information about connection status
	AnnotationKeyConnectionMessage = "argocd.argoproj.io/connection-message"
	// AnnotationConnectionModifiedAt contains timestamp when connection state had been modified
	AnnotationConnectionModifiedAt = "argocd.argoproj.io/connection-modified-at"
	// AnnotationKeyManagedBy is annotation name which indicates that k8s resource is managed by an application.
	AnnotationKeyManagedBy = "managed-by"
	// AnnotationValueManagedByArgoCD is a 'managed-by' annotation value for resources managed by Argo CD
	AnnotationValueManagedByArgoCD = "argocd.argoproj.io"
	// AnnotationKeyHelmHook is the helm hook annotation
	AnnotationKeyHelmHook = "helm.sh/hook"
	// AnnotationValueHelmHookCRDInstall is a value of crd helm hook
	AnnotationValueHelmHookCRDInstall = "crd-install"
	// ResourcesFinalizerName the finalizer value which we inject to finalize deletion of an application
	ResourcesFinalizerName = "resources-finalizer.argocd.argoproj.io"
)

// Environment variables for tuning and debugging Argo CD
const (
	// EnvVarSSODebug is an environment variable to enable additional OAuth debugging in the API server
	EnvVarSSODebug = "ARGOCD_SSO_DEBUG"
	// EnvVarRBACDebug is an environment variable to enable additional RBAC debugging in the API server
	EnvVarRBACDebug = "ARGOCD_RBAC_DEBUG"
	// EnvVarLegacyLabels is an environment variable to use the legacy 'applications.argoproj.io/app-name' label instead of 'app.kubernetes.io/instance'
	EnvVarLegacyLabels = "ARGOCD_LEGACY_LABELS"
)
