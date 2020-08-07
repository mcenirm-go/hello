package hello

import (
	"fmt"
	"io"
	"os"
)

// Hello prints "hello" to out.
func Hello(out io.Writer) {
	fmt.Fprintln(out, "hello")

}

// Main is a wrapper for Hello.
func Main() {
	Hello(os.Stdout)
}
