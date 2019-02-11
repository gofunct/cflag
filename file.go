package cflag

import (
	"github.com/spf13/pflag"
	"os"
	"reflect"
	"strings"
)

// FileFlag is a `flag.Value` for file path arguments.
// By default, any errors from os.Stat are returned.
// Alternatively, the value of the `Validate` field is used as a validator when specified.
type FileFlag struct {
	Key      string
	Help     string
	Validate func(os.FileInfo, error) error

	Value string
}

func (fv FileFlag) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

func (fv FileFlag) Name() string {
	return fv.Key
}

func (fv FileFlag) Usage() string {
	return fv.Help
}

func (fv FileFlag) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

// Set is flag.Value.Set
func (fv FileFlag) Set(v string) error {
	info, err := os.Stat(v)
	fv.Value = v
	if fv.Validate != nil {
		return fv.Validate(info, err)
	}
	return err
}

func (fv FileFlag) String() string {
	return fv.Value
}

// Files is a `flag.Value` for file path arguments.
// By default, any errors from os.Stat are returned.
// Alternatively, the value of the `Validate` field is used as a validator when specified.
type Files struct {
	Key      string
	Help     string
	Validate func(os.FileInfo, error) error

	Values []string
}

func (fv *Files) Type() string {
	return reflect.TypeOf(fv.Values).String()
}

func (fv *Files) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

func (fv *Files) Name() string {
	return fv.Key
}

func (fv *Files) Usage() string {
	return fv.Help
}

// Set is flag.Value.Set
func (fv *Files) Set(v string) error {
	info, err := os.Stat(v)
	fv.Values = append(fv.Values, v)
	if fv.Validate != nil {
		return fv.Validate(info, err)
	}
	return err
}

func (fv *Files) String() string {
	return strings.Join(fv.Values, ",")
}
