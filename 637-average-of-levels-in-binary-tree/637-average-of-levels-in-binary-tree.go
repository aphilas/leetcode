/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    	a := []float64{}

	if root == nil {
		return a
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		l := []float64{}

		for _, n := range q {
			l = append(l, float64(n.Val))

			q = q[1:] // Remove

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		sum := 0.0

		for _, val := range l {
			sum += val
		}

		a = append(a, sum/float64(len(l)))
	}

	return a
}