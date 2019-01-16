package spec

import (
	"bytes"
	"io"
	"os"
)

// IsInvalidError compares errors
func IsInvalidError(actual error, expected error) bool {
	a := (actual != nil && expected != nil && actual.Error() != expected.Error())
	b := (actual == nil && expected != nil)

	return a || b
}

// Capture stdout
type Capture interface {
	Start()
	Copy()
	Release()
}

// Snapshot stdout
type Snapshot struct {
	stdout, r, w *os.File
	c            chan string
	Out          string
}

// Start stdout capture
func (s *Snapshot) Start() {
	s.stdout = os.Stdout
	s.r, s.w, _ = os.Pipe()
	os.Stdout = s.w
}

// Copy stdout
func (s *Snapshot) Copy() {
	s.c = make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, s.r)
		s.c <- buf.String()
	}()
}

// Release capture resources
func (s *Snapshot) Release() {
	s.w.Close()
	os.Stdout = s.stdout
	s.Out = <-s.c
}
