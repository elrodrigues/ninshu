package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "flag"
	// "errors"
	// "strings"
	"github.com/elrodrigues/ninshu/com"
	pb "github.com/elrodrigues/ninshud/jutsu"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	tag      = "v0.1.5"
	help_msg = `Ninshu Client is a tool for interacting with your Ninshu network

Usage:

	ninshu <command> [arguments]

The commands are:

	anchor		create a Ninshu client
	configure	configure Ninshu client settings
	connect		connect to Ninshu network
	disconnect	disconnect from Ninshu network
	version		prints Ninshu version
	help		prints this help message or command info
	ping		pong

Use "ninshu help <command>" for more information about a command

`
)

var (
	addr = "localhost:47001" // fixed for now
)

func anchor(args []string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not dial-in to server: %v\n Is the daemon running?", err)
	}
	defer conn.Close()
	c := pb.NewClusterClient(conn)
	// Contact daemon
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch args[1] {
	case "drop":
		r, err := c.DropAnchor(ctx, &pb.ConnectRequest{HostIP: os.Getenv("hostIP"), Ip: ""})
		if err != nil {
			log.Fatalf("RPC failed: %v\n", err)
		}
		if r.Success {
			fmt.Println("Successfully dropped Ninshu anchor")
		} else {
			log.Fatalf("Failed to drop Ninshu anchor")
		}
	case "raise":
		r, err := c.RaiseAnchor(ctx, &pb.EmptyRequest{})
		if err != nil {
			log.Fatalf("RPC failed: %v\n", err)
		}
		if r.Success {
			fmt.Println("Successfully raised Ninshu anchor")
		} else {
			log.Fatalf("Failed to raise Ninshu anchor")
		}
	default:
		com.FetchHelp(args)
	}
}

func connectTo(ip string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not dial-in to server: %v\n Is the daemon running?", err)
	}
	defer conn.Close()
	c := pb.NewClusterClient(conn)
	// Contact daemon
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Printf("Host IP: %s\n", os.Getenv("hostIP"))
	fmt.Printf("Target IP: %s\n", ip)
	r, err := c.ConnectTo(ctx, &pb.ConnectRequest{HostIP: os.Getenv("hostIP"), Ip: ip})
	if err != nil {
		log.Fatalf("Failed to connect to Ninshu mesh: %v", err)
	} else if r.Success {
		fmt.Println(*r.Reply)
	} else {
		fmt.Println("Already connected to Ninshu mesh")
	}
}

func disconnect() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not dial-in to server: %v\n Is the daemon running?", err)
	}
	defer conn.Close()
	c := pb.NewClusterClient(conn)
	// Contact daemon
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	r, err := c.RaiseAnchor(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("RPC failed: %v\n", err)
	}
	if r.Success {
		fmt.Println("Successfully disconnected from Ninshu mesh")
	}
}

func main() {
	args := os.Args[1:] // ignore script location

	if len(args) < 1 {
		fmt.Fprint(os.Stderr, help_msg)
		os.Exit(2)
	}
	// load config file
	viper.SetDefault("ninRootPath", "")
	viper.SetDefault("hostIP", "127.0.0.1")
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
	os.Setenv("hostIP", viper.GetString("hostIP"))
	// parse command
	switch args[0] {
	case "anchor":
		if len(args) < 2 {
			com.FetchHelp(args)
		} else {
			anchor(args)
		}
	case "configure":
		fmt.Println("Configure Ninshu client and network")
	case "connect":
		if len(args) < 2 {
			com.FetchHelp(args)
		} else {
			fmt.Println("Attempting to connect to Ninshu network...")
			connectTo(args[1])
		}
	case "disconnect":
		disconnect()
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
	case "ping":
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not dial-in to server: %v\n Is the daemon running?", err)
		}
		defer conn.Close()
		c := pb.NewClusterClient(conn)
		// Contact daemon
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.PingNode(ctx, &pb.HelloRequest{Ping: args[1]})
		if err != nil {
			log.Fatalf("RPC failed: %v\n", err)
		}
		fmt.Println(r.Pong)
	default:
		fmt.Fprint(os.Stderr, help_msg)
		os.Exit(2)
	}
}
