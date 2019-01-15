package senpai_test

import (
	"errors"
	"testing"

	"github.com/natural-affinity/kouhai/senpai"
)

var dispatchTests = []struct {
	Name    string
	Command string
	Output  string
	Error   error
}{
	{"shell", "echo $0", "sh\n", nil},
	{"execute", "echo hello", "hello\n", nil},
	{"combined", "echo stdout; echo 1>&2 stderr", "stdout\nstderr\n", nil},
	{"failure", "fake-exe-path", "sh: fake-exe-path: command not found\n", errors.New("exit status 127")},
}

func TestDispatch(t *testing.T) {
	for _, tt := range dispatchTests {
		aout, aerr := senpai.Dispatch(tt.Command)

		out := (aout != tt.Output)
		err := invalidError(aerr, tt.Error)

		if out || err {
			t.Errorf("\nTest: %s\n %s\nExpected:\n %s %s\nActual:\n %s %s",
				tt.Name, tt.Command,
				tt.Output, tt.Error,
				aout, aerr)
		}
	}
}

func invalidError(actual error, expected error) bool {
	a := (actual != nil && expected != nil && actual.Error() != expected.Error())
	b := (actual == nil && expected != nil)

	return a || b
}
