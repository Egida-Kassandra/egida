package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/antonioalfa22/egida/internal/config"
	"github.com/antonioalfa22/egida/internal/dsl"
	"github.com/antonioalfa22/egida/internal/info"
	"github.com/antonioalfa22/egida/internal/menu"
	"os"
)

var parser *argparse.Parser

func main() {
	parser = argparse.NewParser("Egida CLI", "")
	// Commands
	menuCmd := parser.NewCommand("menu", "")
	addGroupCmd := parser.NewCommand("add-group", "")
	infoCmd := parser.NewCommand("info", "")
	compileCmd := parser.NewCommand("compile", "")
	// Flags
	hostsgroup := parser.String("g", "group",
		&argparse.Options{Required: false, Help: "Host group: -group local"})
	hostslist := parser.StringList("H", "hosts",
		&argparse.Options{Required: false, Help:"List of hosts: -H 192.128.2.1 --hosts localhost -H 129.1.1.1"})
	services := parser.Selector("s", "services", []string{"all", "running", "stopped"},
		&argparse.Options{Required: false, Help:"Services info (all | running | stopped): -services all"})
	packages := parser.Selector("p", "packages", []string{"all"} ,
		&argparse.Options{Required: false, Help:"Packages info (all): -packages all"})
	hardening := parser.Selector("z", "hardscores", []string{"lynis"},
		&argparse.Options{Required: false, Help:"Hardening scores info (lynis): -hardscores lynis"})

	err := parser.Parse(os.Args)
	if err != nil {
		invalidArgs()
		return
	}
	if menuCmd.Happened() {
		setMenu()
	} else if compileCmd.Happened() {
		setCompile()
	} else if addGroupCmd.Happened() {
		setAddGroup(*hostsgroup, *hostslist)
	} else if infoCmd.Happened() {
		setInfo(*hostslist, *services, *packages, *hardening)
	} else {
		invalidArgs()
	}
}

// EGIDA OPTIONS
func invalidArgs()  {
	fmt.Print(parser.Usage(nil))
}

func setMenu() {
	menu.SelectHardeningMode()
}

func setCompile() {
	file := os.Args[2]
	dsl.ParseFile(file)
}

func setAddGroup(hostsgroup string, hostslist []string) {
	if hostsgroup != "" && len(hostslist) != 0 {
		config.AddHostGroup(hostsgroup, hostslist)
	} else {
		invalidArgs()
	}
}

func setInfo(hostslist []string, services string, packages string, hardening string) {
	if len(hostslist) != 0 {
		info.GetWorkerInfo(hostslist, services, packages, hardening)
	} else {
		invalidArgs()
	}
}