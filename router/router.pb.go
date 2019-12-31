// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package router

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetParameter struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParameter) Reset()         { *m = GetParameter{} }
func (m *GetParameter) String() string { return proto.CompactTextString(m) }
func (*GetParameter) ProtoMessage()    {}
func (*GetParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *GetParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameter.Unmarshal(m, b)
}
func (m *GetParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameter.Marshal(b, m, deterministic)
}
func (m *GetParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameter.Merge(m, src)
}
func (m *GetParameter) XXX_Size() int {
	return xxx_messageInfo_GetParameter.Size(m)
}
func (m *GetParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameter.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameter proto.InternalMessageInfo

func (m *GetParameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResult struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResult) Reset()         { *m = GetResult{} }
func (m *GetResult) String() string { return proto.CompactTextString(m) }
func (*GetResult) ProtoMessage()    {}
func (*GetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *GetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResult.Unmarshal(m, b)
}
func (m *GetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResult.Marshal(b, m, deterministic)
}
func (m *GetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResult.Merge(m, src)
}
func (m *GetResult) XXX_Size() int {
	return xxx_messageInfo_GetResult.Size(m)
}
func (m *GetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetResult proto.InternalMessageInfo

func (m *GetResult) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetParameter struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetParameter) Reset()         { *m = SetParameter{} }
func (m *SetParameter) String() string { return proto.CompactTextString(m) }
func (*SetParameter) ProtoMessage()    {}
func (*SetParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *SetParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetParameter.Unmarshal(m, b)
}
func (m *SetParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetParameter.Marshal(b, m, deterministic)
}
func (m *SetParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetParameter.Merge(m, src)
}
func (m *SetParameter) XXX_Size() int {
	return xxx_messageInfo_SetParameter.Size(m)
}
func (m *SetParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_SetParameter.DiscardUnknown(m)
}

var xxx_messageInfo_SetParameter proto.InternalMessageInfo

func (m *SetParameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetParameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResult) Reset()         { *m = SetResult{} }
func (m *SetResult) String() string { return proto.CompactTextString(m) }
func (*SetResult) ProtoMessage()    {}
func (*SetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *SetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResult.Unmarshal(m, b)
}
func (m *SetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResult.Marshal(b, m, deterministic)
}
func (m *SetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResult.Merge(m, src)
}
func (m *SetResult) XXX_Size() int {
	return xxx_messageInfo_SetResult.Size(m)
}
func (m *SetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResult.DiscardUnknown(m)
}

var xxx_messageInfo_SetResult proto.InternalMessageInfo

func (m *SetResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*GetParameter)(nil), "GetParameter")
	proto.RegisterType((*GetResult)(nil), "GetResult")
	proto.RegisterType((*SetParameter)(nil), "SetParameter")
	proto.RegisterType((*SetResult)(nil), "SetResult")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe0, 0xe2, 0x71, 0x4f, 0x2d, 0x09,
	0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x14, 0xb9, 0x38, 0xdd, 0x53, 0x4b, 0x82,
	0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0xa1, 0x0a,
	0x20, 0x1c, 0x25, 0x33, 0x2e, 0x9e, 0x60, 0xbc, 0x86, 0x20, 0xf4, 0x31, 0x21, 0xeb, 0x53, 0xe6,
	0xe2, 0x0c, 0x86, 0x1b, 0x2d, 0xc6, 0xc5, 0x56, 0x04, 0x66, 0x81, 0xf5, 0x71, 0x04, 0x41, 0x79,
	0x46, 0x01, 0x5c, 0x6c, 0x41, 0x60, 0x17, 0x0b, 0x29, 0x71, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0xf1,
	0xea, 0x21, 0xbb, 0x58, 0x8a, 0x4b, 0x0f, 0xee, 0x3c, 0x25, 0x06, 0x90, 0x9a, 0x60, 0xb0, 0x9a,
	0x60, 0x54, 0x35, 0xc1, 0x08, 0x35, 0x49, 0x6c, 0x60, 0xaf, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x47, 0xf0, 0x38, 0x0a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Get(ctx context.Context, in *GetParameter, opts ...grpc.CallOption) (*GetResult, error)
	Set(ctx context.Context, in *SetParameter, opts ...grpc.CallOption) (*SetResult, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Get(ctx context.Context, in *GetParameter, opts ...grpc.CallOption) (*GetResult, error) {
	out := new(GetResult)
	err := c.cc.Invoke(ctx, "/Router/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Set(ctx context.Context, in *SetParameter, opts ...grpc.CallOption) (*SetResult, error) {
	out := new(SetResult)
	err := c.cc.Invoke(ctx, "/Router/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Get(context.Context, *GetParameter) (*GetResult, error)
	Set(context.Context, *SetParameter) (*SetResult, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Get(ctx context.Context, req *GetParameter) (*GetResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRouterServer) Set(ctx context.Context, req *SetParameter) (*SetResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Get(ctx, req.(*GetParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Set(ctx, req.(*SetParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Router_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Router_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
