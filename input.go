package terminal

import (
	"fmt"
	"io"
	"os"
)

var (
	Stdin = &Input{
		reader:      os.Stdin,
		interactive: IsTerminal(os.Stdin) && IsTerminal(os.Stdout),
	}
)

// Scan scans text read from standard input, storing successive
// space-separated values into successive arguments. Newlines count
// as space. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
func Scan(a ...interface{}) (n int, err error) {
	return fmt.Fscan(Stdin, a...)
}

// Scanln is similar to Scan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Scanln(a ...interface{}) (n int, err error) {
	return fmt.Fscanln(Stdin, a...)
}

// Scanf scans text read from standard input, storing successive
// space-separated values into successive arguments as determined by
// the format. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
// Newlines in the input must match newlines in the format.
// The one exception: the verb %c always scans the next rune in the
// input, even if it is a space (or tab etc.) or newline.
func Scanf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fscanf(Stdin, format, a...)
}

type Input struct {
	reader      io.Reader
	interactive bool
}

func (input *Input) Fd() uintptr {
	if fd, ok := input.reader.(FdHolder); ok {
		return fd.Fd()
	}

	return 0
}

func NewInput(reader io.Reader) *Input {
	return &Input{
		reader:      reader,
		interactive: IsTerminal(reader),
	}
}

func (input *Input) Read(p []byte) (int, error) {
	return input.reader.Read(p)
}

func (input *Input) IsInteractive() bool {
	return input.interactive
}

func (input *Input) SetInteractive(interactive bool) {
	input.interactive = interactive
}

func (input *Input) SetReader(r io.Reader) {
	input.reader = r
}
