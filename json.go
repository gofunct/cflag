package cflag

import (
	"encoding/json"
	"github.com/spf13/pflag"
	"reflect"
	"strings"
)

// JSON is a `flag.Value` for JSON arguments.
type JSON struct {
	Key   string
	Help  string
	Value interface{}
	Text  string
}

func (fv *JSON) Name() string {
	return fv.Key
}

func (fv *JSON) Usage() string {
	return fv.Help
}

func (fv *JSON) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

func (fv *JSON) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

// Set is flag.Value.Set
func (fv *JSON) Set(v string) error {
	fv.Text = v
	if fv.Value == nil {
		return json.Unmarshal([]byte(v), &fv.Value)
	}
	return json.Unmarshal([]byte(v), fv.Value)
}

func (fv *JSON) String() string {
	return fv.Text
}

// JSONs is a `flag.Value` for JSON arguments. If non-nil, the `Value` field is used to generate template values.
type JSONs struct {
	Key    string
	Help   string
	Value  func() interface{}
	Values []interface{}
	Texts  []string
}

func (fv *JSONs) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

func (fv *JSONs) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

func (fv *JSONs) Name() string {
	return fv.Key
}

func (fv *JSONs) Usage() string {
	return fv.Help
}

// Set is flag.Value.Set
func (fv *JSONs) Set(v string) (err error) {
	var value interface{}
	if fv.Value != nil {
		value = fv.Value()
		err = json.Unmarshal([]byte(v), value)
	} else {
		err = json.Unmarshal([]byte(v), &value)
	}
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, value)
	}
	return err
}

func (fv *JSONs) String() string {
	return strings.Join(fv.Texts, ",")
}
