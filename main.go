package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	docopt "github.com/docopt/docopt-go"
)

// Version identifier
const Version = "0.0.1"

// Usage message (docopt interface)
const Usage = `
  Kouhai
    Run tasks at periodic intervals

  Usage:
    kouhai [--interval n] <cmd>
    kouhai  --help
    kouhai  --version

  Options:
    -h, --help  	  display help information	
	-v, --version  	  display version information
	-i, --interval n  set refresh interval [default: 1s]

`

func main() {
	args, _ := docopt.ParseArgs(Usage, os.Args[1:], Version)
	_, err = time.ParseDuration(args["--interval"].(string))
	cmd := args["<cmd>"].(string)
	fmt.Println(args)

	Execute(cmd)

	//t := Task{Command = cmd, Interval = interval}
	//fmt.Printf("%+v", t)
	//fmt.Println(err)

	/*
		for {
			cmd := exec.Command("sh", "-c", args["<cmd>"].(string))
			out, err := cmd.CombinedOutput()

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			fmt.Printf("%s", out)
			time.Sleep(interval)
		}*/
}

// Execute command
func Execute(cmd string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	command := exec.Command("bash", "-c", cmd)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())
	}
	fmt.Println(stdout.String())

	stdout.Reset()
	stderr.Reset()
}
