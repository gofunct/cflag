package cflag

import (
	"github.com/spf13/pflag"
	"reflect"
)

type Prompt struct {
	Key   string
	Help  string
	Dir  string
	Env  map[string]string
	Value string
	Text  string
}

func (p *Prompt) String() string {
	return p.Text
}

func (p *Prompt) Set(string) error {
	panic("implement me")
}

func (p *Prompt) Type() string {
	return reflect.TypeOf(p.Value).String()
}

func (p *Prompt) FlagSet(set *pflag.FlagSet) {
	set.Var(p, p.Name(), p.Usage())
}

func (p *Prompt) Name() string {
	return p.Key
}

func (p *Prompt) Usage() string {
	return p.Help
}


