package cflag

import (
	"fmt"
	"github.com/spf13/pflag"
	"reflect"
	"sort"
	"strings"
)

// Enum is a `flag.Value` for one-of-a-fixed-set string arguments.
// The value of the `Choices` field defines the valid choices.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type Enum struct {
	Key           string
	Help          string
	Choices       []string
	CaseSensitive bool

	Value string
	Text  string
}

func (fv *Enum) Name() string {
	return fv.Key
}

func (fv *Enum) Usage() string {
	return fv.Help
}

func (fv *Enum) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

func (fv *Enum) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

// Set is flag.Value.Set
func (fv *Enum) Set(v string) error {
	fv.Text = v
	equal := strings.EqualFold
	if fv.CaseSensitive {
		equal = func(a, b string) bool { return a == b }
	}
	for _, c := range fv.Choices {
		if equal(c, v) {
			fv.Value = c
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
}

func (fv *Enum) String() string {
	return fv.Value
}

// EnumSet is a `flag.Value` for one-of-a-fixed-set string arguments.
// Only distinct values are returned.
// The value of the `Choices` field defines the valid choices.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type EnumSet struct {
	Key           string
	Help          string
	Choices       []string
	CaseSensitive bool

	Value map[string]bool
	Texts []string
}

func (fv *EnumSet) Name() string {
	return fv.Key
}

func (fv *EnumSet) Usage() string {
	return fv.Help
}

func (fv *EnumSet) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

func (fv *EnumSet) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

// Values returns a string slice of specified values.
func (fv *EnumSet) Values() (out []string) {
	for v := range fv.Value {
		out = append(out, v)
	}
	sort.Strings(out)
	return
}

// Set is flag.Value.Set
func (fv *EnumSet) Set(v string) error {
	equal := strings.EqualFold
	if fv.CaseSensitive {
		equal = func(a, b string) bool { return a == b }
	}
	var ok bool
	for _, c := range fv.Choices {
		if equal(c, v) {
			v = c
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
	}
	if fv.Value == nil {
		fv.Value = make(map[string]bool)
	}
	fv.Value[v] = true
	fv.Texts = append(fv.Texts, v)
	return nil
}

func (fv *EnumSet) String() string {
	return strings.Join(fv.Values(), ",")
}

// EnumSetCSV is a `flag.Value` for comma-separated enum arguments.
// Only distinct values are returned.
// The value of the `Choices` field defines the valid choices.
// If `Accumulate` is set, the values of all instances of the flag are accumulated.
// The `Separator` field is used instead of the comma when set.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type EnumSetCSV struct {
	Key           string
	Help          string
	Choices       []string
	Separator     string
	Accumulate    bool
	CaseSensitive bool

	Value map[string]bool
	Texts []string
}

func (fv *EnumSetCSV) Name() string {
	return fv.Key
}

func (fv *EnumSetCSV) Usage() string {
	return fv.Help
}

func (fv *EnumSetCSV) Type() string {
	return reflect.TypeOf(fv.Value).String()
}

func (fv *EnumSetCSV) FlagSet(set *pflag.FlagSet) {
	set.Var(fv, fv.Name(), fv.Usage())
}

// Values returns a string slice of specified values.
func (fv *EnumSetCSV) Values() (out []string) {
	for v := range fv.Value {
		out = append(out, v)
	}
	sort.Strings(out)
	return
}

// Set is flag.Value.Set
func (fv *EnumSetCSV) Set(v string) error {
	equal := strings.EqualFold
	if fv.CaseSensitive {
		equal = func(a, b string) bool { return a == b }
	}
	separator := fv.Separator
	if separator == "" {
		separator = ","
	}
	if !fv.Accumulate || fv.Value == nil {
		fv.Value = make(map[string]bool)
	}
	parts := strings.Split(v, separator)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		var ok bool
		var value string
		for _, c := range fv.Choices {
			if equal(c, part) {
				value = c
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
		}
		fv.Value[value] = true
		fv.Texts = append(fv.Texts, part)
	}
	return nil
}

func (fv *EnumSetCSV) String() string {
	return strings.Join(fv.Values(), ",")
}
