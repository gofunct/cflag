package flagclient

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gofunct/cflag/driver"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
	"os"
)

type Client struct {
	
}

func (d *Client) Prompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return text
}


func (c *Client) Signup(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.GoCloud_SignupClient, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	in.Email = c.P
}

func (*Client) Login(ctx context.Context, in *driver.User, opts ...grpc.CallOption) (driver.GoCloud_LoginClient, error) {
	panic("implement me")
}

func (*Client) Debug(ctx context.Context, in *driver.DebugRequest, opts ...grpc.CallOption) (driver.GoCloud_DebugClient, error) {
	panic("implement me")
}

func (*Client) Execute(ctx context.Context, in *driver.ExecRequest, opts ...grpc.CallOption) (driver.GoCloud_ExecuteClient, error) {
	panic("implement me")
}

func (*Client) Recv() (*driver.DebugResponse, error) {
	panic("implement me")
}

func (*Client) Header() (metadata.MD, error) {
	panic("implement me")
}

func (*Client) Trailer() metadata.MD {
	panic("implement me")
}

func (*Client) CloseSend() error {
	panic("implement me")
}

func (*Client) Context() context.Context {
	panic("implement me")
}

func (*Client) SendMsg(m interface{}) error {
	panic("implement me")
}

func (*Client) RecvMsg(m interface{}) error {
	panic("implement me")
}
