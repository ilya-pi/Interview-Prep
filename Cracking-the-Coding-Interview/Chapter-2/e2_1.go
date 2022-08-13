package main

import "fmt"

type LinkedList[T comparable] struct {
	data T
	next *LinkedList[T]
}

func (l *LinkedList[T]) append(v T) {
	n := LinkedList[T]{data: v}
	tail := l
	for ; tail.next != nil; tail = tail.next {
	}
	tail.next = &n
}

func (l *LinkedList[T]) remove(v T) *LinkedList[T] {
	if l == nil {
		return nil
	}

	if l.data == v {
		return l.next
	}

	n := l
	for ; n.next != nil; n = n.next {
		if n.next.data == v {
			n.next = n.next.next
			return l
		}
	}
	return l
}

func (l *LinkedList[T]) String() string {
	var res string
	for k := l; k != nil; k = k.next {
		res += fmt.Sprintf("%v -> ", k.data)
	}
	res += "nil"
	return res
}

func (l *LinkedList[T]) removeDups() {
	if l == nil {
		return
	}

	dups := map[T]int{}
	dups[l.data]++
	for n := l; n != nil; n = n.next {
		dups[n.next.data]++
		if dups[n.next.data] > 1 {
			n.next = n.next.next
		}
	}
}

func (l *LinkedList[T]) removeDups2() {
	if l == nil {
		return
	}

	for n := l; n != nil; n = n.next {
		for k := n.next; k != nil; k = k.next {
			// fmt.Printf("n = %v k = %v\n", n.data, k.data)
			if k.next != nil && k.next.data == n.data {
				//fmt.Printf("Removed %v\n", k.next.data)
				k.next = k.next.next
			}
		}
	}
}

func main() {
	l := &LinkedList[int]{data: 1}
	l.append(2)
	l.append(3)
	l.append(2)
	l.append(5)
	l.append(3)

	fmt.Printf("List is %v\n", l)
	//l.removeDups()
	l.removeDups2()
	fmt.Printf("List is %v\n", l)

	// If extra buffer is not allowed — I would run the inner loop for every value met, that would make it O(n(n-1 + n-2 + n-3 + ...)) -> resulting to O(n^2)
}
