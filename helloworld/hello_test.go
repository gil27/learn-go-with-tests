package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("in French", func(t *testing.T) {
    got := Hello("Gil", "French")
    want := "Bonjour, Gil"

    AssertCorrectMessage(t, got, want)
  })

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Gil", "Spanish")
    want := "Holla, Gil"

    AssertCorrectMessage(t, got, want)

  })

  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Gil", "")
    want := "Hello, Gil"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })
  t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
    got := Hello("Gil", "English")
    want := "Hello, Gil"

    AssertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"

    AssertCorrectMessage(t, got, want)
  })
}

func AssertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
