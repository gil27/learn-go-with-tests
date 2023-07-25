package main

import "testing"

func TestRepeat(t *testing.T) {
  t.Run("repeat 5 times", func(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    AssertReturn(t, expected, repeated)
  })

  t.Run("reapat 0 times", func(t *testing.T) {
    repeated := Repeat("a", 0)
    expected := ""

    AssertReturn(t, expected, repeated)
  })
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 5)
  }
}

func AssertReturn(t testing.TB, expected, repeated string) {
  t.Helper()

  if repeated != expected {
    t.Errorf("expected %q but got %q", expected, repeated)
  }
}
