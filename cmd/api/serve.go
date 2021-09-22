package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"the-one/cmd/api/configuration"
	"the-one/cmd/api/server"

	"github.com/google/subcommands"
)

type ServeCmd struct {
	configFile string
}

func (s *ServeCmd) Name() string {
	return "serve"
}

func (s *ServeCmd) Synopsis() string {
	return "Serve the application"
}

func (s *ServeCmd) Usage() string {
	return "serve"
}

func (s *ServeCmd) SetFlags(set *flag.FlagSet) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	bd := path.Dir(ex)
	set.StringVar(&s.configFile, "f", bd+"/uatu-config.yml", "File to Config file")
}

func (s *ServeCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfg, err := configuration.Load(s.configFile)
	if err != nil {
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}
	serv := server.NewDefaultServer(cfg.Saver, 8080)
	err = serv.Start()
	fmt.Println(err)
	return subcommands.ExitFailure
}
