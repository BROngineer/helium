package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/brongineer/helium/flag"
	"github.com/brongineer/helium/flagset"
	"github.com/brongineer/helium/parser"
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

type customParser[T any] struct {
	*parser.EmbeddedParser
}

func (p *customParser[T]) ParseCmd(input string) (any, error) {
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

func newCustomParser[T any]() *customParser[T] {
	return &customParser[T]{&parser.EmbeddedParser{}}
}

func parse(args []string) (params, error) {
	fs := flagset.New().
		BindFlag(flag.Typed[srvParams]("server-config", flag.Parser(newCustomParser[srvParams]()))).
		BindFlag(flag.Typed[logParams]("log-config", flag.Parser(newCustomParser[logParams]()))).
		Build()

	if err := fs.Parse(args); err != nil {
		return params{}, err
	}

	return params{
		serverOpts: flagset.GetTypedFlag[srvParams](fs, "server-config"),
		loggerOpts: flagset.GetTypedFlag[logParams](fs, "log-config"),
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
