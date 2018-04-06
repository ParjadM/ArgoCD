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
func (*PodLogsQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6f, 0xe4, 0x34,
	0x14, 0x57, 0x3a, 0xb3, 0xd3, 0x8e, 0xbb, 0xd2, 0x22, 0xb3, 0xad, 0xb2, 0xa1, 0x2b, 0x46, 0xde,
	0x5d, 0x54, 0x8a, 0x48, 0xda, 0xc2, 0x01, 0x55, 0x5c, 0x58, 0x76, 0xf9, 0xa7, 0x0a, 0x2d, 0x29,
	0x08, 0x89, 0x0b, 0x78, 0x93, 0x47, 0x6a, 0x9a, 0xd8, 0x59, 0xdb, 0x13, 0x34, 0x5a, 0x7a, 0x80,
	0x2f, 0x80, 0x10, 0xdc, 0xf9, 0x1c, 0x7c, 0x04, 0xce, 0xdc, 0x39, 0xf1, 0x41, 0x90, 0x9d, 0xa4,
	0x71, 0x66, 0xa6, 0xb3, 0x42, 0x9a, 0x03, 0xa7, 0xf8, 0xfd, 0xf1, 0xfb, 0xfd, 0xfc, 0x9e, 0x9f,
	0x5f, 0xd0, 0x7d, 0x05, 0xb2, 0x02, 0x19, 0xd1, 0xb2, 0xcc, 0x59, 0x42, 0x35, 0x13, 0xdc, 0x5d,
	0x87, 0xa5, 0x14, 0x5a, 0xe0, 0x6d, 0x47, 0x15, 0xdc, 0xce, 0x44, 0x26, 0xac, 0x3e, 0x32, 0xab,
	0xda, 0x25, 0xd8, 0xcb, 0x84, 0xc8, 0x72, 0x88, 0x68, 0xc9, 0x22, 0xca, 0xb9, 0xd0, 0xd6, 0x59,
	0x35, 0x56, 0x72, 0xf1, 0x8e, 0x0a, 0x99, 0xb0, 0xd6, 0x44, 0x48, 0x88, 0xaa, 0xa3, 0x28, 0x03,
	0x0e, 0x92, 0x6a, 0x48, 0x1b, 0x9f, 0xb7, 0x3b, 0x9f, 0x82, 0x26, 0xe7, 0x8c, 0x83, 0x9c, 0x45,
	0xe5, 0x45, 0x66, 0x14, 0x2a, 0x2a, 0x40, 0xd3, 0x65, 0xbb, 0x3e, 0xce, 0x98, 0x3e, 0x9f, 0x3e,
	0x0d, 0x13, 0x51, 0x44, 0x54, 0x5a, 0x62, 0xdf, 0xd9, 0xc5, 0x9b, 0x49, 0xda, 0xed, 0x76, 0x8f,
	0x57, 0x1d, 0xd1, 0xbc, 0x3c, 0xa7, 0x0b, 0xa1, 0xc8, 0x6b, 0xe8, 0xa5, 0xf7, 0x3a, 0xbf, 0xcf,
	0xa6, 0x20, 0x67, 0x18, 0xa3, 0x21, 0xa7, 0x05, 0xf8, 0xde, 0xc4, 0xdb, 0x1f, 0xc7, 0x76, 0x4d,
	0x76, 0xd0, 0xcb, 0x8e, 0x5f, 0x0c, 0xaa, 0x14, 0x5c, 0x01, 0x49, 0x91, 0xff, 0x08, 0x72, 0xd0,
	0xd0, 0x33, 0x3e, 0x9b, 0x82, 0xd2, 0xcb, 0xc2, 0xe0, 0x3d, 0x34, 0x36, 0x5f, 0x55, 0xd2, 0x04,
	0xfc, 0x0d, 0x6b, 0xe8, 0x14, 0x78, 0x17, 0x8d, 0xea, 0xd2, 0xf8, 0x03, 0x6b, 0x6a, 0x24, 0xf2,
	0x0d, 0xda, 0x75, 0xe2, 0x9f, 0xcd, 0x78, 0xb2, 0x0a, 0x23, 0x40, 0x5b, 0x12, 0x2a, 0xa6, 0x98,
	0xe0, 0x0d, 0xc4, 0x95, 0x6c, 0x10, 0x52, 0x39, 0x8b, 0xa7, 0xdc, 0x22, 0x6c, 0xc5, 0x8d, 0x44,
	0x0a, 0xb4, 0xb3, 0x80, 0xa0, 0xa6, 0xb9, 0xc6, 0x3e, 0xda, 0x2c, 0x40, 0x29, 0x9a, 0xb5, 0x18,
	0xad, 0x88, 0x4f, 0xd0, 0x58, 0x82, 0x12, 0x53, 0x99, 0x80, 0xf2, 0x37, 0x26, 0x83, 0xfd, 0xed,
	0xe3, 0xbd, 0xd0, 0xbd, 0x46, 0x71, 0x63, 0x7d, 0x04, 0x9a, 0xb2, 0x5c, 0xc5, 0x9d, 0x3b, 0x79,
	0x86, 0x6e, 0xcd, 0x59, 0x97, 0x9e, 0x04, 0xa3, 0xe1, 0x05, 0xe3, 0x69, 0x73, 0x0a, 0xbb, 0xee,
	0x67, 0x70, 0x30, 0x9f, 0x41, 0x87, 0xee, 0xb0, 0x47, 0x97, 0xfc, 0xb6, 0x81, 0x6e, 0x3e, 0x11,
	0xe9, 0xa9, 0xc8, 0x54, 0x5d, 0xe5, 0x7d, 0x74, 0xcb, 0x61, 0xfb, 0x69, 0x87, 0x3d, 0xaf, 0x36,
	0x41, 0x4b, 0x91, 0x5a, 0x8f, 0x9a, 0x49, 0x2b, 0x1a, 0x32, 0x89, 0xe0, 0x9a, 0x9a, 0x7b, 0xdb,
	0x92, 0xb9, 0x52, 0x60, 0x82, 0x6e, 0x2a, 0xc6, 0x13, 0x38, 0x83, 0x44, 0xf0, 0x54, 0x59, 0x46,
	0x83, 0xb8, 0xa7, 0xc3, 0x1f, 0xa1, 0xb1, 0x95, 0x3f, 0x67, 0x05, 0xf8, 0x37, 0x26, 0xde, 0xfe,
	0xf6, 0xf1, 0x41, 0x58, 0x37, 0x45, 0xe8, 0x36, 0x45, 0x58, 0x5e, 0x64, 0x46, 0xa1, 0x42, 0xd3,
	0x14, 0x61, 0x75, 0x14, 0x9a, 0x1d, 0x71, 0xb7, 0xd9, 0x70, 0x31, 0x99, 0x3c, 0x65, 0x1c, 0x94,
	0x3f, 0xb2, 0x50, 0x9d, 0xc2, 0x14, 0xfe, 0x5b, 0x91, 0xe7, 0xe2, 0x7b, 0x7f, 0xb3, 0x2e, 0x7c,
	0x2d, 0x11, 0x8e, 0xb6, 0x4e, 0x45, 0xf6, 0x98, 0x6b, 0x39, 0x33, 0xe7, 0x34, 0xe4, 0x81, 0xeb,
	0xb6, 0xd6, 0x8d, 0x68, 0x58, 0x6a, 0x56, 0xc0, 0x99, 0xa6, 0x45, 0x69, 0x73, 0xf0, 0x1f, 0x59,
	0x5e, 0x6d, 0x3e, 0xfe, 0x7b, 0x8c, 0xb0, 0x7b, 0xd3, 0x40, 0x56, 0x2c, 0x01, 0xfc, 0xb3, 0x87,
	0x86, 0xa7, 0x4c, 0x69, 0x7c, 0xb7, 0x77, 0x85, 0xe6, 0x5b, 0x33, 0xf8, 0x24, 0xec, 0x5a, 0x3f,
	0x6c, 0x5b, 0xdf, 0x2e, 0xbe, 0x4e, 0xd2, 0x0e, 0xdd, 0x8d, 0xd1, 0xb6, 0xbe, 0x1b, 0xcc, 0x40,
	0x91, 0xbd, 0x9f, 0xfe, 0xfa, 0xe7, 0xd7, 0x8d, 0x5d, 0x7c, 0xdb, 0xbe, 0x50, 0xd5, 0x91, 0xfb,
	0x60, 0x28, 0xfc, 0xbb, 0x87, 0x6e, 0x7c, 0x49, 0x75, 0x72, 0xfe, 0x22, 0x4a, 0x4f, 0xd6, 0x43,
	0xc9, 0x62, 0x3d, 0xae, 0x80, 0x6b, 0x72, 0xcf, 0x12, 0xbb, 0x8b, 0x5f, 0x69, 0x89, 0x29, 0x2d,
	0x81, 0x16, 0x3d, 0x7e, 0x87, 0x1e, 0xfe, 0xc3, 0x43, 0xa3, 0xf7, 0x25, 0x50, 0x0d, 0xf8, 0x83,
	0xf5, 0x70, 0x08, 0xd6, 0x14, 0x87, 0xbc, 0x6a, 0x4f, 0x70, 0x87, 0x2c, 0x4d, 0xed, 0x89, 0x77,
	0x80, 0x7f, 0xf1, 0xd0, 0xe0, 0x43, 0x78, 0x61, 0xb9, 0xd7, 0xc5, 0x67, 0x21, 0xa3, 0x2e, 0x9f,
	0xe8, 0xb9, 0x79, 0x3f, 0x2e, 0xf1, 0x9f, 0x1e, 0x1a, 0x7d, 0x51, 0xa6, 0xff, 0xc7, 0x7c, 0x46,
	0x96, 0xff, 0xeb, 0xc1, 0xfd, 0xe5, 0xfc, 0x4d, 0xb3, 0xa5, 0x54, 0xd3, 0xd0, 0x1e, 0xc4, 0xe4,
	0xb7, 0x42, 0xa3, 0x7a, 0x2e, 0xe1, 0x07, 0xbd, 0xe8, 0xd7, 0x0d, 0xab, 0x60, 0x72, 0x5d, 0x21,
	0xae, 0x46, 0x5d, 0x93, 0xc3, 0x83, 0x95, 0x39, 0xfc, 0x01, 0x0d, 0xcd, 0xf0, 0xc0, 0xf7, 0xae,
	0x0b, 0xe7, 0x0c, 0xaf, 0x80, 0xac, 0x76, 0x32, 0xf3, 0x87, 0xbc, 0x61, 0x51, 0x1f, 0x90, 0xc9,
	0x0a, 0xd4, 0x48, 0xcd, 0x78, 0x62, 0x4e, 0xfd, 0xa3, 0x87, 0x36, 0x9b, 0x37, 0x1e, 0xdf, 0xe9,
	0x05, 0x77, 0x5f, 0xfe, 0x60, 0xa7, 0x67, 0x6a, 0x9f, 0x3f, 0xf2, 0xd0, 0x42, 0xbd, 0x8b, 0x4f,
	0x96, 0x43, 0xcd, 0x4d, 0x85, 0xcb, 0xa8, 0x14, 0xa9, 0x8a, 0x9e, 0x37, 0xa3, 0xe0, 0x32, 0xca,
	0x45, 0xa6, 0x0e, 0xbd, 0x87, 0x87, 0x5f, 0x85, 0xab, 0xfe, 0x4e, 0x16, 0x7f, 0xbd, 0x9e, 0x8e,
	0xec, 0x9f, 0xc8, 0x5b, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x25, 0x83, 0xc7, 0xca, 0x97, 0x09,
	0x00, 0x00,
}
