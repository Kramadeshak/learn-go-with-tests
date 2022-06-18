package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sarthak")
	want := "Hello, Sarthak"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}