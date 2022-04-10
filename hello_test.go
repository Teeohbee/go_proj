package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Toby")
	want := "Hello, Toby"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
