type NumArray struct {
	nums []int
}

func (n *NumArray) SumRange(l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += n.nums[i]
	}
	return sum
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}
