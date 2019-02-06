package cflag

import (
	"context"
	"github.com/gofunct/cflag/api/driver"
	"google.golang.org/grpc/metadata"
)

type ServerShutdownFunc func(context.Context, *driver.Empty) (*driver.Empty, error)
type ServerStreamFunc func(driver.GRPCBroker_StartStreamServer) error
type ServerSignUpFunc func(*driver.User, driver.Driver_SignupServer) error
type ServerLoginFunc func(*driver.User, driver.Driver_LoginServer) error
type ServerDebugFunc func(*driver.DebugRequest, driver.Driver_DebugServer) error
type ServerExecFunc func(*driver.ExecRequest, driver.Driver_ExecuteServer) error

type ServerWrapperFunc func(s *Server)

type Server struct {
	ShutdownFunc      ServerShutdownFunc
	StreamFunc        ServerStreamFunc
	SignUpFunc        SignUpFunc
	LoginFunc         ServerLoginFunc
	DebugFunc         ServerDebugFunc
	ExecFunc          ExecFunc
	ServerWrapperFunc ServerWrapperFunc
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
