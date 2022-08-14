package main

import "fmt"

type LL6[T comparable] struct {
	data T
	next *LL6[T]
}

func (l *LL6[T]) append(v T) *LL6[T] {
	n := &LL6[T]{data: v}
	if l == nil {
		return n
	}
	k := l
	for ; k.next != nil; k = k.next {
	}
	k.next = n
	return l
}

func (l *LL6[T]) String() string {
	res := ""
	for k := l; k != nil; k = k.next {
		res += fmt.Sprintf("%v -> ", k.data)
	}
	res += "nil"
	return res
}

func (l *LL6[T]) isPalindrome() bool {
	/*

	   1 -> 2 -> 3 -> 3 -> 2 -> 1

	   We will find the middle element and while doing it will reverse the original list with the pointer running slower and then check the new lists are equal


	   1 2 3 4 5

	   1 2 3 <- middle
	   1 3 5


	   1 2 3 4

	   1 2 3 <- prev
	   1 3 nil


	   1 2 3

	   1.next.next
	*/

	var middle int
	i, j := l, l
	for ; j != nil && j.next != nil; i, j = i.next, j.next.next {
		middle++
	}
	// reverse for the correct amount of elements
	var prev *LL6[T]
	i = l
	for k := 0; k < middle; k++ {
		next := i.next
		i.next = prev
		prev = i
		i = next
	}
	// Check if there was a middle element
	if j != nil {
		prev = prev.next
	}
	return i.equals(prev)
}

func (l *LL6[T]) equals(l2 *LL6[T]) bool {
	i, j := l, l2
	for ; i != nil && j != nil; i, j = i.next, j.next {
		if i.data != j.data {
			return false
		}
	}
	return i == j
}

func main() {
	l := &LL6[int]{data: 1}
	l = l.append(2)
	l = l.append(3)
	l = l.append(3)
	l = l.append(2)
	l = l.append(1)
	// Since we implemented an inplace reverse, that corrupts the original list supplied to the function,
	// hence we remember the serialized output of it beforehand
	str := fmt.Sprintf("%v", l)
	fmt.Printf("%v is palindrome == %v\n", str, l.isPalindrome())
}
