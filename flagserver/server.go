package flagserver

import (
	"context"
	"github.com/gofunct/cflag/driver"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	
}

func NewCflag() *Server {
	return &Server{}
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

func (*Server) Context() context.Context {
	panic("implement me")
}

func (*Server) SendMsg(m interface{}) error {
	panic("implement me")
}

func (*Server) RecvMsg(m interface{}) error {
	panic("implement me")
}

