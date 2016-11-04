package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// build vars
var (
	// Version is set with a linker flag (see Makefile)
	Version string
	// Build is set with a linker flag (see Makefile)
	Build string
	//Global test variables
)

func main() {
	args := os.Args
	testName := ""
	serviceSwarm := false
	amplifierSwarm := false
	if len(args) > 1 {
		for i, arg := range args {
			if i > 0 {
				if arg == "--service-swarm" {
					serviceSwarm = true
				} else if arg == "--amplifier-swarm" {
					amplifierSwarm = true
				} else {
					testName = arg
				}
			}
		}
	}
	testargs := []string{"test", "-v", "./api/rpc/tests"}
	if testName != "" {
		testName = strings.ToUpper(testName[0:1]) + strings.ToLower(testName[1:])
		testargs = append(testargs, "-run")
		testargs = append(testargs, testName+"*")
		fmt.Printf("execute tests: %s\n", testName)
	} else {
		fmt.Println("execute all tests")
	}
	if serviceSwarm {
		os.Setenv("endpoints", "etcd:2379")
		os.Setenv("elasticsearchURL", "elasticsearch:9200")
		os.Setenv("natsURL", "nats:4222")
		os.Setenv("influxURL", "influx:8086")
	}
	if amplifierSwarm {
		os.Setenv("server_address", "amplifier")
	}
	goPath := os.Getenv("GOPATH")
	os.Chdir(path.Join(goPath, "src", "github.com", "appecelerator", "amp"))
	cmd := exec.Command("go", testargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Tests execution exit with error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Tests execution exit without error\n")
		os.Exit(0)
	}
}
