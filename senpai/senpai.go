package senpai

import (
	"fmt"
	"os/exec"
	"time"
)

// Task to execute
type Task struct {
	Stop     bool
	Times    int
	Command  string
	Interval time.Duration
}

// Senpai monitors a task periodically
type Senpai interface {
	Monitor() (string, error)
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
func (t *Task) Monitor() (string, error) {
	i := 0
	if t.Times == 0 {
		i = -1
	}

	for i < t.Times {
		out, err := t.Dispatch()
		if err != nil && t.Stop {
			return out, err
		}

		if t.Times > 0 {
			i = i + 1
		}

		fmt.Printf(out)
		time.Sleep(t.Interval)
	}

	return fmt.Sprintf("monitored: %d times", i), nil
}
