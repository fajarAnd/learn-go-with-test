package hello_world

import (
	"testing"
)

const prefix = "Hello,"

// Hello returns a greeting.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying Hello To People", func(t *testing.T) {
		got := Hello("Fajar")
		want := "Hello,Fajar"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello,World"

		assertCorrectMessage(t, got, want)
	})
}
