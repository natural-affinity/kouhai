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
	{"invokeBash", "echo $0", "sh\n", nil},
	{"executeCmd", "echo hello", "hello\n", nil},
	{"combinedOutput", "echo stdout; echo 1>&2 stderr", "stdout\nstderr\n", nil},
	{"commandError", "fake-exe-path", "sh: fake-exe-path: command not found\n", errors.New("exit status 127")},
}

func TestDispatch(t *testing.T) {
	for _, tt := range dispatchTests {
		aout, aerr := senpai.Dispatch(tt.Command)

		out := (aout != tt.Output)
		err := (aerr != nil && tt.Error != nil && aerr.Error() != tt.Error.Error()) || (aerr == nil && tt.Error != nil)

		if out || err {
			t.Errorf("\nTest: %s\n %s\nExpected:\n %s %s\nActual:\n %s %s",
				tt.Name, tt.Command,
				tt.Output, tt.Error,
				aout, aerr)
		}
	}
}
