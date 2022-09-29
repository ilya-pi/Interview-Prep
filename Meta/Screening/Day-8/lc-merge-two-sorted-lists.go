/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	/*
	   Pick the bigger first node and assign it to head and cur

	   In a loop, have two pointers to the following nodes of lists, at each step weave the smaller one and move it's pointer further, until both pointers are nil

	*/
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head, cur, p1, p2 *ListNode
	if list1.Val < list2.Val {
		head, cur = list1, list1
		p1 = cur.Next
		p2 = list2
	} else {
		head, cur = list2, list2
		p1 = list1
		p2 = cur.Next
	}
	for p1 != nil || p2 != nil {
		if p1 != nil && (p2 == nil || p1.Val < p2.Val) {
			// take p1
			cur.Next = p1
			cur = p1
			p1 = p1.Next
		} else {
			// take p2
			cur.Next = p2
			cur = p2
			p2 = p2.Next
		}
	}
	//cur.Next will be nil as it will be an ending of one of the lists
	return head
}
