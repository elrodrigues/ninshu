package main

import (
	"fmt"
	"os"
	// "flag"
	// "errors"
	// "strings"
)

var tag = "v0.1"
var help_msg = `Ninshu Client is a tool for interacting with your Ninshu network

Usage:

	ninshu <command> [arguments]

The commands are:

	configure	configure Ninshu client settings
	connect		connect to Ninshu network
	version		prints Ninshu version
	help		prints this help message

Use "ninshu help <command>" for more information about a command
`

func main() {
	args := os.Args[1:] // ignore script location

	if len(args) < 1 {
		fmt.Println(help_msg)
		os.Exit(2)
	}
	switch args[0] {
	case "configure":
		fmt.Println("Configure Ninshu client and network")
	case "connect":
		fmt.Println("Connect to Ninshu network")
	case "version":
		fmt.Printf("Ninshu %s へようこそ\n", tag)
	case "help":
		fmt.Println(help_msg)
	case "tskr":
		fmt.Println(help_msg)
	default:
		fmt.Println(help_msg)
		os.Exit(2)
	}
}
