func maxSubArray(nums []int) int {
    global, local, largest := -10000, 0, -10000
    
    max := func (a, b int) int {
        if a > b {
            return a
        }
        
        return b
    }

	for _, n := range nums {
		local = max(local+n, 0)
		global = max(local, global)
		largest = max(largest, n)
	}

	if global == 0 {
		global = largest
	}

	return global
}