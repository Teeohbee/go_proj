package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Toby", "")
		want := "Hello, Toby"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean-Luc", "French")
		want := "Bonjour, Jean-Luc"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Miyakazi-san", "Japanese")
		want := "Konnichiwa, Miyakazi-san"

		assertCorrectMessage(t, got, want)
	})
}
