package main

import "fmt"

type LL3[T comparable] struct {
	data T
	next *LL3[T]
}

func (l *LL3[T]) append(v T) {
	k := l
	for ; k.next != nil; k = k.next {
	}
	n := &LL3[T]{data: v}
	k.next = n
}

func (l *LL3[T]) String() string {
	r := ""
	for k := l; k != nil; k = k.next {
		r += fmt.Sprintf("%v -> ", k.data)
	}
	r += "nil"
	return r
}

func (l *LL3[T]) deleteMiddleNode() {
	// Not a middle node
	if l.next == nil {
		return
	}

	// To delete middle node without access to the previous node we will need to shift data value of the next node and then fix links
	l.data = l.next.data
	l.next = l.next.next
}

func main() {
	l := &LL3[int]{data: 1}
	l.append(2)
	l.append(3)
	l.append(4)
	fmt.Printf("%v\n", l)
	l.next.next.deleteMiddleNode()
	fmt.Printf("%v\n", l)
}
