package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello World"
	got := Hello()
	if got != want {
		t.Errorf("expected: %s, got %s", want, got)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Anish")
	want := "Hello, Anish"
	if got != want {
		t.Errorf("expected: %s, got %s", want, got)
	}
}

