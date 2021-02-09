package main

import (
	"d2-item-sorter/hotkeys"
	"d2-item-sorter/sorter"
	"d2-item-sorter/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
	"github.com/nokka/d2s"
)

var statistics = map[string]int{}

var ignoreFrom = 0
var ignoreTo = 3
var pickupPages = 1

var keys = map[int16]*hotkeys.Hotkey{
	1: {Id: 1, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'O'}, // ALT+CTRL+O
	2: {Id: 2, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'L'}, // ALT+SHIFT+M
	3: {Id: 3, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'X'}, // ALT+CTRL+X
}

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
	if err := cli.Root(serve,
		cli.Tree(help),
		cli.Tree(watch),
		cli.Tree(transfer),
		cli.Tree(itemInfo),
		cli.Tree(sort),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help information")

// child command
type watchT struct {
	cli.Helper
	Groups string `cli:"groups" usage:"configuration file for the sorting" dft:"./groupConfigs/groups.yaml"`
}

type sortT struct {
	cli.Helper
	Name   string `cli:"*name" usage:"character to transfer items from"`
	Groups string `cli:"groups" usage:"configuration file for the sorting" dft:"./groupConfigs/groups.yaml"`
}

type transferT struct {
	cli.Helper
	From   string `cli:"*from" usage:"character to transfer items from"`
	To     string `cli:"*to" usage:"character to transfer items to"`
	Groups string `cli:"groups" usage:"configuration file for the sorting" dft:"./groupConfigs/groups.yaml"`
}

type itemInfoT struct {
	cli.Helper
}

type serveInfoT struct {
	cli.Helper
	Port int `cli:"port" usage:"specify port number for server to listen on" dft:"3000"`
}

var watch = &cli.Command{
	Name: "watch",
	Desc: "Watches the shared stash to record runs and automatically sort items",
	Argv: func() interface{} { return new(watchT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*watchT)
		//ctx.String("Hello, child command, I am %s\n", argv.Name)

		groups := sorter.GetGroups(argv.Groups)
		startWatcher(groups)
		return nil
	},
}

var sort = &cli.Command{
	Name: "sort",
	Desc: "Sorts Items of a character or the shared stash",
	Argv: func() interface{} { return new(sortT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*sortT)
		//ctx.String("Hello, child command, I am %s\n", argv.Name)
		groups := sorter.GetGroups(argv.Groups)
		sortItems(argv.Name, groups)
		return nil
	},
}

var transfer = &cli.Command{
	Name: "transfer",
	Desc: "Transfers items from one character to another or to the shared stash (use _shared instead of a character name)",
	Argv: func() interface{} { return new(transferT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*transferT)
		groups := sorter.GetGroups(argv.Groups)
		transferItems(argv.From, argv.To, groups)
		return nil
	},
}

var itemInfo = &cli.Command{
	Name: "info",
	Desc: "Lists Information about the items you have",
	Argv: func() interface{} { return new(itemInfoT) },
	Fn: func(ctx *cli.Context) error {
		printItemInfo()
		return nil
	},
}

var serve = &cli.Command{
	Name: "serve",
	Desc: "Start Server to access UI on http://localhost:port",
	Argv: func() interface{} { return new(serveInfoT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*serveInfoT)
		startServer(argv.Port)
		return nil
	},
}

func getSaveDir() (string, error) {
	saveDir := utils.GetD2SavePath()
	fmt.Printf("D2 save directory: %s\n", saveDir)

	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		panic("directory " + saveDir + " does not exist")
		//(fmt.Printf("%s does not exist", saveDir)
		//return _, err
	}

	saveDir, err := filepath.Abs(saveDir)

	return saveDir, err
}

func sortItems(name string, groups []sorter.ItemGroup) {
	saveDir, err := getSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	characters, _ := sorter.ParseCharacters(saveDir)

	if _, okFrom := characters[name]; okFrom == false && name != "_shared" {
		fmt.Printf("%s is not a valid character name or does not exist\n", name)
		return
	}

	if name == "_shared" {
		fmt.Printf("sorting items for Shared Stash")

		sharedStash := sorter.ParseSharedStash(saveDir)
		sharedStash.Reload()
		sharedStash.SortItems(groups)
		sharedStash.Save("")

	} else {
		fmt.Printf("sorting items for %s\n", name)

		characters[name].Reload()
		characters[name].SortPersonalStash(groups)
		characters[name].Save("")
	}

	return
}

func transferItems(nameFrom string, nameTo string, groups []sorter.ItemGroup) {
	saveDir, err := getSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := sorter.ParseSharedStash(saveDir)
	characters, _ := sorter.ParseCharacters(saveDir)

	if _, okFrom := characters[nameFrom]; okFrom == false && nameFrom != "_shared" {
		fmt.Printf("%s is not a valid character name or does not exist\n", nameFrom)
		return
	}

	if _, okTo := characters[nameTo]; okTo == false && nameTo != "_shared" {
		fmt.Printf("%s is not a valid character name or does not exist\n", nameTo)
		return
	}

	if nameFrom == "_shared" {
		fmt.Printf("transfering items from Shared Stash to %s\n", nameTo)
		sharedStash.Reload()
		characters[nameTo].Reload()

		sharedStash.TransferPagesTo(&characters[nameTo].PersonalStash)

		sharedStash.SortItems(groups)
		characters[nameTo].SortPersonalStash(groups)

		sharedStash.Save("")
		characters[nameTo].Save("")

	} else if nameTo == "_shared" {
		fmt.Printf("transfering items from %s to Shared Stash\n", nameFrom)
		sharedStash.Reload()
		characters[nameFrom].Reload()

		characters[nameFrom].PersonalStash.TransferPagesTo(&sharedStash)

		sharedStash.SortItems(groups)
		characters[nameFrom].SortPersonalStash(groups)

		sharedStash.Save("")
		characters[nameFrom].Save("")

	} else {
		fmt.Printf("transfering items from %s to %s\n", nameFrom, nameTo)
		characters[nameFrom].Reload()
		characters[nameTo].Reload()

		characters[nameFrom].PersonalStash.TransferPagesTo(&characters[nameTo].PersonalStash)

		characters[nameTo].SortPersonalStash(groups)
		characters[nameFrom].SortPersonalStash(groups)

		characters[nameTo].Save("")
		characters[nameFrom].Save("")
	}

	return
}

func addStatistic(item sorter.SortableItem) {
	if item.IsUnique || item.IsSetItem {
		statistics[item.GetName()]++
	}
}

func printItemInfo() {
	saveDir, err := getSaveDir()

	statistics = map[string]int{}

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := sorter.ParseSharedStash(saveDir)
	characters, _ := sorter.ParseCharacters(saveDir)
	fmt.Printf("%d characters found\n", len(characters))

	for _, name := range d2s.UniqueNames {
		fmt.Printf("%s  \n", name)
	}

	for _, char := range characters {
		//fmt.Printf("[%s] %s (Level %d): %d Items\n", char.D2sCharacter.Header.Class, char.D2sCharacter.Header.Name, char.D2sCharacter.Header.Level, len(char.GetAllItems()))

		//fmt.Printf("\t%000d items in personal stash\n", len(char.PersonalStash.GetAllItems()))
		for _, item := range char.PersonalStash.GetAllItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
		fmt.Printf("\t%000d items in inventory\n", len(char.GetInventoryItems()))
		for _, item := range char.GetInventoryItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
		fmt.Printf("\t%000d items equipped on character\n", len(char.GetEquippedItems()))
		for _, item := range char.GetEquippedItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
		fmt.Printf("\t%000d items equipped on merc\n", len(char.GetMercItems()))
		for _, item := range char.GetMercItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
	}

	fmt.Printf("%d items in shared stash\n", len(sharedStash.GetAllItems()))
	for _, item := range sharedStash.GetAllItems() {
		//fmt.Printf("\t\t %s\n", item.GetName())
		addStatistic(item)
	}

	for key, value := range statistics {
		fmt.Printf("\"%s\";%d\n", key, value)
	}
}

func startWatcher(groups []sorter.ItemGroup) {
	saveDir, err := getSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	fmt.Printf("Starting Run Recorder\n")
	sharedStash := sorter.ParseSharedStash(saveDir)
	sorter.StartRunRecorder(&sharedStash, groups)

}
