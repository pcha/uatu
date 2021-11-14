package main

//
//import (
//	"context"
//	"flag"
//	"fmt"
//
//	"github.com/google/subcommands"
//	"github.com/c-bata/go-prompt"
//)
//
//type InitCmd struct {
//	configFile string
//	saver string
//}
//
//func (i *InitCmd) Name() string {
//	return "init"
//}
//
//func (i *InitCmd) Synopsis() string {
//	return "Create configuration file"
//}
//
//func (i *InitCmd) Usage() string {
//	return "init"
//}
//
//func (i *InitCmd) SetFlags(set *flag.FlagSet) {
//
//}
//
//func (i *InitCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
//	fmt.Printf("Enter configutation file (Default: %v)\n", i.configFile)
//	cfgFile := prompt.Input("> ", nil)
//	if cfgFile != "" {
//		i.configFile = cfgFile
//	}
//	fmt.Printf("Choose the Saver (Default: %v)\n", i.)
//}
////
