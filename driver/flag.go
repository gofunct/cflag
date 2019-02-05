package driver

import (
	"github.com/spf13/pflag"
	"io"
	"os"
	"time"
)

type Input int

const (
	ARGS Input = iota
	FILE
	ENV
	Query
)

var CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

type HandlerFunc func(*pflag.Flag)

type FlagSet interface {
	Name() string
	Usage() string
	Debug()
	AllSettings() map[string]interface{}
	ReadInConfig(reader io.Reader)
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	SafeWriteFileAs(s string) error
	Execute(set FlagSet) error
	VisitAll(f ...*pflag.Flag)
	Input() []Input
}
