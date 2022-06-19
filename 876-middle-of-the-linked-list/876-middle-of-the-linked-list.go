/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    t, h := head, head

	for h != nil {
		// End of the road
		if h.Next == nil {
			return t
		}

		if h.Next.Next == nil {
			return t.Next
		}

		t = t.Next
		h = h.Next.Next
	}

	return head
}