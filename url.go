package cflag

import (
	"fmt"
	"github.com/gofunct/goreflect"
	"net/url"
	"reflect"
	"strings"
)

// URL is a `flag.Value` for `url.URL` arguments.
type URL struct {
	Var     string
	Default *url.URL
	Value   *url.URL
	Text    string
}

func (fv *URL) HasChanged() bool {
	if reflect.DeepEqual(fv.Default, fv.Value) {
		return true
	}
	return false
}

func (fv *URL) Name() string {
	return fv.Var
}

func (fv *URL) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *URL) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *URL) Help() string {
	return "a URL"
}

// Set is flag.Value.Set
func (fv *URL) Set(v string) error {
	u, err := url.Parse(v)
	if err == nil {
		fv.Text = v
		fv.Value = u
	}
	return err
}

func (fv *URL) String() string {
	return fv.Text
}

// URLs is a `flag.Value` for `url.URL` arguments.
type URLs struct {
	Var     string
	Default []*url.URL
	Value   []*url.URL
	Texts   []string
}

func (fv *URLs) HasChanged() bool {
	if reflect.DeepEqual(fv.Default, fv.Value) {
		return true
	}
	return false
}

func (fv *URLs) Name() string {
	return fv.Var
}

func (fv *URLs) ValueString() string {
	return fmt.Sprintf("%s", fv.Value)
}

func (fv *URLs) ValueType() string {
	return goreflect.ValueTypeOf(fv.Value)
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *URLs) Help() string {
	return "a URL"
}

// Set is flag.Value.Set
func (fv *URLs) Set(v string) error {
	u, err := url.Parse(v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Value = append(fv.Value, u)
	}
	return err
}

func (fv *URLs) String() string {
	return strings.Join(fv.Texts, ",")
}
