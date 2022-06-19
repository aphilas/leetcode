/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    t, h := head, head

	for h != nil && h.Next != nil {
		// Returns second node if there are two middle nodes
		t = t.Next
		h = h.Next.Next
	}

	return t
}