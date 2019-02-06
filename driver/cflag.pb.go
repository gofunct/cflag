// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cflag.proto

package driver

/*
 type = "double" | "float" | "int32" | "int64" | "uint32" | "uint64"
     | "sint32" | "sint64" | "fixed32" | "fixed64" | "sfixed32" | "sfixed64"
     | "bool" | "string" | "bytes" | messageType | enumType
maps: map<string, Project> projects = 3;
*/

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

type Command int32

const (
	Command_TERRAFORM Command = 0
	Command_KUBECTL   Command = 1
	Command_GCLOUD    Command = 2
)

var Command_name = map[int32]string{
	0: "TERRAFORM",
	1: "KUBECTL",
	2: "GCLOUD",
}
var Command_value = map[string]int32{
	"TERRAFORM": 0,
	"KUBECTL":   1,
	"GCLOUD":    2,
}

func (x Command) String() string {
	return proto.EnumName(Command_name, int32(x))
}
func (Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{0}
}

type Mode int32

const (
	Mode_DEV  Mode = 0
	Mode_PROD Mode = 1
)

var Mode_name = map[int32]string{
	0: "DEV",
	1: "PROD",
}
var Mode_value = map[string]int32{
	"DEV":  0,
	"PROD": 1,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}
func (Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{1}
}

type Role int32

const (
	Role_GUEST  Role = 0
	Role_MEMBER Role = 1
	Role_ADMIN  Role = 2
)

var Role_name = map[int32]string{
	0: "GUEST",
	1: "MEMBER",
	2: "ADMIN",
}
var Role_value = map[string]int32{
	"GUEST":  0,
	"MEMBER": 1,
	"ADMIN":  2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}
func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{2}
}

type Script struct {
	Command              Command           `protobuf:"varint,1,opt,name=command,proto3,enum=driver.Command" json:"command,omitempty"`
	Args                 []string          `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Env                  map[string]string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Usage                *Usage            `protobuf:"bytes,4,opt,name=usage,proto3" json:"usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Script) Reset()         { *m = Script{} }
func (m *Script) String() string { return proto.CompactTextString(m) }
func (*Script) ProtoMessage()    {}
func (*Script) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{0}
}
func (m *Script) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Script.Unmarshal(m, b)
}
func (m *Script) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Script.Marshal(b, m, deterministic)
}
func (dst *Script) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Script.Merge(dst, src)
}
func (m *Script) XXX_Size() int {
	return xxx_messageInfo_Script.Size(m)
}
func (m *Script) XXX_DiscardUnknown() {
	xxx_messageInfo_Script.DiscardUnknown(m)
}

var xxx_messageInfo_Script proto.InternalMessageInfo

func (m *Script) GetCommand() Command {
	if m != nil {
		return m.Command
	}
	return Command_TERRAFORM
}

func (m *Script) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Script) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Script) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

type Usage struct {
	Usage                map[string]string `protobuf:"bytes,1,rep,name=usage,proto3" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Usage) Reset()         { *m = Usage{} }
func (m *Usage) String() string { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()    {}
func (*Usage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{1}
}
func (m *Usage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Usage.Unmarshal(m, b)
}
func (m *Usage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Usage.Marshal(b, m, deterministic)
}
func (dst *Usage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Usage.Merge(dst, src)
}
func (m *Usage) XXX_Size() int {
	return xxx_messageInfo_Usage.Size(m)
}
func (m *Usage) XXX_DiscardUnknown() {
	xxx_messageInfo_Usage.DiscardUnknown(m)
}

var xxx_messageInfo_Usage proto.InternalMessageInfo

func (m *Usage) GetUsage() map[string]string {
	if m != nil {
		return m.Usage
	}
	return nil
}

type ExecRequest struct {
	User                 *User     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Mode                 Mode      `protobuf:"varint,2,opt,name=mode,proto3,enum=driver.Mode" json:"mode,omitempty"`
	Scripts              []*Script `protobuf:"bytes,3,rep,name=scripts,proto3" json:"scripts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExecRequest) Reset()         { *m = ExecRequest{} }
func (m *ExecRequest) String() string { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()    {}
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{2}
}
func (m *ExecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecRequest.Unmarshal(m, b)
}
func (m *ExecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecRequest.Marshal(b, m, deterministic)
}
func (dst *ExecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecRequest.Merge(dst, src)
}
func (m *ExecRequest) XXX_Size() int {
	return xxx_messageInfo_ExecRequest.Size(m)
}
func (m *ExecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecRequest proto.InternalMessageInfo

func (m *ExecRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ExecRequest) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_DEV
}

func (m *ExecRequest) GetScripts() []*Script {
	if m != nil {
		return m.Scripts
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Role                 Role     `protobuf:"varint,6,opt,name=role,proto3,enum=driver.Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{3}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

type Info struct {
	Text                 []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{4}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

type Bytes struct {
	Resp                 []byte   `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{5}
}
func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (dst *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(dst, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetResp() []byte {
	if m != nil {
		return m.Resp
	}
	return nil
}

type DebugRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugRequest) Reset()         { *m = DebugRequest{} }
func (m *DebugRequest) String() string { return proto.CompactTextString(m) }
func (*DebugRequest) ProtoMessage()    {}
func (*DebugRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{6}
}
func (m *DebugRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugRequest.Unmarshal(m, b)
}
func (m *DebugRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugRequest.Marshal(b, m, deterministic)
}
func (dst *DebugRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugRequest.Merge(dst, src)
}
func (m *DebugRequest) XXX_Size() int {
	return xxx_messageInfo_DebugRequest.Size(m)
}
func (m *DebugRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DebugRequest proto.InternalMessageInfo

func (m *DebugRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DebugResponse struct {
	Usage                []*Usage `protobuf:"bytes,1,rep,name=usage,proto3" json:"usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugResponse) Reset()         { *m = DebugResponse{} }
func (m *DebugResponse) String() string { return proto.CompactTextString(m) }
func (*DebugResponse) ProtoMessage()    {}
func (*DebugResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cflag_c8cb5a99b882f0da, []int{7}
}
func (m *DebugResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugResponse.Unmarshal(m, b)
}
func (m *DebugResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugResponse.Marshal(b, m, deterministic)
}
func (dst *DebugResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugResponse.Merge(dst, src)
}
func (m *DebugResponse) XXX_Size() int {
	return xxx_messageInfo_DebugResponse.Size(m)
}
func (m *DebugResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DebugResponse proto.InternalMessageInfo

func (m *DebugResponse) GetUsage() []*Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func init() {
	proto.RegisterType((*Script)(nil), "driver.Script")
	proto.RegisterMapType((map[string]string)(nil), "driver.Script.EnvEntry")
	proto.RegisterType((*Usage)(nil), "driver.Usage")
	proto.RegisterMapType((map[string]string)(nil), "driver.Usage.UsageEntry")
	proto.RegisterType((*ExecRequest)(nil), "driver.ExecRequest")
	proto.RegisterType((*User)(nil), "driver.User")
	proto.RegisterType((*Info)(nil), "driver.Info")
	proto.RegisterType((*Bytes)(nil), "driver.Bytes")
	proto.RegisterType((*DebugRequest)(nil), "driver.DebugRequest")
	proto.RegisterType((*DebugResponse)(nil), "driver.DebugResponse")
	proto.RegisterEnum("driver.Command", Command_name, Command_value)
	proto.RegisterEnum("driver.Mode", Mode_name, Mode_value)
	proto.RegisterEnum("driver.Role", Role_name, Role_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoCloudClient is the client API for GoCloud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoCloudClient interface {
	Signup(ctx context.Context, in *User, opts ...grpc.CallOption) (GoCloud_SignupClient, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (GoCloud_LoginClient, error)
	Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (GoCloud_DebugClient, error)
	Execute(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (GoCloud_ExecuteClient, error)
}

type goCloudClient struct {
	cc *grpc.ClientConn
}

func NewGoCloudClient(cc *grpc.ClientConn) GoCloudClient {
	return &goCloudClient{cc}
}

func (c *goCloudClient) Signup(ctx context.Context, in *User, opts ...grpc.CallOption) (GoCloud_SignupClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GoCloud_serviceDesc.Streams[0], "/driver.GoCloud/Signup", opts...)
	if err != nil {
		return nil, err
	}
	x := &goCloudSignupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoCloud_SignupClient interface {
	Recv() (*Info, error)
	grpc.ClientStream
}

type goCloudSignupClient struct {
	grpc.ClientStream
}

func (x *goCloudSignupClient) Recv() (*Info, error) {
	m := new(Info)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goCloudClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (GoCloud_LoginClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GoCloud_serviceDesc.Streams[1], "/driver.GoCloud/Login", opts...)
	if err != nil {
		return nil, err
	}
	x := &goCloudLoginClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoCloud_LoginClient interface {
	Recv() (*Info, error)
	grpc.ClientStream
}

type goCloudLoginClient struct {
	grpc.ClientStream
}

func (x *goCloudLoginClient) Recv() (*Info, error) {
	m := new(Info)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goCloudClient) Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (GoCloud_DebugClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GoCloud_serviceDesc.Streams[2], "/driver.GoCloud/Debug", opts...)
	if err != nil {
		return nil, err
	}
	x := &goCloudDebugClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoCloud_DebugClient interface {
	Recv() (*DebugResponse, error)
	grpc.ClientStream
}

type goCloudDebugClient struct {
	grpc.ClientStream
}

func (x *goCloudDebugClient) Recv() (*DebugResponse, error) {
	m := new(DebugResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goCloudClient) Execute(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (GoCloud_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GoCloud_serviceDesc.Streams[3], "/driver.GoCloud/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &goCloudExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoCloud_ExecuteClient interface {
	Recv() (*Bytes, error)
	grpc.ClientStream
}

type goCloudExecuteClient struct {
	grpc.ClientStream
}

func (x *goCloudExecuteClient) Recv() (*Bytes, error) {
	m := new(Bytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoCloudServer is the server API for GoCloud service.
type GoCloudServer interface {
	Signup(*User, GoCloud_SignupServer) error
	Login(*User, GoCloud_LoginServer) error
	Debug(*DebugRequest, GoCloud_DebugServer) error
	Execute(*ExecRequest, GoCloud_ExecuteServer) error
}

func RegisterGoCloudServer(s *grpc.Server, srv GoCloudServer) {
	s.RegisterService(&_GoCloud_serviceDesc, srv)
}

func _GoCloud_Signup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoCloudServer).Signup(m, &goCloudSignupServer{stream})
}

type GoCloud_SignupServer interface {
	Send(*Info) error
	grpc.ServerStream
}

type goCloudSignupServer struct {
	grpc.ServerStream
}

func (x *goCloudSignupServer) Send(m *Info) error {
	return x.ServerStream.SendMsg(m)
}

func _GoCloud_Login_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoCloudServer).Login(m, &goCloudLoginServer{stream})
}

type GoCloud_LoginServer interface {
	Send(*Info) error
	grpc.ServerStream
}

type goCloudLoginServer struct {
	grpc.ServerStream
}

func (x *goCloudLoginServer) Send(m *Info) error {
	return x.ServerStream.SendMsg(m)
}

func _GoCloud_Debug_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DebugRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoCloudServer).Debug(m, &goCloudDebugServer{stream})
}

type GoCloud_DebugServer interface {
	Send(*DebugResponse) error
	grpc.ServerStream
}

type goCloudDebugServer struct {
	grpc.ServerStream
}

func (x *goCloudDebugServer) Send(m *DebugResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GoCloud_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoCloudServer).Execute(m, &goCloudExecuteServer{stream})
}

type GoCloud_ExecuteServer interface {
	Send(*Bytes) error
	grpc.ServerStream
}

type goCloudExecuteServer struct {
	grpc.ServerStream
}

func (x *goCloudExecuteServer) Send(m *Bytes) error {
	return x.ServerStream.SendMsg(m)
}

var _GoCloud_serviceDesc = grpc.ServiceDesc{
	ServiceName: "driver.GoCloud",
	HandlerType: (*GoCloudServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signup",
			Handler:       _GoCloud_Signup_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Login",
			Handler:       _GoCloud_Login_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Debug",
			Handler:       _GoCloud_Debug_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Execute",
			Handler:       _GoCloud_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cflag.proto",
}

func init() { proto.RegisterFile("cflag.proto", fileDescriptor_cflag_c8cb5a99b882f0da) }

var fileDescriptor_cflag_c8cb5a99b882f0da = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0xfa, 0x27, 0x6e, 0x26, 0x6d, 0xb0, 0x86, 0x22, 0x4c, 0xb8, 0x44, 0xe6, 0x80, 0xdb,
	0x43, 0xd4, 0x06, 0x84, 0x2a, 0x6e, 0x6d, 0x62, 0xaa, 0x8a, 0x86, 0xa2, 0x6d, 0xc3, 0xdd, 0x8d,
	0xb7, 0xc1, 0xc2, 0xf1, 0xba, 0x5e, 0x3b, 0xb4, 0xe2, 0x35, 0x78, 0x29, 0x4e, 0xbc, 0x12, 0xda,
	0xdd, 0x38, 0x69, 0x10, 0x12, 0x70, 0xb1, 0x76, 0x7e, 0xbe, 0x99, 0x6f, 0xbe, 0x19, 0x19, 0xda,
	0xd3, 0x9b, 0x34, 0x9a, 0xf5, 0xf3, 0x82, 0x97, 0x1c, 0x9b, 0x71, 0x91, 0x2c, 0x58, 0xe1, 0xff,
	0x24, 0xd0, 0xbc, 0x9c, 0x16, 0x49, 0x5e, 0xe2, 0x1e, 0x38, 0x53, 0x3e, 0x9f, 0x47, 0x59, 0xec,
	0x91, 0x1e, 0x09, 0x3a, 0x83, 0x47, 0x7d, 0x9d, 0xd4, 0x1f, 0x6a, 0x37, 0xad, 0xe3, 0x88, 0x60,
	0x45, 0xc5, 0x4c, 0x78, 0x46, 0xcf, 0x0c, 0x5a, 0x54, 0xbd, 0x71, 0x0f, 0x4c, 0x96, 0x2d, 0x3c,
	0xb3, 0x67, 0x06, 0xed, 0xc1, 0xd3, 0x1a, 0xaa, 0x6b, 0xf7, 0xc3, 0x6c, 0x11, 0x66, 0x65, 0x71,
	0x4f, 0x65, 0x0e, 0xbe, 0x00, 0xbb, 0x12, 0xd1, 0x8c, 0x79, 0x56, 0x8f, 0x04, 0xed, 0xc1, 0x4e,
	0x9d, 0x3c, 0x91, 0x4e, 0xaa, 0x63, 0xdd, 0x37, 0xb0, 0x55, 0xa3, 0xd0, 0x05, 0xf3, 0x0b, 0xbb,
	0x57, 0xb4, 0x5a, 0x54, 0x3e, 0x71, 0x17, 0xec, 0x45, 0x94, 0x56, 0xcc, 0x33, 0x94, 0x4f, 0x1b,
	0x6f, 0x8d, 0x23, 0xe2, 0xdf, 0x82, 0xad, 0xea, 0x60, 0xbf, 0xee, 0x42, 0x14, 0x25, 0x6f, 0xa3,
	0x8b, 0xfe, 0x6a, 0x4e, 0xcb, 0x86, 0x47, 0x00, 0x6b, 0xe7, 0x7f, 0xb5, 0xfc, 0x06, 0xed, 0xf0,
	0x8e, 0x4d, 0x29, 0xbb, 0xad, 0x98, 0x28, 0xb1, 0x07, 0x56, 0x25, 0x58, 0xa1, 0xb0, 0xed, 0xc1,
	0xf6, 0xba, 0x2f, 0x2b, 0xa8, 0x8a, 0xc8, 0x8c, 0x39, 0x8f, 0x75, 0xa5, 0xce, 0x3a, 0x63, 0xcc,
	0x63, 0x46, 0x55, 0x04, 0x03, 0x70, 0x84, 0x92, 0x4e, 0x2c, 0x15, 0xed, 0x6c, 0x2a, 0x4a, 0xeb,
	0xb0, 0xff, 0x9d, 0x80, 0x25, 0x4b, 0x63, 0x07, 0x8c, 0x24, 0x5e, 0x12, 0x36, 0x12, 0xb5, 0x24,
	0x45, 0x43, 0xd3, 0xd5, 0x8d, 0x77, 0xc1, 0x66, 0xf3, 0x28, 0x49, 0x3d, 0x53, 0xcf, 0xa0, 0x0c,
	0xec, 0xc2, 0x56, 0x1e, 0x09, 0xf1, 0x95, 0x17, 0xb1, 0x5a, 0x49, 0x8b, 0xae, 0x6c, 0x89, 0xc8,
	0x3f, 0xf3, 0x8c, 0x79, 0xb6, 0x46, 0x28, 0x43, 0x0e, 0x50, 0xf0, 0x94, 0x79, 0xcd, 0xcd, 0x01,
	0x28, 0x4f, 0x19, 0x55, 0x11, 0xbf, 0x0b, 0xd6, 0x59, 0x76, 0xc3, 0x25, 0x8b, 0x92, 0xdd, 0x95,
	0x6a, 0x09, 0x2d, 0xaa, 0xde, 0xfe, 0x73, 0xb0, 0x4f, 0xee, 0x4b, 0x26, 0x64, 0xb0, 0x60, 0x22,
	0x57, 0xa4, 0xb7, 0xa9, 0x7a, 0xfb, 0x07, 0xb0, 0x3d, 0x62, 0xd7, 0xd5, 0xec, 0x9f, 0xd5, 0xf4,
	0x5f, 0xc3, 0xce, 0x12, 0x21, 0x72, 0x9e, 0x09, 0xb6, 0xbe, 0x2f, 0xbd, 0xf9, 0x3f, 0xde, 0xd7,
	0xfe, 0x21, 0x38, 0xcb, 0xbb, 0xc6, 0x1d, 0x68, 0x5d, 0x85, 0x94, 0x1e, 0xbf, 0xbb, 0xa0, 0x63,
	0xb7, 0x81, 0x6d, 0x70, 0xde, 0x4f, 0x4e, 0xc2, 0xe1, 0xd5, 0xb9, 0x4b, 0x10, 0xa0, 0x79, 0x3a,
	0x3c, 0xbf, 0x98, 0x8c, 0x5c, 0x63, 0xff, 0x19, 0x58, 0x72, 0x45, 0xe8, 0x80, 0x39, 0x0a, 0x3f,
	0xb9, 0x0d, 0xdc, 0x02, 0xeb, 0x23, 0xbd, 0x18, 0xb9, 0x64, 0x3f, 0x00, 0x4b, 0x0e, 0x8f, 0x2d,
	0xb0, 0x4f, 0x27, 0xe1, 0xe5, 0x95, 0xdb, 0x90, 0xc8, 0x71, 0x38, 0x3e, 0x09, 0xa9, 0x4b, 0xa4,
	0xfb, 0x78, 0x34, 0x3e, 0xfb, 0xe0, 0x1a, 0x83, 0x1f, 0x04, 0x9c, 0x53, 0x3e, 0x4c, 0x79, 0x15,
	0x63, 0x00, 0xcd, 0xcb, 0x64, 0x96, 0x55, 0x39, 0x6e, 0xcc, 0xd5, 0x5d, 0x59, 0x52, 0x42, 0xbf,
	0x71, 0x40, 0xf0, 0x25, 0xd8, 0xe7, 0x7c, 0x96, 0x64, 0x7f, 0x4d, 0x3c, 0x02, 0x5b, 0x89, 0x81,
	0xbb, 0x75, 0xe8, 0xa1, 0x9a, 0xdd, 0x27, 0xbf, 0x79, 0xb5, 0x62, 0x0a, 0x79, 0x08, 0x8e, 0xbc,
	0xe2, 0xaa, 0x64, 0xf8, 0xb8, 0xce, 0x7a, 0x70, 0xd6, 0xdd, 0x95, 0x8c, 0x6a, 0x77, 0x12, 0x72,
	0xdd, 0x54, 0x3f, 0x93, 0x57, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x85, 0x8e, 0x0a, 0x55, 0x5b,
	0x04, 0x00, 0x00,
}