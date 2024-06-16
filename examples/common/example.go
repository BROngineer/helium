package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brongineer/helium"
	"github.com/brongineer/helium/flag"
)

type params struct {
	bindAddress     string
	bindPort        uint32
	logLevel        string
	developmentMode bool
	timeout         time.Duration
	peers           []string
}

func parse(args []string) (params, error) {
	fs := helium.NewFlagSet()
	fs.String("bind-address", flag.Description("bind listen address"), flag.DefaultValue("localhost"))
	fs.Uint32("bind-port", flag.Description("bind listen port"), flag.DefaultValue(uint32(80)))
	fs.String("log-level", flag.Description("logging level"), flag.DefaultValue("info"))
	fs.Bool("development-mode", flag.Shorthand("d"), flag.DefaultValue(false))
	fs.Duration("timeout", flag.Description("context timeout"), flag.Shorthand("t"), flag.DefaultValue(time.Minute))
	fs.StringSlice("peers", flag.Description("remote peers"), flag.DefaultValue([]string{}))

	if err := fs.Parse(args); err != nil {
		return params{}, err
	}

	return params{
		bindAddress:     fs.GetString("bind-address"),
		bindPort:        fs.GetUint32("bind-port"),
		logLevel:        fs.GetString("log-level"),
		developmentMode: fs.GetBool("development-mode"),
		timeout:         fs.GetDuration("timeout"),
		peers:           fs.GetStringSlice("peers"),
	}, nil
}

func main() {
	opts, err := parse(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	fmt.Println("Parsed params:", opts)
}
