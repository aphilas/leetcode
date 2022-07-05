func construct2DArray(nums []int, m int, n int) [][]int {
    if len(nums) != m*n {
		return [][]int{}
	}

	res := make([][]int, m)

	for i := 0; i < len(nums); i += n {
		res[i/n] = nums[i : i+n]
	}

	return res
}