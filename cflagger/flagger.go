package cflagger

import (
	"github.com/gofunct/cflag/driver"
	"github.com/spf13/viper"
)

type Flagger struct {
	Name        string
	RequireRoot bool
	Annotations map[string]string
	ConfigFiles []string
	EnvPrefix   string
	Bind        func(fn func(flag driver.Flag))
}

func NewFlagger(opts ...FlagOpt) *Flagger {
	f := &Flagger{}
	for _, o := range opts {
		o(f)
	}
	return f
}

func (f *Flagger) AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func (f *Flagger) Debug() {
	viper.Debug()
}

func (f *Flagger) AllKeys() []string {
	return viper.AllKeys()
}

func (f *Flagger) VisitAll(fn func(flag driver.Flag)) {
	f.Bind(fn)
}
