package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lyle")

	got := buffer.String()
	want := "Hello, Lyle"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}