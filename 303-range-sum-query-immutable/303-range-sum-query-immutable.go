type NumArray struct {
	acc []int
}

func (a *NumArray) SumRange(l, r int) int {
	if l == 0 {
		return a.acc[r]
	}

	return a.acc[r] - a.acc[l-1]
}

func Constructor(nums []int) NumArray {
	n := NumArray{make([]int, len(nums))}
	n.acc[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		n.acc[i] = n.acc[i-1] + nums[i]
	}

	return n
}