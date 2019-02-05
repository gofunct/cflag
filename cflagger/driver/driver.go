package driver

import "github.com/gofunct/cflag/driver"

type Flagger interface {
	VisitAll(fn func(flag driver.Flag))
	AllSettings() map[string]interface{}
	Debug()
	AllKeys() []string
}
