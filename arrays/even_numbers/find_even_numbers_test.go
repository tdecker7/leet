package even

import "testing"

func TestFindEvenNumbers(t *testing.T) {
	t.Run("one even number", func(t *testing.T) {
		nums := []int{12,345,2,6,7896}
		evens := findEvenNumberOfDigits(nums)
		expected := 2
		if evens != expected {
			t.Errorf("expected %d, got %d\n", expected, evens)
		}
	})
}