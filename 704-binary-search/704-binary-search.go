func search(nums []int, target int) int {    
    for lo, hi := 0, len(nums)-1; lo <= hi; {
		mid := (hi + lo) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// Search right
			lo = mid + 1
		} else {
			// Search left
			hi = mid - 1
		}
	}

	return -1
}