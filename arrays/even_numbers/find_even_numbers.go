package even

func findEvenNumberOfDigits(nums []int) int {
	evens := 0
	for i := range nums {
		number := nums[i]
		if findNumDigits(number) % 2 == 0 {
			evens++
		}
	}
	return evens
}

func findNumDigits(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + findNumDigits(num/10)
}