/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(nodes ...(*ListNode)) *ListNode {
    var prev, head *ListNode
	
    head = nodes[0]

    if len(nodes) == 2 {
		prev = nodes[1]
	}

	if head == nil {
		return prev
	}

	tmp := head.Next
	head.Next = prev

	return reverseList(tmp, head)
}
