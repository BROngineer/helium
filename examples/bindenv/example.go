package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brongineer/helium/env"
	"github.com/brongineer/helium/flag"
	"github.com/brongineer/helium/flagset"
)

type params struct {
	bindAddress     string
	bindPort        uint32
	logLevel        string
	developmentMode bool
	timeout         time.Duration
	peers           []string
}

func setEnv() {
	_ = os.Setenv("BIND_ENV_EXAMPLE_BIND_ADDRESS", "1.2.3.4")
	_ = os.Setenv("BIND_ENV_EXAMPLE_BIND_PORT", "65535")
}

func unsetEnv() {
	_ = os.Unsetenv("BIND_ENV_EXAMPLE_BIND_ADDRESS")
	_ = os.Unsetenv("BIND_ENV_EXAMPLE_BIND_PORT")
}

func parse(args []string) (params, error) {
	fs := flagset.New(env.Prefix("BIND_ENV_EXAMPLE"), env.Capitalized(), env.VarNameReplace("-", "_")).
		BindFlag(flag.String("bind-address", flag.Description("bind listen address"), flag.DefaultValue("localhost"))).
		BindFlag(flag.Uint32("bind-port", flag.Description("bind listen port"), flag.DefaultValue(uint32(80)))).
		BindFlag(flag.String("log-level", flag.Description("logging level"), flag.DefaultValue("info"))).
		BindFlag(flag.Bool("development-mode", flag.Shorthand("d"), flag.DefaultValue(false))).
		BindFlag(flag.Duration("timeout", flag.Description("context timeout"), flag.Shorthand("t"), flag.DefaultValue(time.Minute))).
		BindFlag(flag.StringSlice("peers", flag.Description("remote peers"), flag.DefaultValue([]string{}))).
		Build()

	if err := fs.Parse(args); err != nil {
		return params{}, err
	}

	if err := fs.BindEnvVars(); err != nil {
		return params{}, err
	}

	return params{
		bindAddress:     flagset.GetString(fs, "bind-address"),
		bindPort:        flagset.GetUint32(fs, "bind-port"),
		logLevel:        flagset.GetString(fs, "log-level"),
		developmentMode: flagset.GetBool(fs, "development-mode"),
		timeout:         flagset.GetDuration(fs, "timeout"),
		peers:           flagset.GetStringSlice(fs, "peers"),
	}, nil
}

func main() {
	setEnv()
	defer unsetEnv()
	opts, err := parse(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	fmt.Println("Parsed params:", opts)
}
