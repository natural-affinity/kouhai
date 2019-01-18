package senpai_test

import (
	"errors"
	"testing"
	"time"

	capturer "github.com/kami-zh/go-capturer"
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
		Times   int
		Task    *senpai.Task
	}{
		{
			"stop", false, "sh: fake-exe: command not found\n", errors.New("exit status 127"),
			2, &senpai.Task{Command: "fake-exe", Stop: true, Interval: 1 * time.Millisecond},
		},
		{
			"delay", false, "finished monitoring", nil,
			4, &senpai.Task{Command: "echo hello", Stop: false, Interval: 20 * time.Millisecond},
		},
		{
			"print", true, "hello\nhello\n", nil,
			2, &senpai.Task{Command: "echo hello", Stop: false, Interval: 1 * time.Millisecond},
		},
	}

	for _, tc := range cases {
		var actualOutput string
		var actualError error

		forever := func() bool {
			if tc.Times > 0 {
				tc.Times--
				return true
			}

			return false
		}

		start := time.Now()
		if tc.Capture {
			actualOutput = capturer.CaptureStdout(func() {
				_, actualError = tc.Task.Monitor(forever)
			})
		} else {
			actualOutput, actualError = tc.Task.Monitor(forever)
		}
		elapsed := time.Since(start)

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
