package cflag

import "github.com/gofunct/cflag/driver"

var Docker = driver.NewFlagVar(
	"docker",
	"",
	"",
	"",
	"",
	"",
	func(s string) error {
		return nil
	})
