package maxConsecutive

import (
	"testing"
)

func TestMaxConsecutive(t *testing.T) {
	t.Run("three in the middled", func(t *testing.T) {
		threeConsecInMiddle := []int{1,1,0,1,1,1,0}
		consec := findMaxConsecutiveOnes(threeConsecInMiddle)
		if consec != 3 {
			t.Errorf("Expected %d, got %d\n", 3, consec)
		}
	})
	t.Run("two in middle", func(t *testing.T) {
		two := []int{1,0,0,0,1,1,0}
		consec := findMaxConsecutiveOnes(two)
		if consec != 2 {
			t.Errorf("Expected %d, got %d\n", 2, consec)
		}
	})
}