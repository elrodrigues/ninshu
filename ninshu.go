package main

import (
	"fmt"
	"os"

	// "flag"
	// "errors"
	// "strings"
	"github.com/elrodrigues/ninshu/com"
	"github.com/spf13/viper"
)

const (
	tag        = "v0.1"
	configPath = "~/.ninshu/ninshu.json"
	help_msg   = `Ninshu Client is a tool for interacting with your Ninshu network

Usage:

	ninshu <command> [arguments]

The commands are:

	anchor		create a Ninshu client
	configure	configure Ninshu client settings
	connect		connect to Ninshu network
	version		prints Ninshu version
	help		prints this help message or command info

Use "ninshu help <command>" for more information about a command

`
)

func main() {
	args := os.Args[1:] // ignore script location

	if len(args) < 1 {
		fmt.Fprint(os.Stderr, help_msg)
		os.Exit(2)
	}
	// load config file
	viper.SetDefault("ninRootPath", "")
	viper.SetConfigName("ninshu")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.ninshu")
	// viper.AddConfigPath("/etc/ninshu/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("Fatal error reading ninshu config file: %w\n", err))
		}
	}
	os.Setenv("ninRootPath", viper.GetString("ninRootPath"))
	// parse command
	switch args[0] {
	case "configure":
		fmt.Println("Configure Ninshu client and network")
	case "connect":
		fmt.Println("Connect to Ninshu network")
	case "version":
		fmt.Printf("Ninshu %s へようこそ\n", tag)
	case "help":
		if len(args) < 2 {
			fmt.Print(help_msg)
		} else {
			com.FetchHelp(args[1:])
		}
	case "tskr":
		fmt.Print(help_msg)
	default:
		fmt.Fprint(os.Stderr, help_msg)
		os.Exit(2)
	}
}
