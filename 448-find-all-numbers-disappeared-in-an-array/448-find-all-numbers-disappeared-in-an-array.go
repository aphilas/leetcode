func findDisappearedNumbers(nums []int) []int {
    s := make([]int, 0)
	mask := make([]bool, len(nums)) // Extra 0(n) space

	for _, num := range nums {
		mask[num-1] = true
	}

	for i, v := range mask {
		if !v {
			s = append(s, i+1)
		}
	}

	return s
}