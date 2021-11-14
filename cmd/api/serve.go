package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"the-one/internal/app/configuration"
	"the-one/internal/app/server"

	"github.com/google/subcommands"
)

type ServeCmd struct {
	configFile string
	port       uint
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
	set.UintVar(&s.port, "p", 8080, "Port to serve")
}

func (s *ServeCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfg, err := configuration.Load(s.configFile)
	if err != nil {
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}
	serv := server.NewDefaultServer(cfg.Saver, uint16(s.port))
	err = serv.Start()
	fmt.Println(err)
	return subcommands.ExitFailure
}
