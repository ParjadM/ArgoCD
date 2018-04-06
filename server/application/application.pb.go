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
*/
package application

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "k8s.io/api/core/v1"
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

func init() {
	proto.RegisterType((*ApplicationQuery)(nil), "application.ApplicationQuery")
	proto.RegisterType((*ApplicationResponse)(nil), "application.ApplicationResponse")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "application.DeleteApplicationRequest")
	proto.RegisterType((*ApplicationSyncRequest)(nil), "application.ApplicationSyncRequest")
	proto.RegisterType((*ApplicationSyncResult)(nil), "application.ApplicationSyncResult")
	proto.RegisterType((*ResourceDetails)(nil), "application.ResourceDetails")
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
	},
	Metadata: "server/application/application.proto",
}

func init() { proto.RegisterFile("server/application/application.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x9b, 0xfc, 0xf2, 0x6b, 0xb6, 0x07, 0xd0, 0xd2, 0x46, 0xc6, 0xa4, 0x22, 0xda, 0xb6,
	0xa8, 0x04, 0x61, 0x37, 0xe5, 0x82, 0x72, 0x03, 0x0a, 0x08, 0xc4, 0x01, 0x8c, 0x10, 0x12, 0x17,
	0xd8, 0xda, 0x23, 0xc7, 0xc4, 0xde, 0x75, 0x77, 0xd7, 0x96, 0x22, 0xe0, 0xc2, 0x0b, 0x20, 0xc4,
	0x03, 0xf0, 0x1c, 0x3c, 0x02, 0x67, 0x6e, 0x9c, 0x79, 0x10, 0xe4, 0x8d, 0xd3, 0xd8, 0xf9, 0xd7,
	0x4b, 0x0e, 0x9c, 0x32, 0xbb, 0x33, 0x99, 0xef, 0x9b, 0xcf, 0x9f, 0x66, 0xd1, 0xbe, 0x04, 0x91,
	0x81, 0x70, 0x68, 0x92, 0x44, 0xa1, 0x47, 0x55, 0xc8, 0x59, 0x39, 0xb6, 0x13, 0xc1, 0x15, 0xc7,
	0x5b, 0xa5, 0x2b, 0x6b, 0x3b, 0xe0, 0x01, 0xd7, 0xf7, 0x4e, 0x1e, 0x8d, 0x4b, 0xac, 0x76, 0xc0,
	0x79, 0x10, 0x81, 0x43, 0x93, 0xd0, 0xa1, 0x8c, 0x71, 0xa5, 0x8b, 0x65, 0x91, 0x25, 0xc3, 0xbb,
	0xd2, 0x0e, 0xb9, 0xce, 0x7a, 0x5c, 0x80, 0x93, 0xf5, 0x9c, 0x00, 0x18, 0x08, 0xaa, 0xc0, 0x2f,
	0x6a, 0x9e, 0x04, 0xa1, 0x1a, 0xa4, 0xa7, 0xb6, 0xc7, 0x63, 0x87, 0x0a, 0x0d, 0xf1, 0x5e, 0x07,
	0xb7, 0x3d, 0xdf, 0x49, 0x86, 0x41, 0xfe, 0x67, 0x59, 0x21, 0x9a, 0xf5, 0x68, 0x94, 0x0c, 0xe8,
	0x5c, 0x2b, 0x72, 0x03, 0x5d, 0xbe, 0x37, 0xad, 0x7b, 0x91, 0x82, 0x18, 0x61, 0x8c, 0xea, 0x8c,
	0xc6, 0x60, 0x1a, 0x1d, 0xe3, 0xb0, 0xe9, 0xea, 0x98, 0xec, 0xa0, 0x2b, 0xa5, 0x3a, 0x17, 0x64,
	0xc2, 0x99, 0x04, 0xe2, 0x23, 0xf3, 0x04, 0x22, 0x50, 0x50, 0x49, 0x9e, 0xa5, 0x20, 0xd5, 0xa2,
	0x36, 0xb8, 0x8d, 0x9a, 0xf9, 0xaf, 0x4c, 0xa8, 0x07, 0xe6, 0x86, 0x4e, 0x4c, 0x2f, 0x70, 0x0b,
	0x35, 0xc6, 0x22, 0x9b, 0x35, 0x9d, 0x2a, 0x4e, 0xe4, 0x1d, 0x6a, 0x95, 0xfa, 0xbf, 0x1c, 0x31,
	0x6f, 0x15, 0x86, 0x85, 0x36, 0x05, 0x64, 0xa1, 0x0c, 0x39, 0x2b, 0x20, 0xce, 0xcf, 0x39, 0x82,
	0x2f, 0x46, 0x6e, 0xca, 0x34, 0xc2, 0xa6, 0x5b, 0x9c, 0x48, 0x8c, 0x76, 0xe6, 0x10, 0x64, 0x1a,
	0x29, 0x6c, 0xa2, 0xff, 0x63, 0x90, 0x92, 0x06, 0x13, 0x8c, 0xc9, 0x11, 0xf7, 0x51, 0x53, 0x80,
	0xe4, 0xa9, 0xf0, 0x40, 0x9a, 0x1b, 0x9d, 0xda, 0xe1, 0xd6, 0x71, 0xdb, 0x2e, 0x1b, 0xc2, 0x2d,
	0xb2, 0x27, 0xa0, 0x68, 0x18, 0x49, 0x77, 0x5a, 0x4e, 0xce, 0xd0, 0xa5, 0x99, 0xec, 0xc2, 0x49,
	0x30, 0xaa, 0x0f, 0x43, 0xe6, 0x17, 0x53, 0xe8, 0xb8, 0xaa, 0x60, 0x6d, 0x56, 0xc1, 0x12, 0xdd,
	0x7a, 0x85, 0xee, 0xf1, 0xef, 0x4d, 0x84, 0xcb, 0x23, 0x82, 0xc8, 0x42, 0x0f, 0xf0, 0x17, 0x03,
	0xd5, 0x9f, 0x85, 0x52, 0xe1, 0xdd, 0x0a, 0xf7, 0x59, 0x4f, 0x58, 0x4f, 0xed, 0xa9, 0xe7, 0xec,
	0x89, 0xe7, 0x74, 0xf0, 0xd6, 0xf3, 0xed, 0x64, 0x18, 0xd8, 0xb9, 0xe7, 0x2a, 0x3d, 0x26, 0x9e,
	0x2b, 0x37, 0xcb, 0xa1, 0x48, 0xfb, 0xf3, 0xaf, 0x3f, 0xdf, 0x36, 0x5a, 0x78, 0x5b, 0x9b, 0x3c,
	0xeb, 0x95, 0x9d, 0x2a, 0xf1, 0x77, 0x03, 0xfd, 0xf7, 0x9a, 0x2a, 0x6f, 0x70, 0x11, 0xa5, 0xe7,
	0xeb, 0xa1, 0xa4, 0xb1, 0x1e, 0x66, 0xc0, 0x14, 0xd9, 0xd3, 0xc4, 0x76, 0xf1, 0xb5, 0x09, 0x31,
	0xa9, 0x04, 0xd0, 0xb8, 0xc2, 0xef, 0xc8, 0xc0, 0x3f, 0x0c, 0xd4, 0x78, 0x20, 0x80, 0x2a, 0xc0,
	0x8f, 0xd6, 0xc3, 0xc1, 0x5a, 0x53, 0x1f, 0x72, 0x5d, 0x4f, 0x70, 0x95, 0x2c, 0x94, 0xb6, 0x6f,
	0x74, 0xf1, 0x57, 0x03, 0xd5, 0x1e, 0xc3, 0x85, 0x9f, 0x7b, 0x5d, 0x7c, 0xe6, 0x14, 0x2d, 0xf3,
	0x71, 0x3e, 0xe4, 0xc6, 0xfd, 0x84, 0x7f, 0x1a, 0xa8, 0xf1, 0x2a, 0xf1, 0xff, 0x45, 0x3d, 0x1d,
	0xcd, 0xff, 0xa6, 0xb5, 0xbf, 0x98, 0x7f, 0x0c, 0x8a, 0xfa, 0x54, 0x51, 0x5b, 0x0f, 0x92, 0xeb,
	0x9b, 0xa1, 0xc6, 0x78, 0x21, 0xe2, 0x83, 0x4a, 0xf7, 0x65, 0x5b, 0xd2, 0xea, 0x2c, 0xfb, 0x10,
	0xe7, 0x3b, 0xb6, 0xd0, 0xb0, 0xbb, 0x52, 0xc3, 0x8f, 0xa8, 0x9e, 0x6f, 0x2d, 0xbc, 0xb7, 0xac,
	0x5d, 0x69, 0x6b, 0x5a, 0x64, 0x75, 0x51, 0xbe, 0xf8, 0xc8, 0x2d, 0x8d, 0x7a, 0x40, 0x3a, 0x2b,
	0x50, 0x1d, 0x39, 0x62, 0x5e, 0xdf, 0xe8, 0xde, 0x3f, 0x7a, 0x63, 0xaf, 0x7a, 0x92, 0xe6, 0x5f,
	0xce, 0xd3, 0x86, 0x7e, 0x7e, 0xee, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x64, 0x8b, 0x57, 0xe4,
	0x56, 0x07, 0x00, 0x00,
}
