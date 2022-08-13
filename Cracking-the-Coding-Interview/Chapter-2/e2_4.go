package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type LL4[T constraints.Ordered] struct {
	data T
	next *LL4[T]
}

func (l *LL4[T]) append(v T) *LL4[T] {
	if l == nil {
		n := &LL4[T]{data: v}
		return n
	}
	k := l
	for ; k.next != nil; k = k.next {
	}
	n := &LL4[T]{data: v}
	k.next = n
	return l
}

func (l *LL4[T]) partition(v T) *LL4[T] {
	/*
	   Maintain two lists and add to either left of right part of the partition and then glue them togeter
	*/
	var left *LL4[T]
	var right *LL4[T]

	for k := l; k != nil; k = k.next {
		if v > k.data {
			left = left.append(k.data)
		} else {
			right = right.append(k.data)
		}
	}
	k := left
	for ; k.next != nil; k = k.next {
	}
	k.next = right
	return left
}

func (l *LL4[T]) String() string {
	res := ""
	for k := l; k != nil; k = k.next {
		res += fmt.Sprintf("%v -> ", k.data)
	}
	res += "nil"
	return res
}

func main() {
	l := &LL4[int]{data: 8}
	l = l.append(7)
	l = l.append(6)
	l = l.append(5)
	l = l.append(4)
	l = l.append(3)
	l = l.append(2)
	l = l.append(1)
	fmt.Printf("%v\n", l)
	l = l.partition(5)
	fmt.Printf("%v\n", l)
}
