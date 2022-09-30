/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	/*
	   1. Weave new list into old list as second nodes
	   2. Copy all Random links
	   3. Remove all original list nodes
	*/

	if head == nil {
		return nil
	}

	// Weave new list into old list
	for cur := head; ; {
		next := cur.Next
		n := &Node{Val: cur.Val, Next: next}
		cur.Next = n
		cur = next
		if cur == nil {
			break
		}
	}

	// Copy all random links
	for cur := head; ; {
		next := cur.Next.Next
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = next
		if cur == nil {
			break
		}
	}

	ans := head.Next
	// Remove original list nodes
	for cur := head; ; {
		if cur == nil {
			break
		}
		// unweave lists
		oNext := cur.Next.Next
		var nNext *Node
		if cur.Next.Next != nil {
			nNext = cur.Next.Next.Next // 0 - 01 - 2 - 21 - 3 - 31
		}
		// unweave copy
		cur.Next.Next = nNext
		// unweave original
		cur.Next = oNext
		cur = oNext

	}
	return ans

}
