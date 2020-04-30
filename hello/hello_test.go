package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrentMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	
	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("lyle", "")
		want := "Hello,lyle"

		assertCorrentMessage(t, got, want)
	})

	t.Run("say 'hello,world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello,world"
		assertCorrentMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola,Elodie"
		assertCorrentMessage(t, got, want)
	})
	
	t.Run("in French", func(t *testing.T) {
		got := Hello("Dushang", "French")
		want := "Bonjour,Dushang"
		assertCorrentMessage(t, got, want)
	})
}