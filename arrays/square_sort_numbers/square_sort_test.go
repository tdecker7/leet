package square_sort_numbers

import "testing"

func TestSquareSort(t *testing.T) {
	t.Run("scenario 1", func(t *testing.T) {
		input := []int{-4,-1,0,3,10}
		expected := []int{0,1,9,16,100}
		output := squareSort(input)
		for i, v := range output {
			if v != expected[i] {
				t.Errorf("expected: %v, got %v", expected, output)
			}
		}
	})
}