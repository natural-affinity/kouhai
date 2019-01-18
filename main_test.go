package main_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/natural-affinity/kouhai/spec"
)

var update = flag.Bool("update", false, "update .golden files")

func TestUsage(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{"help.long"},
		{"help.short"},
		{"version.long"},
		{"version.short"},
		{"stop.short"},
		{"stop.long"},
		{"invalid.interval.short"},
		{"invalid.interval.long"},
	}

	for _, tc := range cases {
		_, command := spec.LoadTestFile(t, "testdata", tc.Name+".input")
		golden, expected := spec.LoadTestFile(t, "testdata", tc.Name+".golden")
		aout, _ := spec.Run(string(command))

		if *update {
			ioutil.WriteFile(golden, aout, 0644)
		}

		expected, _ = ioutil.ReadFile(golden)
		out := !bytes.Equal(aout, expected)

		if out {
			t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, aout, expected)
		}
	}
}
