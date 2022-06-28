/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var (
		nodeDiam func(n *TreeNode) int
		diam     int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	nodeDiam = func(n *TreeNode) int {
		if n == nil {
			return 1
		}

		l := nodeDiam(n.Left)
		r := nodeDiam(n.Right)
		diam = max(diam, l+r)

		return max(l, r) + 1
	}

	nodeDiam(root)

	return diam - 2
}