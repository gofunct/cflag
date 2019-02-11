package flag

import (
	"github.com/spf13/pflag"
	"reflect"
)

type CFLAG interface {
	pflag.Value
	FlagSet(set *pflag.FlagSet)
	Name() string
	Usage() string
}

type Cflag struct {
	*Flag
	SetFunc func(s string) error
}
func (this *Cflag) Set(s string) error {
	return this.SetFunc(s)
}

func (this *Cflag) FlagSet(set *pflag.FlagSet) {
	set.Var(this, this.Name(), this.Usage())
}

func (this *Flag) Type() string {
	return reflect.TypeOf(this.Value).String()
}


func (this *Flag) Name() string {
	return this.GetKey()
}

func (this *Flag) Usage() string {
	return this.GetHellp()
}