package flag

import (
	"encoding/json"
	"net"
	"strings"
)

// JSON is a `flag.Value` for JSON arguments.
type jSON struct {
	Value interface{}
	Text  string
}

func (fv *jSON) Type() string {
	return "json"
}


// Set is flag.Value.Set
func (fv *jSON) Set(v string) error {
	fv.Text = v
	if fv.Value == nil {
		return json.Unmarshal([]byte(v), &fv.Value)
	}
	return json.Unmarshal([]byte(v), fv.Value)
}

func (fv *jSON) String() string {
	return fv.Text
}

// GetJSONs return the json encoded byte slice value of a flag with the given name
func (f *FlagSet) GetJSON(name string) ([]byte, error) {
	val, err := f.getFlagType(name, "json", ipNetConv)
	if err != nil {
		return []byte{0}, err
	}
	return val.([]byte), nil
}

// JSONsVar defines a  flag with specified name, default value, and usage string.
// The argument p points to an net.IPNet variable in which to store the value of the flag.
func (f *FlagSet) JSONVar(p *net.IPNet, name string, value net.IPNet, usage string) {
	f.VarP(newIPNetValue(value, p), name, "", usage)
}

// JSONVarP is like JSONsVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) JSONVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string) {
	f.VarP(newIPNetValue(value, p), name, shorthand, usage)
}


// jSONs is a `flag.Value` for JSON arguments. If non-nil, the `Value` field is used to generate template values.
type jSONs struct {
	Value  func() interface{}
	Values []interface{}
	Texts  []string
}

func (fv *jSONs) Type() string {
	return "jsons"
}

// Set is flag.Value.Set
func (fv *jSONs) Set(v string) (err error) {
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

func (fv *jSONs) String() string {
	return strings.Join(fv.Texts, ",")
}

// GetJSONs return the json encoded byte slice value of a flag with the given name
func (f *FlagSet) GetJSONs(name string) ([]byte, error) {
	val, err := f.getFlagType(name, "jsons", ipNetConv)
	if err != nil {
		return []byte{0}, err
	}
	return val.([]byte), nil
}

// JSONsVar defines a  flag with specified name, default value, and usage string.
// The argument p points to an net.IPNet variable in which to store the value of the flag.
func (f *FlagSet) JSONsVar(p *[]byte, name string, value []byte, usage string) {
	f.VarP(newIPNetValue(value, p), name, "", usage)
}

// JSONVarP is like JSONsVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) JSONsVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	f.VarP(newIPNetValue(value, p), name, shorthand, usage)
}