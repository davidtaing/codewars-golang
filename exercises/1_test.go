package exercises

import "testing"

func TestDoubleInteger(t *testing.T) {
	t.Run("input 2, output 4", func(t *testing.T) {
		got := DoubleInteger(2)
		want := 4

		if got != want {
			t.Errorf("expected '%d', got '%d'", want, got)
		}
	})

	t.Run("input 12, output 24", func(t *testing.T) {
		got := DoubleInteger(2)
		want := 4

		if got != want {
			t.Errorf("expected '%d', got '%d'", want, got)
		}
	})
}
