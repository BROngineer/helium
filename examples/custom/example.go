package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/brongineer/helium"
	"github.com/brongineer/helium/flag"
)

type params struct {
	serverOpts srvParams
	loggerOpts logParams
}

type srvParams struct {
	BindAddress string        `json:"bindAddress"`
	BindPort    uint32        `json:"bindPort"`
	Timeout     time.Duration `json:"timeout"`
	Peers       []string      `json:"peers"`
}

type logParams struct {
	LogLevel  string `json:"logLevel"`
	LogFormat string `json:"logFormat"`
	DevMode   bool   `json:"devMode"`
}

func parser[T any](input string) (any, error) {
	var (
		b    []byte
		err  error
		opts T
	)
	if input == "" {
		return nil, fmt.Errorf("path to config file is empty")
	}
	b, err = os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &opts); err != nil {
		return nil, err
	}
	return &opts, nil
}

func parse(args []string) (params, error) {
	fs := helium.NewFlagSet()
	helium.CustomFlag[srvParams](fs, "server-config", flag.CommandLineParser(parser[srvParams]))
	helium.CustomFlag[logParams](fs, "log-config", flag.CommandLineParser(parser[logParams]))

	if err := fs.Parse(args); err != nil {
		return params{}, err
	}

	return params{
		serverOpts: helium.GetCustomFlag[srvParams](fs, "server-config"),
		loggerOpts: helium.GetCustomFlag[logParams](fs, "log-config"),
	}, nil
}

func main() {
	opts, err := parse(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	fmt.Println("Parsed params:")
	fmt.Println(opts.serverOpts)
	fmt.Println(opts.loggerOpts)
}
