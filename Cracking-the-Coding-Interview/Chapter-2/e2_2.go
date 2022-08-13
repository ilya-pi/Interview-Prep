package main

import "fmt"

type LL2[T comparable] struct {
	data T
	next *LL2[T]
}

func (l *LL2[T]) append(v T) {
	i := l
	for ; i.next != nil; i = i.next {
	}
	n := &LL2[T]{data: v}
	i.next = n
}

func (l *LL2[T]) String() string {
	var r string
	for i := l; i != nil; i = i.next {
		r += fmt.Sprintf("%v -> ", i.data)
	}
	r += "nil"
	return r
}

func (l *LL2[T]) removeKthFromEnd(k int) {
	if k < 1 {
		//o-th element from end is nil
		return
	}
	// Make k steps forward and then iterate with both pointers to the end
	kOffset := l
	end := l
	for i := 0; i < k; i++ {
		if end == nil {
			// List is not long enough, such element doesn't exist
			return
		}
		end = end.next
	}
	for end.next != nil {
		kOffset = kOffset.next
		end = end.next
	}
	kOffset.next = kOffset.next.next
}

func main() {
	l := &LL2[int]{data: 1}
	l.append(2)
	l.append(3)
	l.append(4)
	l.append(5)
	fmt.Printf("%v\n", l)
	l.removeKthFromEnd(3)
	fmt.Printf("%v\n", l)
}
