package driver

import (
	"github.com/gofunct/gocfg"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

type Cflag interface {
	pflag.Value
	Name() string
	Usage() string
	Runnable() bool
	FlagSet(f *pflag.FlagSet)
}
type SetFunc func(s string) error

type WrapSet func(s string) string

type FlagVar struct {
	key     string
	typstr  string
	def     string
	setFunc SetFunc
	val     string
	*gocfg.GoCfg
	help string
}

func (f *FlagVar) SetFunc(fn func(s string) error) {
	f.setFunc = fn
}

func NewFlagVar(name string, def string, configFile string, envprefix string, typ string, usage string, set SetFunc) *FlagVar {
	f := &FlagVar{key: name, def: def, typstr: typ, help: usage, setFunc: set}
	f.GoCfg = gocfg.New(configFile, envprefix)
	if def != "" {
		f.GoCfg.SetDefault(name, def)
		f.val = def
	}
	return f
}

func (f *FlagVar) String() string {
	return f.val
}

func (f *FlagVar) Runnable() bool {
	if f.key == "" || f.help == "" {
		return false
	}
	if f.setFunc == nil {
		return false
	}
	return true
}

func (f *FlagVar) Set(s string) error {
	if !f.Runnable() {
		return errors.New("flag is not runnable, required: key, help, and setfunc")
	}
	if s == "" {
		s = f.def
		f.val = s
	}
	return f.setFunc(s)
}

func (f *FlagVar) Type() string {
	if f.typstr == "" {
		return "n/a"
	}
	return f.typstr
}

func (f *FlagVar) Name() string {
	return f.key
}

func (f *FlagVar) Usage() string {
	if f.help == "" {
		return "n/a"
	}
	return f.help
}

func (f *FlagVar) FlagSet(set *pflag.FlagSet) {
	set.Var(f, f.Name(), f.Usage())
}
