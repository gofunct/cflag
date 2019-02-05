package cflag

import (
	"fmt"
	"github.com/gofunct/goreflect"
	"os"
	"reflect"
	"strings"
)

// File is a `flag.Value` for file path arguments.
// By default, any errors from os.Stat are returned.
// Alternatively, the value of the `Validate` field is used as a validator when specified.
type File struct {
	Var      string
	Validate func(os.FileInfo, error) error
	Default  string
	Value    string
}

func (fv *File) Help() string {
	return "a file path"
}

func (fv *File) HasChanged() bool {
	if reflect.DeepEqual(fv.Default, fv.Value) {
		return true
	}
	return false
}

func (fv *File) Name() string {
	return fv.Var
}

func (fv *File) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *File) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Set is flag.Value.Set
func (fv *File) Set(v string) error {
	info, err := os.Stat(v)
	fv.Value = v
	if fv.Validate != nil {
		return fv.Validate(info, err)
	}
	return err
}

func (fv *File) String() string {
	return fv.Value
}

// Files is a `flag.Value` for file path arguments.
// By default, any errors from os.Stat are returned.
// Alternatively, the value of the `Validate` field is used as a validator when specified.
type Files struct {
	Var      string
	Validate func(os.FileInfo, error) error
	Default  []string
	Value    []string
}

func (fv *Files) Help() string {
	return "the path to a file"
}

func (fv *Files) HasChanged() bool {
	if goreflect.StringSliceMatches(fv.Default, fv.Value) {
		return true
	}
	return false
}

func (fv *Files) Name() string {
	return fv.Var
}

func (fv *Files) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *Files) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Set is flag.Value.Set
func (fv *Files) Set(v string) error {
	info, err := os.Stat(v)
	fv.Value = append(fv.Value, v)
	if fv.Validate != nil {
		return fv.Validate(info, err)
	}
	return err
}

func (fv *Files) String() string {
	return strings.Join(fv.Value, ",")
}
