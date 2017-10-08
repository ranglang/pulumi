// Code generated by protoc-gen-go.
// source: languages.proto
// DO NOT EDIT!

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// RunRequest asks the interpreter to execute a program.
type RunRequest struct {
	Pwd      string            `protobuf:"bytes,1,opt,name=pwd" json:"pwd,omitempty"`
	Program  string            `protobuf:"bytes,2,opt,name=program" json:"program,omitempty"`
	Args     []string          `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	Config   map[string]string `protobuf:"bytes,4,rep,name=config" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DryRun   bool              `protobuf:"varint,5,opt,name=dryRun" json:"dryRun,omitempty"`
	Parallel int32             `protobuf:"varint,6,opt,name=parallel" json:"parallel,omitempty"`
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *RunRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *RunRequest) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *RunRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RunRequest) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *RunRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *RunRequest) GetParallel() int32 {
	if m != nil {
		return m.Parallel
	}
	return 0
}

// RunResponse is the response back from the interpreter/source back to the monitor.
type RunResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *RunResponse) Reset()                    { *m = RunResponse{} }
func (m *RunResponse) String() string            { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()               {}
func (*RunResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RunResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// NewResourceRequest contains information about a resource object that was newly allocated.
type NewResourceRequest struct {
	Type   string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name   string                  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Object *google_protobuf.Struct `protobuf:"bytes,3,opt,name=object" json:"object,omitempty"`
}

func (m *NewResourceRequest) Reset()                    { *m = NewResourceRequest{} }
func (m *NewResourceRequest) String() string            { return proto.CompactTextString(m) }
func (*NewResourceRequest) ProtoMessage()               {}
func (*NewResourceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *NewResourceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NewResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewResourceRequest) GetObject() *google_protobuf.Struct {
	if m != nil {
		return m.Object
	}
	return nil
}

// NewResourceResponse reflects back the properties initialized during creation, if applicable.
type NewResourceResponse struct {
	Id      string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Urn     string                  `protobuf:"bytes,2,opt,name=urn" json:"urn,omitempty"`
	Object  *google_protobuf.Struct `protobuf:"bytes,3,opt,name=object" json:"object,omitempty"`
	Stable  bool                    `protobuf:"varint,4,opt,name=stable" json:"stable,omitempty"`
	Stables []string                `protobuf:"bytes,5,rep,name=stables" json:"stables,omitempty"`
}

func (m *NewResourceResponse) Reset()                    { *m = NewResourceResponse{} }
func (m *NewResourceResponse) String() string            { return proto.CompactTextString(m) }
func (*NewResourceResponse) ProtoMessage()               {}
func (*NewResourceResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *NewResourceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NewResourceResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *NewResourceResponse) GetObject() *google_protobuf.Struct {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *NewResourceResponse) GetStable() bool {
	if m != nil {
		return m.Stable
	}
	return false
}

func (m *NewResourceResponse) GetStables() []string {
	if m != nil {
		return m.Stables
	}
	return nil
}

func init() {
	proto.RegisterType((*RunRequest)(nil), "pulumirpc.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "pulumirpc.RunResponse")
	proto.RegisterType((*NewResourceRequest)(nil), "pulumirpc.NewResourceRequest")
	proto.RegisterType((*NewResourceResponse)(nil), "pulumirpc.NewResourceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LanguageRuntime service

type LanguageRuntimeClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type languageRuntimeClient struct {
	cc *grpc.ClientConn
}

func NewLanguageRuntimeClient(cc *grpc.ClientConn) LanguageRuntimeClient {
	return &languageRuntimeClient{cc}
}

func (c *languageRuntimeClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.LanguageRuntime/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LanguageRuntime service

type LanguageRuntimeServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
}

func RegisterLanguageRuntimeServer(s *grpc.Server, srv LanguageRuntimeServer) {
	s.RegisterService(&_LanguageRuntime_serviceDesc, srv)
}

func _LanguageRuntime_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageRuntime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.LanguageRuntime",
	HandlerType: (*LanguageRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _LanguageRuntime_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "languages.proto",
}

// Client API for ResourceMonitor service

type ResourceMonitorClient interface {
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	NewResource(ctx context.Context, in *NewResourceRequest, opts ...grpc.CallOption) (*NewResourceResponse, error)
}

type resourceMonitorClient struct {
	cc *grpc.ClientConn
}

func NewResourceMonitorClient(cc *grpc.ClientConn) ResourceMonitorClient {
	return &resourceMonitorClient{cc}
}

func (c *resourceMonitorClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) NewResource(ctx context.Context, in *NewResourceRequest, opts ...grpc.CallOption) (*NewResourceResponse, error) {
	out := new(NewResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/NewResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceMonitor service

type ResourceMonitorServer interface {
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	NewResource(context.Context, *NewResourceRequest) (*NewResourceResponse, error)
}

func RegisterResourceMonitorServer(s *grpc.Server, srv ResourceMonitorServer) {
	s.RegisterService(&_ResourceMonitor_serviceDesc, srv)
}

func _ResourceMonitor_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_NewResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).NewResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/NewResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).NewResource(ctx, req.(*NewResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceMonitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.ResourceMonitor",
	HandlerType: (*ResourceMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _ResourceMonitor_Invoke_Handler,
		},
		{
			MethodName: "NewResource",
			Handler:    _ResourceMonitor_NewResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "languages.proto",
}

func init() { proto.RegisterFile("languages.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0x4d, 0x1b, 0xb6, 0x13, 0x69, 0x8b, 0x0c, 0x2c, 0x26, 0x02, 0x14, 0xc2, 0x25, 0xa7,
	0x54, 0x2a, 0x12, 0x62, 0xb9, 0x70, 0x40, 0x1c, 0x56, 0x82, 0x3d, 0x98, 0x5f, 0x90, 0xa6, 0xb3,
	0x51, 0x68, 0x62, 0x07, 0xc7, 0xee, 0x2a, 0x7f, 0x85, 0x23, 0x3f, 0x92, 0x33, 0x8a, 0x3f, 0x96,
	0x02, 0xcb, 0x61, 0x6f, 0xef, 0x8d, 0x9f, 0x67, 0x9e, 0xe7, 0x19, 0x56, 0x6d, 0xc9, 0x6b, 0x5d,
	0xd6, 0x38, 0x14, 0xbd, 0x14, 0x4a, 0x90, 0x65, 0xaf, 0x5b, 0xdd, 0x35, 0xb2, 0xaf, 0x92, 0x67,
	0xb5, 0x10, 0x75, 0x8b, 0x6b, 0x73, 0xb0, 0xd5, 0x57, 0xeb, 0x41, 0x49, 0x5d, 0x29, 0x2b, 0x4c,
	0x4e, 0x7b, 0x29, 0x0e, 0xcd, 0x0e, 0xa5, 0xe5, 0xd9, 0xcf, 0x00, 0x80, 0x69, 0xce, 0xf0, 0x9b,
	0xc6, 0x41, 0x91, 0x07, 0x10, 0xf6, 0xd7, 0x3b, 0x1a, 0xa4, 0x41, 0xbe, 0x64, 0x13, 0x24, 0x14,
	0xee, 0xf7, 0x52, 0xd4, 0xb2, 0xec, 0xe8, 0xcc, 0x54, 0x3d, 0x25, 0x04, 0xe6, 0xa5, 0xac, 0x07,
	0x1a, 0xa6, 0x61, 0xbe, 0x64, 0x06, 0x93, 0x73, 0x88, 0x2a, 0xc1, 0xaf, 0x9a, 0x9a, 0xce, 0xd3,
	0x30, 0x8f, 0x37, 0x2f, 0x8b, 0x1b, 0x63, 0xc5, 0xef, 0x31, 0xc5, 0x07, 0xa3, 0xf9, 0xc8, 0x95,
	0x1c, 0x99, 0xbb, 0x40, 0xce, 0x20, 0xda, 0xc9, 0x91, 0x69, 0x4e, 0x17, 0x69, 0x90, 0x9f, 0x30,
	0xc7, 0x48, 0x02, 0x27, 0x7d, 0x29, 0xcb, 0xb6, 0xc5, 0x96, 0x46, 0x69, 0x90, 0x2f, 0xd8, 0x0d,
	0x4f, 0xce, 0x21, 0x3e, 0x6a, 0x35, 0xb9, 0xdf, 0xe3, 0xe8, 0xdd, 0xef, 0x71, 0x24, 0x8f, 0x60,
	0x71, 0x28, 0x5b, 0x8d, 0xce, 0xbb, 0x25, 0xef, 0x66, 0x6f, 0x83, 0xec, 0x15, 0xc4, 0xc6, 0xd0,
	0xd0, 0x0b, 0x3e, 0xe0, 0x24, 0x44, 0x29, 0x85, 0x74, 0x97, 0x2d, 0xc9, 0x3a, 0x20, 0x97, 0x78,
	0xcd, 0x70, 0x10, 0x5a, 0x56, 0xe8, 0x97, 0x44, 0x60, 0xae, 0xc6, 0x1e, 0x9d, 0xd4, 0xe0, 0xa9,
	0xc6, 0xcb, 0xce, 0xcf, 0x31, 0x98, 0xac, 0x21, 0x12, 0xdb, 0xaf, 0x58, 0x29, 0x1a, 0xa6, 0x41,
	0x1e, 0x6f, 0x9e, 0x14, 0x36, 0x9a, 0xc2, 0x47, 0x53, 0x7c, 0x31, 0xd1, 0x30, 0x27, 0xcb, 0xbe,
	0x07, 0xf0, 0xf0, 0x8f, 0x79, 0xce, 0xdc, 0x29, 0xcc, 0x1a, 0x1f, 0xca, 0xac, 0xd9, 0x4d, 0xef,
	0xd4, 0x92, 0xbb, 0x59, 0x13, 0xbc, 0xf3, 0xa8, 0x69, 0xdb, 0x83, 0x2a, 0xb7, 0x2d, 0xd2, 0xb9,
	0xdd, 0xb6, 0x65, 0x53, 0xdc, 0x16, 0x0d, 0x74, 0x61, 0x72, 0xf5, 0x74, 0x73, 0x01, 0xab, 0x4f,
	0xee, 0xd7, 0x31, 0xcd, 0x55, 0xd3, 0x21, 0x79, 0x03, 0xe1, 0x94, 0xd0, 0xe3, 0x5b, 0x43, 0x4e,
	0xce, 0xfe, 0x2e, 0xdb, 0xd7, 0x64, 0xf7, 0x36, 0x3f, 0x02, 0x58, 0xf9, 0x47, 0x7e, 0x16, 0xbc,
	0x51, 0x42, 0x92, 0xf7, 0x10, 0x5d, 0xf0, 0x83, 0xd8, 0x23, 0xa1, 0x47, 0xf7, 0x6c, 0xc9, 0x77,
	0x7c, 0x7a, 0xcb, 0x89, 0x6f, 0x4a, 0x2e, 0x21, 0x3e, 0xda, 0x1d, 0x79, 0x7e, 0xa4, 0xfd, 0x37,
	0xc3, 0xe4, 0xc5, 0xff, 0x8e, 0x7d, 0xbf, 0x6d, 0x64, 0x56, 0xf7, 0xfa, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0xbc, 0x47, 0x30, 0x6c, 0x03, 0x00, 0x00,
}