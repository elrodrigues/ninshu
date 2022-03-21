package com

import (
	"fmt"
	"os"
)

// root usage message
var usage = `Usage:

	ninshu help [command]

The commands are:

	configure	configure Ninshu client settings
	connect		connect to Ninshu network
	version		prints Ninshu version
	help		prints this help message or command info
`

var pathmsg = `Please configure Ninshu's root path!
You can do this by setting ninRootPath in ninshu.json (usually in ~/.ninshu/)
This path must be an absolute/full path.
`

func check(e error, cmd string) {
	if e != nil {
		fmt.Printf("Cannot find help page for %s\n", cmd)
		os.Exit(2)
	}
}

// FetchHelp fetches the relevant documentation for a command.
//
// Ninshu's `ninDocPath' environment variable must be set to the
// documentation's directory for this function to work correctly.
func FetchHelp(args []string) {
	path := os.Getenv("ninRootPath")
	if path == "" {
		fmt.Println(pathmsg)
		os.Exit(2)
	}
	switch args[0] {
	case "help":
		fmt.Println(usage)
	case "tskr":
		fmt.Println("Sure bud")
	default:
		filePath := fmt.Sprintf("%s/docs/%s.doc", path, args[0])
		dat, err := os.ReadFile(filePath)
		check(err, args[0])
		fmt.Print(string(dat))
	}
}
