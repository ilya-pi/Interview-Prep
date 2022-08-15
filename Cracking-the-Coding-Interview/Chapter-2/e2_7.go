package main

import "fmt"

type LL7[T comparable] struct {
	data T
	next *LL7[T]
}

func (l *LL7[T]) append(v T) *LL7[T] {
	n := &LL7[T]{data: v}
	if l == nil {
		return n
	}
	k := l
	for ; k.next != nil; k = k.next {
	}
	k.next = n
	return l
}

func (l *LL7[T]) String() string {
	res := ""
	for k := l; k != nil; k = k.next {
		res += fmt.Sprintf("%v -> ", k.data)
	}
	res += "nil"
	return res
}

func (l *LL7[T]) intersection(l2 *LL7[T]) *LL7[T] {
	/*

	   1 -> 2 -> 3 -> 4 -> 5
	   0 -> 0 -> 0 -> 0 -> 3 -> 4 -> 5

	   The way I understand it, two lists intersect means they

	*/

	/*
	   Let's implement bruteforce and then see if we can optimize
	   O(n*m)
	*/
	/*
		for i := l; i != nil; i = i.next {
			for j := l2; j != nil; j = j.next {
				if i == j {
					return i
				}
			}
		}
		return nil
	*/

	/*
		Optimize with map the second operation
		O(m + n) and O(m) space
	*/
	/*
		m := map[*LL7[T]]bool{}
		for j := l2; j != nil; j = j.next {
			m[j] = true
		}
		for i := l; i != nil; i = i.next {
			if _, ok := m[i]; ok {
				return i
			}
		}
		return nil
	*/

	/* What we don't use is that tail is the same, but I don't see how we can smartly utilise it */
	/*
		Let's try using it with checking
		1) Do the lists intersect?
		2) What is the length difference?
		3) Iterating one by one with length difference adjustment
	*/

	var len1 int
	var len2 int
	i := l
	j := l2
	// Get last nodes and compare
	for ; i.next != nil; i = i.next {
		len1++
	}
	for ; j.next != nil; j = j.next {
		len2++
	}
	// Lists don't intersect
	if i != j {
		return nil
	}
	// Lists intersect
	if len1 < len2 {
		l, i, len1, l2, j, len2 = l2, j, len2, l, i, len1
	}
	i, j = l, l2
	// Skip the first elements in the shorter list
	for k := 0; k < len1-len2; k++ {
		fmt.Printf("Skipping %v\n", i.data)
		i = i.next
	}
	for ; i != nil && j != nil; i, j = i.next, j.next {
		if i == j {
			return i
		}
	}
	// NB: Technically unreachable
	return nil
}

func main() {
	l1 := &LL7[int]{data: 1}
	l1 = l1.append(2)
	l1 = l1.append(3)
	l1 = l1.append(4)
	l1 = l1.append(5)
	l2 := &LL7[int]{data: 0}
	l2 = l2.append(0)
	l2 = l2.append(0)
	l2.next.next.next = l1.next.next
	l3 := &LL7[int]{data: 1}
	l3 = l3.append(3)
	l3 = l3.append(4)
	fmt.Printf("l1 = %v\n", l1)
	fmt.Printf("l2 = %v\n", l2)
	fmt.Printf("l3 = %v\n", l3)
	fmt.Printf("Intersection l1 l2 = %v\n", l1.intersection(l2))
	fmt.Printf("Intersection l1 l3 = %v\n", l1.intersection(l3))

}
