func twoSum(nums []int, target int) []int {
    diff := make(map[int]int)

	for i, n := range nums {
		diff[target-n] = i
	}

	for j, m := range nums {
		if r, ok := diff[m]; ok && r != j {
			return []int{j, r}
		}
	}

	return []int{}
}