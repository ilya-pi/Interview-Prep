package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Stack5[T constraints.Ordered] []T

func (s Stack5[T]) push(v T) Stack5[T] {
	s = append(s, v)
	return s
}

func (s Stack5[T]) pop() (Stack5[T], *T) {
	if len(s) == 0 {
		return s, nil
	}
	return s[:len(s)-1], &s[len(s)-1]
}

func (s Stack5[T]) peek() *T {
	if len(s) == 0 {
		return nil
	}
	return &s[len(s)-1]
}

func sort[T constraints.Ordered](s Stack5[T]) Stack5[T] {
	if len(s) == 0 {
		return nil
	}
	/*
		Approach:
			1. Find out length of the stack

			2. Find max element in first n..1

			3. Push the max element to the bottom

			4. Push the rest on top, apart from max element (skip it in it's order)
	*/

	// 1. Find stack depth
	var n int
	var tmp Stack5[T]
	var v *T
	for s, v = s.pop(); v != nil; s, v = s.pop() {
		n++
		tmp = tmp.push(*v)
	}
	//Revert back
	s, tmp = tmp, s

	var max *T
	//2. Max element in first n..1
	for k := n; k > 0; k-- {
		// fmt.Printf("First %v elements...\n", k)
		max = s.peek()
		var v *T
		// fmt.Printf("s is %v tmp is %v\n", s, tmp)
		for l := 0; l < k; l++ {
			s, v = s.pop()
			tmp = tmp.push(*v)
			// fmt.Printf("v= %v max = %v\n", v, max)
			if *v > *max {
				max = v
			}
		}
		// Found pointer to max
		s = s.push(*max)
		// fmt.Printf("Tmp is %v\n", tmp)
		// Push the rest to s, apart from element pointing to max
		for tmp, v = tmp.pop(); v != nil; tmp, v = tmp.pop() {
			// fmt.Printf("Pushing elem %v\n", *v)
			if max == nil {
				s = s.push(*v)
			} else if *v != *max {
				s = s.push(*v)
			} else {
				max = nil
			}
		}
	}

	return s
}

func main() {
	s := Stack5[int]{}
	s = s.push(1)
	s = s.push(2)
	s = s.push(3)
	s = s.push(4)
	s = sort(s)
	s, v := s.pop()
	for ; v != nil; s, v = s.pop() {
		fmt.Printf("Popped %v\n", *v)
	}
}
