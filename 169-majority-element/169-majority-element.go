func majorityElement(nums []int) int {
    l := len(nums) / 2
	m := make(map[int]int)

	for _, n := range nums {
		m[n] += 1
	}

	for k, v := range m {
		if v > l {
			return k
		}
	}

	return 0
}