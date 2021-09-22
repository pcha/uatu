package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func init() {

}

type Command interface {
	Run([]string) error
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&ServeCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
	//fs := flag.Args()
	//fmt.Println(fs)
	//cmd := flag.Arg(0)
	//if cmd == "" {
	//	log.Fatal("Must indicate command")
	//}
	//switch cmd {
	//case "serve":
	//
	//	cfg, err := configuration.Load(file)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	s := server.NewDefaultServer(cfg.Saver, port)
	//	err = s.Start()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}

//func getServeCommand() *subcmd.SubCommand {
//	fs := flag.NewFlagSet("")
//	return subcmd.NewSubCommand("serve", )
//}
