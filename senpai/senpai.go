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

// Senpai monitors a task periodically
type Senpai interface {
	Monitor(forever func() bool) (string, error)
	Dispatch() (string, error)
}

// Dispatch command
func (t *Task) Dispatch() (string, error) {
	command := exec.Command("sh", "-c", t.Command)
	out, err := command.CombinedOutput()

	if err != nil {
		return string(out), err
	}

	return string(out), nil
}

// Monitor task
func (t *Task) Monitor(forever func() bool) (string, error) {
	for forever() {
		out, err := t.Dispatch()
		if err != nil && t.Stop {
			return out, err
		}

		fmt.Printf(out)
		time.Sleep(t.Interval)
	}

	return "finished monitoring", nil
}
