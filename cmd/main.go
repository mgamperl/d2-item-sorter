package main

import (
	"d2-item-sorter/pkg/api"
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

var finishRunner chan bool

type argT struct {
	Help bool `cli:"h,help" usage:"show help"`
}

// AutoHelp implements cli.AutoHelper interface
// NOTE: cli.Helper is a predefined type which implements cli.AutoHelper
func (argv *argT) AutoHelp() bool {
	return argv.Help
}

func main() {
	if err := cli.Root(cmdServe,
		cli.Tree(cmdHelp),
		cli.Tree(cmdWatch),
		cli.Tree(cmdTransfer),
		cli.Tree(cmdItemInfo),
		cli.Tree(cmdSort),
		cli.Tree(cmdGenerate),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var cmdHelp = cli.HelpCommand("display help information")

// child command
type watchT struct {
	cli.Helper
	Groups string `cli:"groups" usage:"configuration file for the sorting" dft:"./config/groups.yaml"`
}

type sortT struct {
	cli.Helper
	Name   string `cli:"*name" usage:"character to transfer items from"`
	Groups string `cli:"groups" usage:"configuration file for the sorting" dft:"./config/groups.yaml"`
}

type transferT struct {
	cli.Helper
	From   string `cli:"*from" usage:"character to transfer items from"`
	To     string `cli:"*to" usage:"character to transfer items to"`
	Groups string `cli:"groups" usage:"configuration file for the sorting" dft:"./config/groups.yaml"`
}

type itemInfoT struct {
	cli.Helper
}

type serveInfoT struct {
	cli.Helper
	Port int `cli:"port" usage:"specify port number for server to listen on" dft:"3000"`
}

type generateInfoT struct {
	cli.Helper
	D2dataFolder string `cli:"*d2dataFolder" usage:"specify d2data/json folder to generate data from json files" dft:"c:/dev/go/d2data/json"`
}

var cmdGenerate = &cli.Command{
	Name: "generate",
	Desc: "Generate some data for this program",
	Argv: func() interface{} { return new(generateInfoT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*generateInfoT)
		api.GenerateData(argv.D2dataFolder)
		return nil
	},
}

var cmdWatch = &cli.Command{
	Name: "watch",
	Desc: "Watches the shared stash to record runs and automatically sort items",
	Argv: func() interface{} { return new(watchT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*watchT)
		//ctx.String("Hello, child command, I am %s\n", argv.Name)

		groups := api.GetGroups(argv.Groups)
		api.StartWatcher(groups)
		return nil
	},
}

var cmdSort = &cli.Command{
	Name: "sort",
	Desc: "Sorts Items of a character or the shared stash",
	Argv: func() interface{} { return new(sortT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*sortT)
		//ctx.String("Hello, child command, I am %s\n", argv.Name)
		groups := api.GetGroups(argv.Groups)
		api.SortItems(argv.Name, groups)
		return nil
	},
}

var cmdTransfer = &cli.Command{
	Name: "transfer",
	Desc: "Transfers items from one character to another or to the shared stash (use _shared instead of a character name)",
	Argv: func() interface{} { return new(transferT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*transferT)
		groups := api.GetGroups(argv.Groups)
		api.TransferItems(argv.From, argv.To, groups)
		return nil
	},
}

var cmdItemInfo = &cli.Command{
	Name: "info",
	Desc: "Lists Information about the items you have",
	Argv: func() interface{} { return new(itemInfoT) },
	Fn: func(ctx *cli.Context) error {
		api.PrintItemInfo()
		return nil
	},
}

var cmdServe = &cli.Command{
	Name: "serve",
	Desc: "Start Server to access UI on http://localhost:port",
	Argv: func() interface{} { return new(serveInfoT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*serveInfoT)
		api.StartServer(argv.Port)
		return nil
	},
}
