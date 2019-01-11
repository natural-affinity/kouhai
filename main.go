package main

import (
	"bytes"
	"fmt"
	"log"
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

// Task to execute
type Task struct {
	Command  string
	Interval time.Duration
}

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and fetch args
	args, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	// extract options and args
	cmd := args["<cmd>"].(string)
	interval, err := time.ParseDuration(args["--interval"].(string))
	if err != nil {
		log.Fatalf("invalid interval: %s", err.Error())
	}

	// build and execute task
	t := &Task{Command: cmd, Interval: interval}
	for {
		out, err := Execute(t)
		if err != nil {
			log.Fatalf("invalid command: %s", err.Error())
		}

		fmt.Printf(out)
		time.Sleep(interval)
	}

}

// Execute command and fetch results
func Execute(t *Task) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	defer stdout.Reset()
	defer stderr.Reset()

	command := exec.Command("sh", "-c", t.Command)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}
