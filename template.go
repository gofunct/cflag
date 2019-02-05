package cflag

import (
	"fmt"
	"github.com/gofunct/goreflect"
	"reflect"
)

import (
	htemplate "html/template"
	"text/template"
)

// Template is a `flag.Value` for `text.Template` arguments.
// The value of the `Root` field is used as a root template when specified.
type Template struct {
	Var     string
	Root    *template.Template
	Default *template.Template
	Value   *template.Template
	Text    string
}

func (fv *Template) HasChanged() bool {
	if reflect.DeepEqual(fv.Default, fv.Value) {
		return true
	}
	return false
}

func (fv *Template) Name() string {
	return fv.Var
}

func (fv *Template) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *Template) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Template) Help() string {
	return "input a go text template to compile"
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
	Root *template.Template

	Values []*template.Template
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Templates) Help() string {
	return "a go template"
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

// HTMLTemplates is a `flag.Value` for `text.Template` arguments.
// The value of the `Root` field is used as a root template when specified.
type HTMLTemplates struct {
	Var     string
	Root    *htemplate.Template
	Default *htemplate.Template
	Value   []*htemplate.Template
	Texts   []string
}

func (fv *HTMLTemplates) HasChanged() bool {
	if reflect.DeepEqual(fv.Default, fv.Value) {
		return true
	}
	return false
}

func (fv *HTMLTemplates) Name() string {
	return fv.Var
}

func (fv *HTMLTemplates) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *HTMLTemplates) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *HTMLTemplates) Help() string {
	return "input a go html template to compile"
}

// Set is flag.Value.Set
func (fv *HTMLTemplates) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = htemplate.New("")
	}
	t, err := root.New(fmt.Sprintf("%T(%p)", fv, fv)).Parse(v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Value = append(fv.Value, t)
	}
	return err
}

func (fv *HTMLTemplates) String() string {
	return fmt.Sprint(fv.Texts)
}
