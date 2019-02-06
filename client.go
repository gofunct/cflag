package cflag

import (
	"context"
	"github.com/gofunct/cflag/api/driver"
	"google.golang.org/grpc"
)

type ClientSignUpFunc func(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.Driver_SignupClient, error)
type ClientLoginFunc func(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.Driver_LoginClient, error)
type ClientDebugFunc func(ctx context.Context, in *driver.DebugRequest, opts ...grpc.CallOption) (driver.Driver_DebugClient, error)
type ClientExecFunc func(ctx context.Context, in *driver.ExecRequest, opts ...grpc.CallOption) (driver.Driver_ExecuteClient, error)
type ClientStremFunc func(ctx context.Context, opts ...grpc.CallOption) (driver.GRPCBroker_StartStreamClient, error)
type ClientWrapperFunc func(s *Client)

type Client struct {
	StreamFunc ClientStremFunc
	SignUpFunc  ClientSignUpFunc
	LoginFunc   ClientLoginFunc
	DebugFunc   ClientDebugFunc
	ExecFunc    ClientExecFunc
	WrapperFunc ClientWrapperFunc
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) StartStream(ctx context.Context, opts ...grpc.CallOption) (driver.GRPCBroker_StartStreamClient, error) {
	return c.StreamFunc(ctx, opts...)
}

func (c *Client) Signup(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.Driver_SignupClient, error) {
	return c.SignUpFunc(ctx, in, opts...)
}

func (c *Client) Login(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.Driver_LoginClient, error) {
	return c.LoginFunc(ctx, in, opts...)
}

func (c *Client) Debug(ctx context.Context, in *driver.DebugRequest, opts ...grpc.CallOption) (driver.Driver_DebugClient, error) {
	return c.DebugFunc(ctx, in, opts...)
}

func (c *Client) Execute(ctx context.Context, in *driver.ExecRequest, opts ...grpc.CallOption) (driver.Driver_ExecuteClient, error) {
	return c.ExecFunc(ctx, in, opts...)
}
