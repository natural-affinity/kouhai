package main

import (
	"fmt"
	"log"
	"os"
	"time"

	docopt "github.com/docopt/docopt-go"
	"github.com/natural-affinity/kouhai/senpai"
)

// Version identifier
const Version = "0.0.1"

// Usage message (docopt interface)
const Usage = `
  Kouhai (kōhai)
    Run tasks at periodic intervals

  Usage:
    kouhai [--interval n] [--stop] <cmd>
    kouhai  --help
    kouhai  --version

  Options:
	-s, --stop        stop on error
    -h, --help  	  display help information	
	-v, --version  	  display version information
	-i, --interval n  set refresh interval [default: 1s]
`

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and collect args
	args, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	// extract options and args
	cmd := args["<cmd>"].(string)
	stop := args["--stop"].(bool)
	interval, err := time.ParseDuration(args["--interval"].(string))
	if err != nil {
		log.Fatalf("invalid interval: %s", err.Error())
	}

	// build and execute task
	task := &senpai.Task{Command: cmd, Interval: interval}
	for {
		out, err := senpai.Dispatch(task)
		if err != nil && stop {
			log.Fatalf("%s\n", out)
		}

		fmt.Printf(out)
		time.Sleep(interval)
	}
}
