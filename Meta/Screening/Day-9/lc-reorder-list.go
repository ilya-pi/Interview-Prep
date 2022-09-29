/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	/*
	   rev

	   next = head.next
	   head.next = last
	   pre-to-last.next= null
	   head.next.next = rev rest
	   return head
	*/
	/*
	   1 2 3 4
	   next = 2
	   preLast = 3
	   h.Next = 4
	   3 -> nil
	   1 -> 4 -> ro(2->3)
	*/

	var reorder func(*ListNode) *ListNode
	reorder = func(h *ListNode) *ListNode {
		if h == nil || h.Next == nil || h.Next.Next == nil {
			return h
		}
		next := h.Next
		preLast := h
		for ; preLast.Next.Next != nil; preLast = preLast.Next {
		}
		h.Next = preLast.Next
		preLast.Next = nil
		h.Next.Next = reorder(next)
		return h
	}
	reorder(head)

}
