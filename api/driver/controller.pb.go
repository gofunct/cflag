// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller.proto

package driver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_controller_6ee40fdde8d469f8, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "driver.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GRPCControllerClient is the client API for GRPCController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCControllerClient interface {
	Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type gRPCControllerClient struct {
	cc *grpc.ClientConn
}

func NewGRPCControllerClient(cc *grpc.ClientConn) GRPCControllerClient {
	return &gRPCControllerClient{cc}
}

func (c *gRPCControllerClient) Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/driver.GRPCController/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCControllerServer is the server API for GRPCController service.
type GRPCControllerServer interface {
	Shutdown(context.Context, *Empty) (*Empty, error)
}

func RegisterGRPCControllerServer(s *grpc.Server, srv GRPCControllerServer) {
	s.RegisterService(&_GRPCController_serviceDesc, srv)
}

func _GRPCController_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCControllerServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.GRPCController/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCControllerServer).Shutdown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "driver.GRPCController",
	HandlerType: (*GRPCControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shutdown",
			Handler:    _GRPCController_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}

func init() { proto.RegisterFile("controller.proto", fileDescriptor_controller_6ee40fdde8d469f8) }

var fileDescriptor_controller_6ee40fdde8d469f8 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x29,
	0xca, 0x2c, 0x4b, 0x2d, 0x52, 0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x34, 0xb2, 0xe2,
	0xe2, 0x73, 0x0f, 0x0a, 0x70, 0x76, 0x86, 0x2b, 0x14, 0xd2, 0xe0, 0xe2, 0x08, 0xce, 0x28, 0x2d,
	0x49, 0xc9, 0x2f, 0xcf, 0x13, 0xe2, 0xd5, 0x83, 0xa8, 0xd7, 0x03, 0x2b, 0x96, 0x42, 0xe5, 0x26,
	0xb1, 0x81, 0xcd, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x49, 0x69, 0x4f, 0xbe, 0x67, 0x00,
	0x00, 0x00,
}