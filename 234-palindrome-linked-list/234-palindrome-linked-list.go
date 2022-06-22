func isPalindrome(head *ListNode) bool {
	var prev *ListNode

	if head == nil || head.Next == nil {
		return true
	}

	f, s := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	for s != nil {
		temp := s.Next
		s.Next = prev
		prev = s
		s = temp
	}

	f, s = head, prev

	for s != nil {
		if s.Val != f.Val {
			return false
		}

		s = s.Next
		f = f.Next
	}

	return true
}