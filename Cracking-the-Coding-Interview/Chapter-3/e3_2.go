package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Elem3[T constraints.Ordered] struct {
	v   T
	min T
}

type Stack3[T constraints.Ordered] []Elem3[T]

func (s Stack3[T]) push(v T) Stack3[T] {
	min := v
	if len(s) > 0 && s[len(s)-1].min < min {
		min = s[len(s)-1].min
	}
	e := Elem3[T]{v: v, min: min}
	return append(s, e)
}

func (s Stack3[T]) pop() (Stack3[T], *T) {
	if len(s) == 0 {
		return nil, nil
	}
	r := s[len(s)-1]
	s = s[:len(s)-1]
	return s, &(r.v)
}

func (s Stack3[T]) min() *T {
	if len(s) == 0 {
		return nil
	}
	return &(s[len(s)-1].min)
}

/*
push
pop
min

All in O(1)

Thought #1:

Since min is in O(1) we should know what min is at all points. We can pre-calculate
that on insert, but since we mihgt "pop" the min, we need to know what is the next
min after that. But since we might also pop the middle of the chain of mins, we'll
need to maintain what is the next min, in order to update it's record. All can be
implemented on top of a slice probably.

Thought #2:

We can just maintain the min value at each level of stack, we won't know what node
it is then, but we'll know the min value.
*/

func main() {
	s := Stack3[int]{}
	s = s.push(2)
	s = s.push(3)
	s = s.push(1)
	fmt.Printf("S is %v\n", s)
	fmt.Printf("Min is %v\n", *(s.min()))
	s, _ = s.pop()
	fmt.Printf("Min is %v\n", *(s.min()))
}
