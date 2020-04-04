// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"github.com/tanem/learn-go-with-tests/math/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
