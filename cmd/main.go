package main

import (
	"flag"
	"fmt"
	"github.com/antonioalfa22/egida/internal/config"
	"github.com/antonioalfa22/egida/internal/dsl"
	"github.com/antonioalfa22/egida/internal/menu"
	"os"
)

func main() {
	var connection string
	var hostsgroup string
	var hostslist string
	flag.StringVar(&connection, "connection", "local", "Connection type (default local): local | ssh")
	flag.StringVar(&hostsgroup, "group", "", "Host group")
	flag.StringVar(&hostslist, "hosts", "", "List of hosts, ej: 192.128.2.1,localhost,129.1.1.1")
	mode := os.Args[1]
	if mode == "menu" {
		menu.SelectHardeningMode()
	} else if mode == "compile" {
		file := os.Args[2]
		dsl.ParseFile(file)
	} else if mode == "add-group" {
		fmt.Println(hostslist,hostsgroup)
		if hostsgroup != "" && hostslist != "" {
			config.AddHostGroup(hostsgroup, hostslist)
		} else {
			fmt.Println("Command not found, Try: egida add-group -group=example -hosts=192.128.2.1,localhost")
		}
	} else if mode == "info" {

	} else {
		invalidArgs()
	}
}

func invalidArgs() {
	fmt.Println("EGIDA Mode invalid -> [menu | compile | add-group | info]")
}
