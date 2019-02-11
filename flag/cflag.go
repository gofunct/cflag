package flag

import (
	"github.com/spf13/pflag"
)

type Cflag interface {
	pflag.Value
	FlagSet(set *pflag.FlagSet)
	Name() string
	Usage() string
}

type SetFunc func(s string) error


func (this *Flag) Type() string {
	return this.Typestr
}


func (this *Flag) Name() string {
	return this.GetKey()
}

func (this *Flag) Usage() string {
	return this.GetHelp()
}