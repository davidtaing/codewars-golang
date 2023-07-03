package exercises

import "testing"

func TestSmallestItemInArray(t *testing.T) {
	t.Run("array = [34, 15, 88, 2]", func(t *testing.T) {
		got := SmallestItemInArray([]int{34, 15, 88, 2})
		want := 2

		if got != want {
			t.Errorf("expected '%d', got '%d", want, got)
		}
	})

	t.Run("array = [34, -345, -1, 100]", func(t *testing.T) {
		got := SmallestItemInArray([]int{34, -345, -1, 100})
		want := -345

		if got != want {
			t.Errorf("expected '%d', got '%d", want, got)
		}
	})
}
