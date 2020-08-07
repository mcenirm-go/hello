package hello_test

import (
	"bytes"
	"testing"

	"github.com/mcenirm-go/hello"
)

func TestHello(t *testing.T) {
	t.Run("it writes hello", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		hello.Hello(out)
		got := out.String()
		want := "hello\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
