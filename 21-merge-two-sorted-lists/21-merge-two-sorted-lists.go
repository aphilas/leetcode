/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
    head := &ListNode{Val: 0, Next: nil}
	curr := head

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	for a != nil && b != nil {
		if a.Val < b.Val {
			curr.Next = a
			curr = a
			a = a.Next
		} else {
			curr.Next = b
			curr = b
			b = b.Next
		}
	}

	if a != nil {
		curr.Next = a
	} else {
		curr.Next = b
	}

	return head.Next
}