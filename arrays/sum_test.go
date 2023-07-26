package main

import "testing"

func TestSum(t *testing.T) {
  t.Run("Sum collections of 5 items", func(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    AssertSum(t, got, want, numbers)
  })

  t.Run("Sum collections of any size", func(t *testing.T) {
    numbers := []int{1, 2, 3}

    got := Sum(numbers)
    want := 6

    AssertSum(t, got, want, numbers)
  })
}

func AssertSum(t testing.TB, got, want int, numbers []int) {
  t.Helper()

  if got != want {
    t.Errorf("got %d want %d given, %v", got, want, numbers)
  }
}
