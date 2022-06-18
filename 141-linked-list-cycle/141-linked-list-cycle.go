/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    
    hm := make(map[*ListNode]int)
	node := head

	for node.Next != nil {
		hm[node] += 1

		if hm[node] == 2 {
			return true
		}

		node = node.Next
	}

	return false
}