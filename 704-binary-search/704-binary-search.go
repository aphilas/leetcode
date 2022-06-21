func search(nums []int, target int) int {    
    for lo, hi := 0, len(nums)-1; lo <= hi; {
		if nums[(hi + lo) / 2] == target {
			return (hi + lo) / 2
		} else if nums[(hi + lo) / 2] < target {
			lo = (hi + lo) / 2 + 1
		} else {
			hi = (hi + lo) / 2 - 1
		}
	}

	return -1
}