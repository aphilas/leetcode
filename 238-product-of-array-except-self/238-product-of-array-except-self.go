func productExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))

	// Build lefts
	left := 1
	for i, n := range nums {
		answer[i] = left
		left *= n
	}

	// Build rights + answer
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}