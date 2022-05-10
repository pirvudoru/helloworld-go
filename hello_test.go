package main

import "testing"

func TestSayHello(t *testing.T) {
	t.Run("saying hello to people	", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertEqual(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
