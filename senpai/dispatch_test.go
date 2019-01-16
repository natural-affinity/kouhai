package senpai_test

import (
	"errors"
	"testing"

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
