// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/application/application.proto

/*
Package application is a generated protocol buffer package.

Application Service

Application Service API performs CRUD actions against application resources

It is generated from these files:
	server/application/application.proto

It has these top-level messages:
	ApplicationQuery
	ApplicationResponse
	DeleteApplicationRequest
	ApplicationSyncRequest
	ApplicationSyncResult
	ResourceDetails
	DeletePodQuery
	PodLogsQuery
	LogEntry
*/
package application

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "k8s.io/api/core/v1"
import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ApplicationQuery is a query for application resources
type ApplicationQuery struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ApplicationQuery) Reset()                    { *m = ApplicationQuery{} }
func (m *ApplicationQuery) String() string            { return proto.CompactTextString(m) }
func (*ApplicationQuery) ProtoMessage()               {}
func (*ApplicationQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplicationQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplicationResponse struct {
}

func (m *ApplicationResponse) Reset()                    { *m = ApplicationResponse{} }
func (m *ApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*ApplicationResponse) ProtoMessage()               {}
func (*ApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeleteApplicationRequest struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Server    string `protobuf:"bytes,3,opt,name=server" json:"server,omitempty"`
	Force     bool   `protobuf:"varint,4,opt,name=force" json:"force,omitempty"`
}

func (m *DeleteApplicationRequest) Reset()                    { *m = DeleteApplicationRequest{} }
func (m *DeleteApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationRequest) ProtoMessage()               {}
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteApplicationRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteApplicationRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *DeleteApplicationRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// ApplicationSyncRequest is a request to apply the config state to live state
type ApplicationSyncRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Revision string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	DryRun   bool   `protobuf:"varint,3,opt,name=dryRun" json:"dryRun,omitempty"`
}

func (m *ApplicationSyncRequest) Reset()                    { *m = ApplicationSyncRequest{} }
func (m *ApplicationSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplicationSyncRequest) ProtoMessage()               {}
func (*ApplicationSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ApplicationSyncRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplicationSyncRequest) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ApplicationSyncRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

// ApplicationSyncResult is a result of a sync requeswt
type ApplicationSyncResult struct {
	Message   string             `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Resources []*ResourceDetails `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
}

func (m *ApplicationSyncResult) Reset()                    { *m = ApplicationSyncResult{} }
func (m *ApplicationSyncResult) String() string            { return proto.CompactTextString(m) }
func (*ApplicationSyncResult) ProtoMessage()               {}
func (*ApplicationSyncResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ApplicationSyncResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApplicationSyncResult) GetResources() []*ResourceDetails {
	if m != nil {
		return m.Resources
	}
	return nil
}

type ResourceDetails struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind      string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Message   string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *ResourceDetails) Reset()                    { *m = ResourceDetails{} }
func (m *ResourceDetails) String() string            { return proto.CompactTextString(m) }
func (*ResourceDetails) ProtoMessage()               {}
func (*ResourceDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResourceDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceDetails) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ResourceDetails) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ResourceDetails) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeletePodQuery struct {
	ApplicationName string `protobuf:"bytes,1,opt,name=applicationName" json:"applicationName,omitempty"`
	PodName         string `protobuf:"bytes,2,opt,name=podName" json:"podName,omitempty"`
}

func (m *DeletePodQuery) Reset()                    { *m = DeletePodQuery{} }
func (m *DeletePodQuery) String() string            { return proto.CompactTextString(m) }
func (*DeletePodQuery) ProtoMessage()               {}
func (*DeletePodQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeletePodQuery) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *DeletePodQuery) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

type PodLogsQuery struct {
	ApplicationName string                                     `protobuf:"bytes,1,opt,name=applicationName" json:"applicationName,omitempty"`
	PodName         string                                     `protobuf:"bytes,2,opt,name=podName" json:"podName,omitempty"`
	Container       string                                     `protobuf:"bytes,3,opt,name=container" json:"container,omitempty"`
	SinceSeconds    int64                                      `protobuf:"varint,4,opt,name=sinceSeconds" json:"sinceSeconds,omitempty"`
	SinceTime       *k8s_io_apimachinery_pkg_apis_meta_v1.Time `protobuf:"bytes,5,opt,name=sinceTime" json:"sinceTime,omitempty"`
	TailLines       int64                                      `protobuf:"varint,6,opt,name=tailLines" json:"tailLines,omitempty"`
	Follow          bool                                       `protobuf:"varint,7,opt,name=follow" json:"follow,omitempty"`
}

func (m *PodLogsQuery) Reset()                    { *m = PodLogsQuery{} }
func (m *PodLogsQuery) String() string            { return proto.CompactTextString(m) }
func (*PodLogsQuery) ProtoMessage()               {}
func (*PodLogsQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PodLogsQuery) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *PodLogsQuery) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *PodLogsQuery) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *PodLogsQuery) GetSinceSeconds() int64 {
	if m != nil {
		return m.SinceSeconds
	}
	return 0
}

func (m *PodLogsQuery) GetSinceTime() *k8s_io_apimachinery_pkg_apis_meta_v1.Time {
	if m != nil {
		return m.SinceTime
	}
	return nil
}

func (m *PodLogsQuery) GetTailLines() int64 {
	if m != nil {
		return m.TailLines
	}
	return 0
}

func (m *PodLogsQuery) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

type LogEntry struct {
	Content   string                                     `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	TimeStamp *k8s_io_apimachinery_pkg_apis_meta_v1.Time `protobuf:"bytes,2,opt,name=timeStamp" json:"timeStamp,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LogEntry) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *LogEntry) GetTimeStamp() *k8s_io_apimachinery_pkg_apis_meta_v1.Time {
	if m != nil {
		return m.TimeStamp
	}
	return nil
}

func init() {
	proto.RegisterType((*ApplicationQuery)(nil), "application.ApplicationQuery")
	proto.RegisterType((*ApplicationResponse)(nil), "application.ApplicationResponse")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "application.DeleteApplicationRequest")
	proto.RegisterType((*ApplicationSyncRequest)(nil), "application.ApplicationSyncRequest")
	proto.RegisterType((*ApplicationSyncResult)(nil), "application.ApplicationSyncResult")
	proto.RegisterType((*ResourceDetails)(nil), "application.ResourceDetails")
	proto.RegisterType((*DeletePodQuery)(nil), "application.DeletePodQuery")
	proto.RegisterType((*PodLogsQuery)(nil), "application.PodLogsQuery")
	proto.RegisterType((*LogEntry)(nil), "application.LogEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApplicationService service

type ApplicationServiceClient interface {
	// List returns list of applications
	List(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList, error)
	// Watch returns stream of application change events.
	Watch(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (ApplicationService_WatchClient, error)
	// Create creates an application
	Create(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Get returns an application by name
	Get(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Update updates an application
	Update(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Delete deletes an application
	Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// Sync syncs an application to its target state
	Sync(ctx context.Context, in *ApplicationSyncRequest, opts ...grpc.CallOption) (*ApplicationSyncResult, error)
	// PodLogs returns stream of log entries for the specified pod. Pod
	DeletePod(ctx context.Context, in *DeletePodQuery, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// PodLogs returns stream of log entries for the specified pod. Pod
	PodLogs(ctx context.Context, in *PodLogsQuery, opts ...grpc.CallOption) (ApplicationService_PodLogsClient, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) List(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList)
	err := grpc.Invoke(ctx, "/application.ApplicationService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Watch(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (ApplicationService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ApplicationService_serviceDesc.Streams[0], c.cc, "/application.ApplicationService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationService_WatchClient interface {
	Recv() (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent, error)
	grpc.ClientStream
}

type applicationServiceWatchClient struct {
	grpc.ClientStream
}

func (x *applicationServiceWatchClient) Recv() (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent, error) {
	m := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *applicationServiceClient) Create(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Get(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Update(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Sync(ctx context.Context, in *ApplicationSyncRequest, opts ...grpc.CallOption) (*ApplicationSyncResult, error) {
	out := new(ApplicationSyncResult)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) DeletePod(ctx context.Context, in *DeletePodQuery, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := grpc.Invoke(ctx, "/application.ApplicationService/DeletePod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) PodLogs(ctx context.Context, in *PodLogsQuery, opts ...grpc.CallOption) (ApplicationService_PodLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ApplicationService_serviceDesc.Streams[1], c.cc, "/application.ApplicationService/PodLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationServicePodLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationService_PodLogsClient interface {
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type applicationServicePodLogsClient struct {
	grpc.ClientStream
}

func (x *applicationServicePodLogsClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ApplicationService service

type ApplicationServiceServer interface {
	// List returns list of applications
	List(context.Context, *ApplicationQuery) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList, error)
	// Watch returns stream of application change events.
	Watch(*ApplicationQuery, ApplicationService_WatchServer) error
	// Create creates an application
	Create(context.Context, *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Get returns an application by name
	Get(context.Context, *ApplicationQuery) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Update updates an application
	Update(context.Context, *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Delete deletes an application
	Delete(context.Context, *DeleteApplicationRequest) (*ApplicationResponse, error)
	// Sync syncs an application to its target state
	Sync(context.Context, *ApplicationSyncRequest) (*ApplicationSyncResult, error)
	// PodLogs returns stream of log entries for the specified pod. Pod
	DeletePod(context.Context, *DeletePodQuery) (*ApplicationResponse, error)
	// PodLogs returns stream of log entries for the specified pod. Pod
	PodLogs(*PodLogsQuery, ApplicationService_PodLogsServer) error
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).List(ctx, req.(*ApplicationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplicationQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationServiceServer).Watch(m, &applicationServiceWatchServer{stream})
}

type ApplicationService_WatchServer interface {
	Send(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent) error
	grpc.ServerStream
}

type applicationServiceWatchServer struct {
	grpc.ServerStream
}

func (x *applicationServiceWatchServer) Send(m *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ApplicationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Create(ctx, req.(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Get(ctx, req.(*ApplicationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Update(ctx, req.(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Delete(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Sync(ctx, req.(*ApplicationSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).DeletePod(ctx, req.(*DeletePodQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_PodLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PodLogsQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationServiceServer).PodLogs(m, &applicationServicePodLogsServer{stream})
}

type ApplicationService_PodLogsServer interface {
	Send(*LogEntry) error
	grpc.ServerStream
}

type applicationServicePodLogsServer struct {
	grpc.ServerStream
}

func (x *applicationServicePodLogsServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "application.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ApplicationService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ApplicationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApplicationService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApplicationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApplicationService_Delete_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _ApplicationService_Sync_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _ApplicationService_DeletePod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ApplicationService_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PodLogs",
			Handler:       _ApplicationService_PodLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/application/application.proto",
}

func init() { proto.RegisterFile("server/application/application.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0x97, 0xb3, 0x9b, 0x4d, 0x76, 0x52, 0x51, 0x34, 0x24, 0x91, 0xeb, 0xa6, 0x62, 0x35, 0x6d,
	0xd1, 0x12, 0x84, 0x9d, 0x04, 0x10, 0x28, 0x82, 0x03, 0xa5, 0xe5, 0x9f, 0x22, 0x14, 0x9c, 0x22,
	0x24, 0x2e, 0x30, 0xb5, 0x5f, 0x1d, 0xb3, 0xf6, 0x8c, 0x3b, 0x33, 0x6b, 0xb4, 0x94, 0x1c, 0xe0,
	0xc6, 0x09, 0x21, 0xb8, 0xf3, 0x39, 0xf8, 0x08, 0x9c, 0xf9, 0x0a, 0x9c, 0xf8, 0x14, 0x68, 0xc6,
	0xf6, 0xda, 0xde, 0xdd, 0x6c, 0x00, 0xed, 0xa1, 0x27, 0xcf, 0xfb, 0x33, 0xef, 0xf7, 0x7b, 0x6f,
	0x66, 0xde, 0x33, 0xba, 0x23, 0x41, 0xe4, 0x20, 0x3c, 0x9a, 0x65, 0x49, 0x1c, 0x50, 0x15, 0x73,
	0xd6, 0x5c, 0xbb, 0x99, 0xe0, 0x8a, 0xe3, 0xad, 0x86, 0xca, 0xd9, 0x8e, 0x78, 0xc4, 0x8d, 0xde,
	0xd3, 0xab, 0xc2, 0xc5, 0xd9, 0x8b, 0x38, 0x8f, 0x12, 0xf0, 0x68, 0x16, 0x7b, 0x94, 0x31, 0xae,
	0x8c, 0xb3, 0x2c, 0xad, 0x64, 0xf4, 0x96, 0x74, 0x63, 0x6e, 0xac, 0x01, 0x17, 0xe0, 0xe5, 0x87,
	0x5e, 0x04, 0x0c, 0x04, 0x55, 0x10, 0x96, 0x3e, 0xaf, 0xd7, 0x3e, 0x29, 0x0d, 0xce, 0x63, 0x06,
	0x62, 0xe2, 0x65, 0xa3, 0x48, 0x2b, 0xa4, 0x97, 0x82, 0xa2, 0x8b, 0x76, 0x7d, 0x14, 0xc5, 0xea,
	0x7c, 0xfc, 0xc8, 0x0d, 0x78, 0xea, 0x51, 0x61, 0x88, 0x7d, 0x6d, 0x16, 0xaf, 0x06, 0x61, 0xbd,
	0xbb, 0x99, 0x5e, 0x7e, 0x48, 0x93, 0xec, 0x9c, 0xce, 0x85, 0x22, 0x2f, 0xa1, 0xe7, 0xdf, 0xad,
	0xfd, 0x3e, 0x1d, 0x83, 0x98, 0x60, 0x8c, 0xba, 0x8c, 0xa6, 0x60, 0x5b, 0x03, 0x6b, 0xd8, 0xf7,
	0xcd, 0x9a, 0xec, 0xa0, 0x17, 0x1a, 0x7e, 0x3e, 0xc8, 0x8c, 0x33, 0x09, 0xe4, 0x5b, 0x64, 0xdf,
	0x87, 0x04, 0x14, 0xb4, 0x8c, 0x4f, 0xc6, 0x20, 0xd5, 0xa2, 0x30, 0x78, 0x0f, 0xf5, 0xf5, 0x57,
	0x66, 0x34, 0x00, 0x7b, 0xcd, 0x18, 0x6a, 0x05, 0xde, 0x45, 0xbd, 0xe2, 0x68, 0xec, 0x8e, 0x31,
	0x95, 0x12, 0xde, 0x46, 0xeb, 0x8f, 0xb9, 0x08, 0xc0, 0xee, 0x0e, 0xac, 0xe1, 0xa6, 0x5f, 0x08,
	0xe4, 0x2b, 0xb4, 0xdb, 0x40, 0x3d, 0x9b, 0xb0, 0x60, 0x19, 0xb2, 0x83, 0x36, 0x05, 0xe4, 0xb1,
	0x8c, 0x39, 0x2b, 0x81, 0xa7, 0xb2, 0xc6, 0x0d, 0xc5, 0xc4, 0x1f, 0x33, 0x83, 0xbb, 0xe9, 0x97,
	0x12, 0x49, 0xd1, 0xce, 0x1c, 0x82, 0x1c, 0x27, 0x0a, 0xdb, 0x68, 0x23, 0x05, 0x29, 0x69, 0x54,
	0x61, 0x54, 0x22, 0x3e, 0x46, 0x7d, 0x01, 0x92, 0x8f, 0x45, 0x00, 0xd2, 0x5e, 0x1b, 0x74, 0x86,
	0x5b, 0x47, 0x7b, 0x6e, 0xf3, 0x72, 0xf9, 0xa5, 0xf5, 0x3e, 0x28, 0x1a, 0x27, 0xd2, 0xaf, 0xdd,
	0xc9, 0x13, 0x74, 0x7d, 0xc6, 0xba, 0x30, 0x13, 0x8c, 0xba, 0xa3, 0x98, 0x85, 0x65, 0x16, 0x66,
	0xdd, 0xae, 0x6b, 0x67, 0xb6, 0xae, 0x0d, 0xba, 0xdd, 0x16, 0x5d, 0xf2, 0x10, 0x3d, 0x57, 0x9c,
	0xdf, 0x29, 0x0f, 0x8b, 0xc3, 0x1f, 0xa2, 0xeb, 0x0d, 0xba, 0x9f, 0xd4, 0xe0, 0xb3, 0x6a, 0x1d,
	0x35, 0xe3, 0xa1, 0xf1, 0x28, 0xa8, 0x54, 0x22, 0xf9, 0x75, 0x0d, 0x5d, 0x3b, 0xe5, 0xe1, 0x09,
	0x8f, 0xe4, 0xca, 0x82, 0xea, 0x14, 0x03, 0xce, 0x14, 0xd5, 0x6f, 0xa4, 0x4a, 0x71, 0xaa, 0xc0,
	0x04, 0x5d, 0x93, 0x31, 0x0b, 0xe0, 0x0c, 0x02, 0xce, 0x42, 0x69, 0xf2, 0xec, 0xf8, 0x2d, 0x1d,
	0xfe, 0x10, 0xf5, 0x8d, 0xfc, 0x30, 0x4e, 0xc1, 0x5e, 0x1f, 0x58, 0xc3, 0xad, 0xa3, 0x7d, 0xb7,
	0x78, 0x80, 0x6e, 0xf3, 0x01, 0xba, 0xd9, 0x28, 0xd2, 0x0a, 0xe9, 0xea, 0x07, 0xe8, 0xe6, 0x87,
	0xae, 0xde, 0xe1, 0xd7, 0x9b, 0x35, 0x17, 0x7d, 0x3e, 0x27, 0x31, 0x03, 0x69, 0xf7, 0x0c, 0x54,
	0xad, 0xd0, 0xd7, 0xe9, 0x31, 0x4f, 0x12, 0xfe, 0x8d, 0xbd, 0x51, 0x5c, 0xa7, 0x42, 0x22, 0x0c,
	0x6d, 0x9e, 0xf0, 0xe8, 0x01, 0x53, 0x62, 0xa2, 0xf3, 0xd4, 0xe4, 0x81, 0xa9, 0xea, 0x06, 0x95,
	0xa2, 0x66, 0xa9, 0xe2, 0x14, 0xce, 0x14, 0x4d, 0x33, 0x53, 0x83, 0xff, 0xc8, 0x72, 0xba, 0xf9,
	0xe8, 0x6f, 0x84, 0x70, 0xf3, 0xfe, 0x82, 0xc8, 0xe3, 0x00, 0xf0, 0x4f, 0x16, 0xea, 0x9e, 0xc4,
	0x52, 0xe1, 0x5b, 0xad, 0x8b, 0x39, 0xdb, 0x06, 0x9c, 0x8f, 0xdd, 0xba, 0xcd, 0xb8, 0x55, 0x9b,
	0x31, 0x8b, 0x2f, 0x83, 0xb0, 0x46, 0x6f, 0xc6, 0xa8, 0xda, 0x4c, 0x33, 0x98, 0x86, 0x22, 0x7b,
	0x3f, 0xfc, 0xf9, 0xd7, 0x2f, 0x6b, 0xbb, 0x78, 0xdb, 0x74, 0xc3, 0xfc, 0xb0, 0xd9, 0x9c, 0x24,
	0xfe, 0xcd, 0x42, 0xeb, 0x9f, 0x53, 0x15, 0x9c, 0x5f, 0x45, 0xe9, 0x74, 0x35, 0x94, 0x0c, 0xd6,
	0x83, 0x1c, 0x98, 0x22, 0xb7, 0x0d, 0xb1, 0x5b, 0xf8, 0x66, 0x45, 0x4c, 0x2a, 0x01, 0x34, 0x6d,
	0xf1, 0x3b, 0xb0, 0xf0, 0xef, 0x16, 0xea, 0xbd, 0x27, 0x80, 0x2a, 0xc0, 0xef, 0xaf, 0x86, 0x83,
	0xb3, 0xa2, 0x38, 0xe4, 0x45, 0x93, 0xc1, 0x0d, 0xb2, 0xb0, 0xb4, 0xc7, 0xd6, 0x3e, 0xfe, 0xd9,
	0x42, 0x9d, 0x0f, 0xe0, 0xca, 0xe3, 0x5e, 0x15, 0x9f, 0xb9, 0x8a, 0x36, 0xf9, 0x78, 0x4f, 0x75,
	0x57, 0xba, 0xc0, 0x7f, 0x58, 0xa8, 0xf7, 0x59, 0x16, 0x3e, 0x8b, 0xf5, 0xf4, 0x0c, 0xff, 0x97,
	0x9d, 0x3b, 0x8b, 0xf9, 0xeb, 0xc7, 0x16, 0x52, 0x45, 0x5d, 0x93, 0x88, 0xae, 0x6f, 0x8e, 0x7a,
	0x45, 0x0f, 0xc5, 0x77, 0x5b, 0xd1, 0x2f, 0x1b, 0x8c, 0xce, 0xe0, 0xb2, 0x83, 0x98, 0x8e, 0xd5,
	0xb2, 0x86, 0xfb, 0x4b, 0x6b, 0xf8, 0x1d, 0xea, 0xea, 0x91, 0x84, 0x6f, 0x5f, 0x16, 0xae, 0x31,
	0x12, 0x1d, 0xb2, 0xdc, 0x49, 0x4f, 0x35, 0xf2, 0x8a, 0x41, 0xbd, 0x4b, 0x06, 0x4b, 0x50, 0x3d,
	0x39, 0x61, 0x81, 0xce, 0xfa, 0x47, 0x0b, 0xf5, 0xa7, 0xa3, 0x03, 0xdf, 0x5c, 0x90, 0x79, 0x35,
	0x52, 0xfe, 0x45, 0xbe, 0xef, 0x18, 0xe4, 0x37, 0xf7, 0xdf, 0x58, 0x8c, 0x3c, 0x33, 0x24, 0x2e,
	0xbc, 0x8c, 0x87, 0xd2, 0x7b, 0x5a, 0x4e, 0x86, 0x0b, 0xfc, 0xbd, 0x85, 0x36, 0xca, 0x79, 0x83,
	0x6f, 0xb4, 0xc0, 0x9a, 0x53, 0xc8, 0xd9, 0x69, 0x99, 0xaa, 0x56, 0x4c, 0xee, 0x19, 0xf0, 0xb7,
	0xf1, 0xf1, 0xff, 0x02, 0xf7, 0x12, 0x1e, 0xc9, 0x03, 0xeb, 0xde, 0xc1, 0x17, 0xee, 0xb2, 0xbf,
	0xb2, 0xf9, 0x5f, 0xce, 0x47, 0x3d, 0xf3, 0x07, 0xf6, 0xda, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5f, 0x73, 0xc7, 0xe2, 0x8f, 0x0a, 0x00, 0x00,
}
