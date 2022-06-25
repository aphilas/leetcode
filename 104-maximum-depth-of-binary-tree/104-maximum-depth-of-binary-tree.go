func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode, rest ...int) int {
    var c int

	if len(rest) == 1 {
		c = rest[0]
	}

	if root == nil {
		return c
	}

	if root.Left == nil && root.Right == nil {
		return c + 1
	}

	return max(maxDepth(root.Left, c+1), maxDepth(root.Right, c+1))
}