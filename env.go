package cflag

import (
	"fmt"
	"github.com/gofunct/goreflect"
	"reflect"
	"strings"
)

type EnvMap struct {
	Key       string
	Separator string
	Defaults  map[string]string
	Value     map[string]string
	Texts     []string
}

func (fv *EnvMap) Name() string {
	return fv.Key
}

func (fv *EnvMap) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *EnvMap) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

func (fv *EnvMap) HasChanged() bool {
	return reflect.DeepEqual(fv.Defaults, fv.Value)
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *EnvMap) Help() string {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	return fmt.Sprintf("a key/value pair KEY%sVALUE", separator)
}

// Set is flag.Value.Set
func (fv *EnvMap) Set(v string) error {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	i := strings.Index(v, separator)
	if i < 0 {
		return fmt.Errorf(`"%s" must have the form KEY%sVALUE`, v, separator)
	}
	fv.Texts = append(fv.Texts, v)
	if fv.Value == nil {
		fv.Value = make(map[string]string)
	}
	fv.Value[v[:i]] = v[i+len(separator):]
	return nil
}

func (fv *EnvMap) String() string {
	return strings.Join(fv.Texts, ", ")
}
