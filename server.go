package cflag

import (
	"context"
	"github.com/gofunct/cflag/api/driver"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	ShutdownFunc func(context.Context, *driver.Empty) (*driver.Empty, error)
	StreamFunc func(driver.GRPCBroker_StartStreamServer) error
	SignUpFunc func(*driver.User, driver.Driver_SignupServer) error
	LoginFunc func(*driver.User, driver.Driver_LoginServer) error
	DebugFunc func(*driver.DebugRequest, driver.Driver_DebugServer) error
	ExecFunc func(*driver.ExecRequest, driver.Driver_ExecuteServer) error
	SendFunc func(*driver.Info) error
	SetHeaderFunc func(metadata.MD) error
	HeaderFunc func(metadata.MD) error
	TrailerFunc func(metadata.MD)
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Signup(r *driver.User, w driver.Driver_SignupServer) error {
	return s.SignUpFunc(r, w)
}

func (s *Server) Login(r *driver.User, w driver.Driver_LoginServer) error {
	return s.LoginFunc(r, w)
}

func (s *Server) Debug(r *driver.DebugRequest, w driver.Driver_DebugServer) error {
	return s.DebugFunc(r, w)
}

func (s *Server) Execute(r *driver.ExecRequest, w driver.Driver_ExecuteServer) error {
	return s.ExecFunc(r, w)
}
