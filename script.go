package cflag

import (
	"github.com/spf13/pflag"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

type Script struct {
	Key   string
	Help  string
	Dir  string
	Env  map[string]string
	Value string
	Text  string
}

func (s *Script) String() string {
	return s.Text
}

func (s *Script) Set(script string) error {
	s.Text = script
	cmd := exec.Command("bash", "-c", script)
	if cmd.Env != nil {
		for k, v := range  s.Env {
			_ =os.Setenv(strings.ToUpper(k), v)
		}
	}
	cmd.Env = os.Environ()
	if s.Dir != "" {
		cmd.Dir = s.Dir
	} else {
		cmd.Dir = os.Getenv("PWD")
	}
	return cmd.Run()
}

func (s *Script) Type() string {
	return reflect.TypeOf(s.Value).String()
}

func (s *Script) FlagSet(set *pflag.FlagSet) {
	set.Var(s, s.Name(), s.Usage())
}

func (s *Script) Name() string {
	return s.Key
}

func (s *Script) Usage() string {
	return s.Help
}

