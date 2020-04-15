package poker

import (
	"os"
)

// Tape implements "when we write we go from the beginning".
type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, 0)

	return t.File.Write(p)
}
