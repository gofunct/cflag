package cflag

import (
	"bufio"
	"fmt"
	"github.com/gofunct/cflag/driver"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"os"
	"time"
)

type FlagSet struct {
	Use            string
	Info           string
	ConfigPaths    []string
	ConfigType     string
	Env            map[string]string
	Inputs         []driver.Input
	SafeWrite      bool
	FlagHandler    driver.HandlerFunc
	FlagSetHandler func(d driver.FlagSet) error
	v              viper.Viper
}

func (f *FlagSet) Init() {
	f.v.SetEnvPrefix(f.Use)
	for _, p := range f.ConfigPaths {
		f.v.AddConfigPath(p)
	}
	f.v.AutomaticEnv()
	for k, v := range f.Env {
		ifErr(os.Setenv(k, v))
		ifErr(f.v.BindEnv(k))
	}
	if f.ConfigType != "" {
		f.v.SetConfigType(f.ConfigType)
	}
	err := f.v.ReadInConfig()
	if err != nil && f.SafeWrite {
		fmt.Println(f.v.SafeWriteConfig())
	}
}

func (*FlagSet) Name() string {
	return os.Args[0]
}

func (f *FlagSet) Usage() string {
	return f.Info
}

func (f *FlagSet) Debug() {
	f.v.Debug()
}

func (f *FlagSet) AllSettings() map[string]interface{} {
	return f.v.AllSettings()
}

func (f *FlagSet) ReadInConfig(reader io.Reader) {
	ifErr(f.v.ReadConfig(reader))
}

func (f *FlagSet) Get(key string) interface{} {
	return f.v.Get(key)
}

func (f *FlagSet) GetBool(key string) bool {
	return f.v.GetBool(key)
}

func (f *FlagSet) GetDuration(key string) time.Duration {
	return f.v.GetDuration(key)
}

func (f *FlagSet) GetFloat64(key string) float64 {
	return f.v.GetFloat64(key)
}

func (f *FlagSet) GetInt(key string) int {
	return f.v.GetInt(key)
}

func (f *FlagSet) GetInt32(key string) int32 {
	return f.v.GetInt32(key)
}

func (f *FlagSet) GetInt64(key string) int64 {
	return f.v.GetInt64(key)
}

func (f *FlagSet) GetSizeInBytes(key string) uint {
	return f.v.GetSizeInBytes(key)
}

func (f *FlagSet) GetString(key string) string {
	return f.v.GetString(key)
}

func (f *FlagSet) GetStringMap(key string) map[string]interface{} {
	return f.v.GetStringMap(key)
}

func (f *FlagSet) GetStringMapString(key string) map[string]string {
	return f.v.GetStringMapString(key)
}

func (f *FlagSet) GetStringMapStringSlice(key string) map[string][]string {
	return f.v.GetStringMapStringSlice(key)
}

func (f *FlagSet) GetStringSlice(key string) []string {
	return f.v.GetStringSlice(key)
}

func (f *FlagSet) GetTime(key string) time.Time {
	return f.v.GetTime(key)
}

func (f *FlagSet) InConfig(key string) bool {
	return f.v.InConfig(key)
}

func (f *FlagSet) SafeWriteFileAs(s string) error {
	return f.v.SafeWriteConfigAs(s)
}

func (f *FlagSet) VisitAll(flags ...*pflag.Flag) {
	for _, fl := range flags {
		f.FlagHandler(fl)
	}
}

func (f *FlagSet) Execute(set driver.FlagSet) error {
	return f.FlagSetHandler(f)
}

func (f *FlagSet) Input() []driver.Input {
	return f.Inputs
}

func NewFlagSet(name string, configPaths []string, configType string, env map[string]string, inputs []driver.Input, safeWrite bool, flagHandler driver.HandlerFunc, flagSetHandler func(f *FlagSet) error) *FlagSet {
	return &FlagSet{Use: name, ConfigPaths: configPaths, ConfigType: configType, Env: env, Inputs: inputs, SafeWrite: safeWrite, FlagHandler: flagHandler, FlagSetHandler: flagSetHandler}
}

func NewHandler(i ...driver.Input) driver.HandlerFunc {
	return func(fl *pflag.Flag) {
		for _, input := range i {
			switch input {
			case driver.ARGS:

			case driver.FILE:
			case driver.ENV:
			case driver.Query:

			}
		}
	}
}

func Prompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return text
}

func ifErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}
