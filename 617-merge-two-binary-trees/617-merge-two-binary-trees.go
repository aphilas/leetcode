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

	a.Val += b.Val
	a.Left = mergeTrees(a.Left, b.Left)
	a.Right = mergeTrees(a.Right, b.Right)

	return a
}