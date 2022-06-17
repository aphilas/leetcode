func maxSubArray(nums []int) int {
    best := -10000
	current := 0
	neg := true
	hin := -10000 // Highest negative

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for _, n := range nums {
		if n > 0 {
			neg = false
		}

		if neg && n > hin {
			hin = n
		}

		current = max(0, current+n)
		best = max(best, current)
	}

	if neg {
		return hin
	} else {
		return best
	}
}