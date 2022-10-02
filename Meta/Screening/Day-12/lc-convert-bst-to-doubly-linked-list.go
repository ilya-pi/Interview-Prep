/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *Node) *Node {
	/*
	   Well, since it must be sorted we need to walk the tree left-self-right

	   So we will neet to pass self to left, to be part of the sequence there, or we need to get last element of the sequence on the left, to put self as it's right element, and self.left as that element

	   But in the right we will need the first element of that sequence to attach self to it

	   So we will need to return first and last element from the inOrder walk

	   First, last of nil would be nil
	   if First and last on the left are nil -> then first in this iteration is self
	   if first and last of the right iteration are nil -> then last of this iteration is self

	   In the end we will need to glue first last of the whole tree and return first
	*/
	var inOrder func(*Node) (*Node, *Node)
	inOrder = func(t *Node) (*Node, *Node) {
		if t == nil {
			return nil, nil
		}
		firstL, lastL := inOrder(t.Left)
		firstR, lastR := inOrder(t.Right)
		if lastL != nil {
			lastL.Right = t
			t.Left = lastL
		}
		if firstR != nil {
			firstR.Left = t
			t.Right = firstR
		}

		var first, last *Node
		if firstL == nil {
			first = t
		} else {
			first = firstL
		}
		if lastR == nil {
			last = t
		} else {
			last = lastR
		}
		return first, last
	}
	f, l := inOrder(root)
	if f != nil {
		f.Left = l
		l.Right = f
	}
	return f
}
