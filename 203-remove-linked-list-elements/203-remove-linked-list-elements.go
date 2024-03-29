/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var prev *ListNode
	curr := head

	for curr != nil {
		if curr.Val == val {
			if prev == nil {
				head = curr.Next
				curr = head
				continue
			} else {
				if curr.Next != nil {
					prev.Next = curr.Next
				} else {
					prev.Next = nil
				}
			}
		} else {
			// Only move prev is current is valid
			prev = curr
		}

		curr = curr.Next
	}

	return head
}