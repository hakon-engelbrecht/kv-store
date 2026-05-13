package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/hakon-engelbrecht/kv-store/internal/command"
	"github.com/hakon-engelbrecht/kv-store/internal/server"
	"github.com/hakon-engelbrecht/kv-store/internal/store"
)

const defaultPort = 6380

var dataStore store.Store

func main() {
	args := os.Args[1:]
	port, err := parsePort(args)
	if err != nil {
		log.Fatal(err)
	}

	dataStore = store.NewConcurrentStore()

	srv := server.NewTCPServer(port, handleConnection)
	err = srv.Listen()
	if err != nil {
		panic(err)
	}
}

// Parses the command line arguments to find, if a port has been configured
func parsePort(args []string) (int, error) {
	if len(args) == 0 {
		return defaultPort, nil // return the default port
	}

	if len(args) > 2 {
		return 0, fmt.Errorf("invalid number of arguments: %v", len(args))
	}

	command := args[0]
	switch command {
	case "-h":
		fmt.Println("Usage: kv-store [--port <port>]")
		os.Exit(0)
	case "--port":
		if len(args) != 2 {
			return 0, fmt.Errorf("invalid number of arguments: %v", len(args))
		}
		portArg := args[1]
		port, err := strconv.Atoi(portArg)
		if err != nil {
			return 0, fmt.Errorf("invalid port argument: %v", portArg)
		}
		return port, nil
	default:
		return 0, fmt.Errorf("unknown argument")
	}
	return defaultPort, nil
}

func handleConnection(conn net.Conn) {
	log.Printf("serving %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		log.Printf("received message: %s", msg)
		cmd, err := command.ParseCommand(msg)
		if err != nil {
			_, err = conn.Write([]byte("invalid message\n"))
			if err != nil {
				log.Println("write error: ", err)
				break
			}
		} else {
			res := cmd.Execute(dataStore)
			_, err = conn.Write([]byte(res + "\n"))
			if err != nil {
				log.Println("write error: ", err)
				break
			}

			if cmd.Name == "QUIT" {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("read error: ", err)
	} else {
		log.Println("client disconnected")
	}
}
