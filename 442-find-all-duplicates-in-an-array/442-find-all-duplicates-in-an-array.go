func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1

		if nums[idx] < 0 {
			res = append(res, idx+1)
		} else {
			nums[idx] = -nums[idx]
		}
	}

	return res
}