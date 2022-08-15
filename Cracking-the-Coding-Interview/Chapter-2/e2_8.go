package main

import "fmt"

type LL8[T comparable] struct {
	data T
	next *LL8[T]
}

func (l *LL8[T]) append(v T) *LL8[T] {
	n := &LL8[T]{data: v}
	if l == nil {
		return n
	}
	k := l
	for ; k.next != nil; k = k.next {
	}
	k.next = n
	return l
}

func (l *LL8[T]) String() string {
	// simplified stringer that prints only current node
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%v -> ...", l.data)
}

func (l *LL8[T]) loop() *LL8[T] {
	/*
	   1. We want to understand if there is a loop, so we iterate with two pointers with 1 and 2 speed and see if they meet
	   -> they do, there is a loop
	   2. We use the point where they meet as a beginning of a another list and find intersection point between them
	*/
	i, j := l, l.next
	for ; i != nil && j != nil && j.next != nil; i, j = i.next, j.next.next {
		if i == j {
			break
		}
	}
	// There is no loop
	if j == nil || j.next == nil {
		return nil
	}
	// So there is a loop
	l1, l2 := l, i.next // We assign it to "end" so that we could understand where it ends by comparisson to end
	end := i

	// Find length
	len1, len2 := 0, 0
	for i := l1; i != end; i = i.next {
		len1++
	}
	for j := l2; j != end; j = j.next {
		len2++
	}
	if len1 < len2 {
		l1, len1, l2, len2 = l2, len2, l1, len1
	}
	// Equalize lists in length
	i, j = l1, l2
	for k := 0; k < len1-len2; k++ {
		i = i.next
	}
	for ; ; i, j = i.next, j.next {
		if i == j {
			return i
		}
	}
	// Technically not possible
	return nil
}

func main() {
	l := &LL8[int]{data: 1}
	l = l.append(2)
	l = l.append(3)
	l = l.append(4)
	l.next.next.next.next = l.next.next
	fmt.Printf("Loop is at %v\n", l.loop())
}
