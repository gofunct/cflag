package cflag

import (
	"context"
	"github.com/gofunct/cflag/flag"
	"github.com/gofunct/cflag/util"
	"github.com/hashicorp/go-getter"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"log"
	"os"
	"os/signal"
	"sync"
)

type Download struct {
	*flag.Flag

	Dest  string
}

func (d *Download) Set(s string) error {
	return InitLoad(d)(s)
}

func (d *Download) FlagSet(set *pflag.FlagSet) {
	set.Var(d, d.GetKey(), d.GetHelp())
}

func NewDownloadVar(dfault string, dest string) *Download{
	d := &Download{
			Flag:    &flag.Flag{
				Key:        "load",
				Help:       "download a file or repository from git, s3, or http",
				Required:   false,
				Default:    dfault,
				Typestr:    "string",
				Text:       "",
			},
		Dest:  dest,
	}
	return d
}

func InitLoad(d *Download) func(s string) error {
	return func(s string) error {
		pwd, err := os.Getwd()
		if err != nil {
			return errors.Wrapf(err, "Error getting wd: %s", pwd)
		}

		opts := []getter.ClientOption{}
		opts = append(opts, getter.WithProgress(util.DefaultProgressBar))

		ctx, cancel := context.WithCancel(context.Background())
		// Build the client
		client := &getter.Client{
			Ctx:     ctx,
			Src:     s,
			Dst:     d.Dest,
			Pwd:     pwd,
			Mode:    getter.ClientModeAny,
			Options: opts,
		}

		wg := sync.WaitGroup{}
		wg.Add(1)
		errChan := make(chan error, 2)
		go func() {
			defer wg.Done()
			defer cancel()
			if err := client.Get(); err != nil {
				errChan <- err
			}
		}()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)

		select {
		case sig := <-c:
			signal.Reset(os.Interrupt)
			cancel()
			wg.Wait()
			log.Printf("signal %v", sig)
		case <-ctx.Done():
			wg.Wait()
			log.Printf("success!")
		case err := <-errChan:
			wg.Wait()
			return errors.Wrapf(err, "Error downloading: %s", s)
		}
		return nil
	}
}