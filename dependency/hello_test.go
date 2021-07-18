package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Pratap", "")
		want := "Hello Pratap"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Cuban", "Spanish")
		want := "Hola Cuban"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Cuban", "French")
		want := "Bonjour Cuban"
		assertCorrectMessage(t, got, want)

	})

}

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello Chris"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}
