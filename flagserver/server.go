package flagserver

import (
	"context"
	"github.com/gofunct/cflag/api/driver"
	"google.golang.org/grpc/metadata"
)

func NewCflag() *Server {
	return &Server{}
}

type Server struct {
}

func (*Server) Shutdown(context.Context, *driver.Empty) (*driver.Empty, error) {
	panic("implement me")
}

func (*Server) StartStream(driver.GRPCBroker_StartStreamServer) error {
	panic("implement me")
}

func (*Server) Signup(*driver.User, driver.GoCloud_SignupServer) error {
	panic("implement me")
}

func (*Server) Login(*driver.User, driver.GoCloud_LoginServer) error {
	panic("implement me")
}

func (*Server) Debug(*driver.DebugRequest, driver.GoCloud_DebugServer) error {
	panic("implement me")
}

func (*Server) Execute(*driver.ExecRequest, driver.GoCloud_ExecuteServer) error {
	panic("implement me")
}

func (*Server) Send(*driver.Info) error {
	panic("implement me")
}

func (*Server) SetHeader(metadata.MD) error {
	panic("implement me")
}

func (*Server) SendHeader(metadata.MD) error {
	panic("implement me")
}

func (*Server) SetTrailer(metadata.MD) {
	panic("implement me")
}
