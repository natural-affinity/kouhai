package main_test

import (
	"flag"
	"testing"

	"github.com/natural-affinity/gotanda"
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
		r := gotanda.CompareCommand(t, tc, update)
		r.Assert(t, tc)
	}
}
