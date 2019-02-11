package cflag

import (
	"fmt"
	"github.com/spf13/pflag"
	"reflect"
	"text/template"
)

// Template is a `flag.Value` for `text.Template` arguments.
// The value of the `Root` field is used as a root template when specified.
type Template struct {
	Key  string
	Help string
	Root *template.Template

	Value *template.Template
	Text  string
}

func (fv *Template) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

func (fv *Template) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

func (fv *Template) Name() string {
	return fv.Key
}

func (fv *Template) Usage() string {
	return fv.Help
}

// Set is flag.Value.Set
func (fv *Template) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = template.New("")
	}
	t, err := root.New(fmt.Sprintf("%T(%p)", fv, fv)).Parse(v)
	if err == nil {
		fv.Value = t
	}
	return err
}

func (fv *Template) String() string {
	return fv.Text
}

// Templates is a `flag.Value` for `text.Template` arguments.
// The value of the `Root` field is used as a root template when specified.
type Templates struct {
	Key  string
	Help string
	Root *template.Template

	Values []*template.Template
	Texts  []string
}

func (fv *Templates) Type() string {
	return reflect.TypeOf(fv.Values).String()
}

func (fv *Templates) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

func (fv *Templates) Name() string {
	return fv.Key
}

func (fv *Templates) Usage() string {
	return fv.Help
}

// Set is flag.Value.Set
func (fv *Templates) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = template.New("")
	}
	t, err := root.New(fmt.Sprintf("%T(%p)", fv, fv)).Parse(v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, t)
	}
	return err
}

func (fv *Templates) String() string {
	return fmt.Sprint(fv.Texts)
}
