/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
	   Maintain overflow and two pointers, at each step try to move pointers and add overflow, until both are at nil

	   At the end -> drain overflow if present
	*/
	var overflow int
	ans := &ListNode{}
	cur := ans
	var prev *ListNode
	for l1 != nil || l2 != nil {
		if l1 != nil {
			overflow += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			overflow += l2.Val
			l2 = l2.Next
		}
		// record value
		cur.Val = overflow % 10
		overflow /= 10
		cur.Next = &ListNode{}
		prev = cur
		cur = cur.Next
	}
	if overflow > 0 {
		cur.Val = overflow
	} else {
		prev.Next = nil
	}
	return ans
}
