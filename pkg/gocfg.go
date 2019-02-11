package pkg

import (
	"github.com/gofunct/gocfg"
	"os"
)

var GoCfg = gocfg.New(os.Getenv("PWD")+"/.config.yaml", "CFLAG")