/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, target int) bool {
    if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return true
	}

	return hasPathSum(root.Left, target-root.Val) || hasPathSum(root.Right, target-root.Val)
}