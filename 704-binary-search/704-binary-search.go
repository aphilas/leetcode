func search(nums []int, rest ...int) int {
    if len(nums) == 0 {
		return -1
	}
    
    offset := 0
	target := rest[0]

	if len(rest) == 2 {
		offset = rest[1]
	}

	i := len(nums) / 2
	n := nums[i]

	if n == target {
		return i + offset
	} else if n < target {
		return search(nums[i+1:], target, offset+i+1)
	} else {
		return search(nums[:i], target, offset)
	}
}