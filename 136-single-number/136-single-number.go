func singleNumber(nums []int) int {
    h := make(map[int]bool)

	for _, n := range nums {
		if _, ok := h[n]; ok {
			h[n] = false
		} else {
			h[n] = true
		}
	}

	for k, v := range h {
		if v {
			return k
		}
	}

	return 0
}