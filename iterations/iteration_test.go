package iteration

import "testing"

func TestRepeat(t *testing.T) {
	printerHelper := func(t testing.TB, expected, repeated string) {
		t.Helper()

		if repeated != expected {
			t.Errorf("Expeted %q but got %q", expected, repeated)
		}
	}

	t.Run("Printing 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		printerHelper(t, expected, repeated)
	})

	t.Run("Printing 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		printerHelper(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
