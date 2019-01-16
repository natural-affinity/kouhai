package senpai_test

import (
	"errors"
	"testing"
	"time"

	"github.com/natural-affinity/kouhai/senpai"
	"github.com/natural-affinity/kouhai/spec"
)

func TestDispatch(t *testing.T) {
	cases := []struct {
		Name string
		Out  string
		Err  error
		Task *senpai.Task
	}{
		{
			"shell", "sh\n", nil,
			&senpai.Task{Command: "echo $0"},
		},
		{
			"execute", "hello\n", nil,
			&senpai.Task{Command: "echo hello"},
		},
		{
			"combined", "stdout\nstderr\n", nil,
			&senpai.Task{Command: "echo stdout; echo 1>&2 stderr"},
		},
		{
			"failure", "sh: fake-exe: command not found\n", errors.New("exit status 127"),
			&senpai.Task{Command: "fake-exe"},
		},
	}

	for _, tc := range cases {
		actualOutput, actualError := tc.Task.Dispatch()

		out := (actualOutput != tc.Out)
		err := spec.IsInvalidError(actualError, tc.Err)

		if out || err {
			t.Errorf("\nTest: %s\n %s\nExpected:\n %s %s\nActual:\n %s %s",
				tc.Name, tc.Task.Command,
				tc.Out, tc.Err,
				actualOutput, actualError)
		}
	}
}

func TestMonitor(t *testing.T) {
	cases := []struct {
		Name    string
		Capture bool
		Out     string
		Err     error
		Task    *senpai.Task
	}{
		{
			"stop", false, "sh: fake-exe: command not found\n", errors.New("exit status 127"),
			&senpai.Task{Command: "fake-exe", Stop: true, Times: 2, Interval: 1 * time.Millisecond},
		},
		{
			"times", false, "monitored: 4 times", nil,
			&senpai.Task{Command: "echo hello", Stop: false, Times: 4, Interval: 1 * time.Millisecond},
		},
		{
			"delay", false, "monitored: 2 times", nil,
			&senpai.Task{Command: "echo hello", Stop: false, Times: 2, Interval: 50 * time.Millisecond},
		},
		{
			"print", true, "hello\nhello\n", nil,
			&senpai.Task{Command: "echo hello", Stop: false, Times: 2, Interval: 1 * time.Millisecond},
		},
	}

	for _, tc := range cases {
		capture := &spec.Snapshot{}
		if tc.Capture {
			capture.Start()
		}

		start := time.Now()
		actualOutput, actualError := tc.Task.Monitor()
		elapsed := time.Since(start)

		if tc.Capture {
			capture.Copy()
			capture.Release()
			actualOutput = capture.Out
		}

		out := (actualOutput != tc.Out)
		err := spec.IsInvalidError(actualError, tc.Err)
		dur := (elapsed < tc.Task.Interval)

		if out || err {
			t.Errorf("\nTest: %s\n %s\nExpected:\n %s %s\nActual:\n %s %s",
				tc.Name, tc.Task.Command,
				tc.Out, tc.Err,
				actualOutput, actualError)
		}

		if dur {
			t.Errorf("\nTest: %s \nExecution Time Error\nExpected: %s\nActual:%s\n",
				tc.Name, tc.Task.Interval, elapsed)
		}
	}
}
