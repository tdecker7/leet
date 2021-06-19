package maxConsecutive 

// import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
    maxConsecutive := 0
    currentConsecutive := 0
	for i := range nums {
		curr := nums[i]
		if curr == 1 {
			currentConsecutive++
		} else if curr == 0 {
			currentConsecutive = 0
		}
		if currentConsecutive > maxConsecutive {
			maxConsecutive = currentConsecutive
		}
	}
    return maxConsecutive
}