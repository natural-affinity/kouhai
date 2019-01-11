package senpai

import (
	"os/exec"
	"time"
)

// Task to execute
type Task struct {
	Command  string
	Interval time.Duration
}

// Dispatch command and fetch results
func Dispatch(t *Task) (string, error) {
	command := exec.Command("sh", "-c", t.Command)
	out, err := command.CombinedOutput()

	if err != nil {
		return string(out), err
	}

	return string(out), nil
}
