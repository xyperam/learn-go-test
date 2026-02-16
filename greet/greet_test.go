package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ilham")

	got := buffer.String()
	if got != "Hello, Ilham" {
		t.Errorf("got %q want %q", got, "Hello, Ilham")
	}
}
