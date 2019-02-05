package example

import (
	"context"
	"github.com/gofunct/cflag"
	"github.com/gofunct/cflag/driver"
	"github.com/spf13/pflag"
	"log"
	"reflect"
)

func init() {
	Set.AddFlags(sport)
}

var (
	sport = &pflag.Flag{
		Name:                "sport",
		Shorthand:           "s",
		Usage:               "your favorite sport",
		DefValue:            "basketball",
	}
	Set = cflag.NewFlagSet(
	"example",
	"just an example",
	map[string]string{
		"FAVORITE_SPORT" : "FOOTBALL",
	},
	"example",
	[]string{"config"},
	true,

	func(d driver.FlagSet) error {

		return nil
	},
	)
)
func main() {
	log.Fatalln(driver.Execute(context.Background(), Set))
}

func NewFlagFunc(f driver.FlagSet) driver.HandlerFunc {
	return func(flag *pflag.Flag) {
		if flag.Value == nil {
			if f.InConfig(flag.Name) {
				reflect.TypeOf()
			}
		}
	}
}