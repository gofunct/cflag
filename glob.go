package cflag

import (
	"fmt"
	"github.com/gobwas/glob"
	"github.com/gofunct/goreflect"
	"github.com/pkg/errors"
	"path/filepath"
	"reflect"
)

// Glob is a `flag.Value` for glob syntax arguments.
// By default, `filepath.Separator` is used as a separator.
// If `Separators` is non-nil, its elements are used as separators.
// To have no separators, set `Separators` to a (non-nil) pointer to an empty slice.
type Glob struct {
	Var        string
	Separators *[]rune
	Default    glob.Glob
	Pattern    string
	Seperator  string
	Value      glob.Glob
	Text       string
}

func (fv *Glob) HasChanged() bool {
	r := []rune(fv.Seperator)
	g, err := glob.Compile(fv.Pattern, r...)
	if err != nil {
		panic(errors.WithStack(err))
	}

	if reflect.DeepEqual(g, fv.Value) {
		return true
	}
	return false
}

func (fv *Glob) Name() string {
	return fv.Var
}

func (fv *Glob) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *Glob) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Glob) Help() string {
	separators := []rune{filepath.Separator}
	if fv.Separators != nil {
		separators = *fv.Separators
	}
	if len(separators) == 0 {
		return "a glob expression"
	}
	if len(separators) == 1 {
		return fmt.Sprintf("a glob expression with separator %q", separators[0])
	}
	return fmt.Sprintf("a glob expression with separators %q", separators)
}

// Set is flag.Value.Set
func (fv *Glob) Set(v string) error {
	separators := fv.Separators
	if separators == nil {
		separators = &[]rune{filepath.Separator}
	}
	g, err := glob.Compile(v, *separators...)
	if err != nil {
		return err
	}
	fv.Text = v
	fv.Value = g
	return nil
}

func (fv *Glob) String() string {
	return fv.Text
}
