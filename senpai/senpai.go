package senpai

import (
	"bytes"
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
