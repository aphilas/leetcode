/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	var head *ListNode

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.Val < b.Val {
		head = a
		a = a.Next
	} else {
		head = b
		b = b.Next
	}

	curr := head

	for a != nil && b != nil {
		if a.Val < b.Val {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}

		curr = curr.Next
	}

	if a == nil {
		curr.Next = b
	} else {
		curr.Next = a
	}

	return head
}