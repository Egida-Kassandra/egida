package main

import (
	"flag"
	"fmt"
	"github.com/antonioalfa22/egida/internal/config"
	"github.com/antonioalfa22/egida/internal/dsl"
	"github.com/antonioalfa22/egida/internal/info"
	"github.com/antonioalfa22/egida/internal/menu"
	"os"
)

func main() {
	hostsgroup := flag.String("group", "", "Host group")
	hostslist := flag.String("hosts", "", "List of hosts, ej: 192.128.2.1,localhost,129.1.1.1")
	services := flag.String("services", "", "Services info: all | running | stopped")
	packages := flag.String("packages", "", "Packages info: all")
	hardening := flag.String("hardscores", "", "Hardening scores info: lynis")
	//connection := flag.String("connection", "local", "Connection type (default local): local | ssh")
	if len(os.Args) <= 1 { invalidArgs() }
	mode := os.Args[1]
	if mode == "menu" {
		setMenu()
	} else if mode == "compile" {
		setCompile()
	} else if mode == "add-group" {
		setAddGroup(*hostsgroup, *hostslist)
	} else if mode == "info" {
		setInfo(*hostslist, *services, *packages, *hardening)
	} else {
		invalidArgs()
	}
}

func invalidArgs() {
	fmt.Println("EGIDA Mode invalid -> [menu | compile | add-group | info]")
}

// EGIDA OPTIONS

func setMenu() {
	menu.SelectHardeningMode()
}

func setCompile() {
	file := os.Args[2]
	dsl.ParseFile(file)
}

func setAddGroup(hostsgroup string, hostslist string) {
	if hostsgroup != "" && hostslist != "" {
		config.AddHostGroup(hostsgroup, hostslist)
	} else {
		fmt.Println("Command not found, Try: egida add-group -group=example -hosts=192.128.2.1,localhost")
	}
}

func setInfo(hostslist string, services string, packages string, hardening string) {
	if hostslist != "" && checkInfoArgs(services, packages, hardening){
		info.GetWorkerInfo(hostslist, services, packages, hardening)
	} else {
		fmt.Println("Command not found, Try: egida info -services=running -packages=all -hardscore=lynis -hosts=192.128.2.1,localhost")
	}
}

func checkInfoArgs(ser string, pack string, har string) bool {
	val := ser != "" || pack != "" || har != ""
	if ! val { return false }
	if ser != "" && !(ser == "all" || ser == "running" || ser == "stopped") { return false }
	if pack != "" && !(pack == "all") { return false }
	if har != "" && !(har == "lynis") { return false }
	return true
}