package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello World"
	got := Hello()
	checkIfMatching(t,got,want)
}

func TestHelloNameNew(t *testing.T) {
	t.Parallel()
	t.Run("Handling empty string case",func(t *testing.T) {
		got := HelloName("")
		want := "Hello World"
		checkIfMatching(t,got,want)
	})
	t.Run("If the input is provided",func(t *testing.T) {
		got := HelloName("Anish")
		want := "Hello Anish"
		checkIfMatching(t,got,want)
	})
}

func checkIfMatching(t *testing.T,got string, want string){
	t.Helper()
	if got!=want{
		t.Error("Error: The output didn't matched!")
		t.Errorf("expected: %s, got %s", want, got)
	}
}