/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	curr := head
	m := make(map[int]struct{})

	for curr.Next != nil {
		m[curr.Val] = struct{}{}

		// If next has been seen, bypass it
		if _, ok := m[curr.Next.Val]; ok {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}