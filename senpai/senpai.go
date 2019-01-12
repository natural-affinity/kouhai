package senpai

import (
	"fmt"
	"os/exec"
	"time"
)

// Task to execute
type Task struct {
	Stop     bool
	Command  string
	Interval time.Duration
}

// Dispatch command and fetch results
func Dispatch(cmd string) (string, error) {
	command := exec.Command("sh", "-c", cmd)
	out, err := command.CombinedOutput()

	if err != nil {
		return string(out), err
	}

	return string(out), nil
}

// Monitor task and
func (t *Task) Monitor() (string, error) {
	for {
		out, err := Dispatch(t.Command)
		if err != nil && t.Stop {
			return out, err
		}

		fmt.Printf(out)
		time.Sleep(t.Interval)
	}
}
