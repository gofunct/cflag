package driver

import (
	"flag"
)

type Flag interface {
	flag.Value
	Help() string
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}
