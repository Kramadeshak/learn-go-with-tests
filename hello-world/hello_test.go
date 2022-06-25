package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func (t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T){
		got := Hello("Sarthak", "")
		want := "Hello, Sarthak"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is passed", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Sarthak", "Spanish")
		want := "Hola, Sarthak"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T){
		got := Hello("Sarthak", "French")
		want := "Bonjour, Sarthak"
		assertCorrectMessage(t, got, want)
	})
}