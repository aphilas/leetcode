/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    	if root == nil {
		return nil
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		for _, n := range q {
			q = q[1:]
			n.Left, n.Right = n.Right, n.Left

			if n.Right != nil {
				q = append(q, n.Right)
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
		}
	}

	return root
}