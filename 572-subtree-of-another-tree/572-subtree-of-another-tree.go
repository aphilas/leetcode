func Equal(a, b *TreeNode) bool {
    if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return Equal(a.Left, b.Left) && Equal(a.Right, b.Right)
}


func isSubtree(r, s *TreeNode) bool {
    if Equal(r, s) {
		return true
	}
    
    if r == nil {
        return false
    }

	return isSubtree(r.Left, s) || isSubtree(r.Right, s)
}
