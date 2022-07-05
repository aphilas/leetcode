func majorityElement(nums []int) int {
    var cand, count int

	for _, num := range nums {
		if count == 0 {
			cand = num
		}

		if num == cand {
			count += 1
		} else {
			count -= 1
		}
	}

	return cand
}