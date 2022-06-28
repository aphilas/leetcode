/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(a,b *TreeNode) *TreeNode {
    if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	return &TreeNode{Val: a.Val + b.Val, Left: mergeTrees(a.Left, b.Left), Right: mergeTrees(a.Right, b.Right)}
}