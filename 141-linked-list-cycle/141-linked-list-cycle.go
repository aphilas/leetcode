/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}

	// Tortoise and hare
	t := head
	h := head

	for h != nil && h.Next != nil {
		t = t.Next
		h = h.Next.Next
		if h == t {
			return true
		}
	}

	return false
}