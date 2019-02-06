package cflag

import (
	"github.com/gofunct/cflag/api/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"context"
)

type Client struct {
	SignUpFunc func(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.Driver_SignupClient, error)
	LoginFunc func(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.Driver_LoginClient, error)
	DebugFunc func(ctx context.Context, in *driver.DebugRequest, opts ...grpc.CallOption) (driver.Driver_DebugClient, error)
	DebugRcvFunc func() (*driver.DebugResponse, error)
	ExecFunc func(ctx context.Context, in *driver.ExecRequest, opts ...grpc.CallOption) (driver.Driver_ExecuteClient, error)
	HeaderFunc func() (metadata.MD, error)
	TrailerFunc func() metadata.MD
}

func NewClient() *Client {
	return &Client{}
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
