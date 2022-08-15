package main

import "fmt"

type Elem[T comparable] struct {
	v    T
	prev int
}

type Stack1[T comparable] struct {
	storage [1000]*Elem[T]
	heads   [3]int
}

func newStack[T comparable]() *Stack1[T] {
	s := &Stack1[T]{
		storage: [1000]*Elem[T]{},
		heads:   [3]int{},
	}
	// -1 means we reached the bottom of the stack
	for i := 0; i < len(s.heads); i++ {
		s.heads[i] = -1
	}
	return s
}

func (s *Stack1[T]) peek(i int) *T {
	if i >= len(s.heads) {
		return nil
	}
	if s.heads[i] < 0 {
		return nil
	}
	r := s.storage[s.heads[i]].v
	return &r
}

func (s *Stack1[T]) pop(i int) *T {
	if i >= len(s.heads) {
		return nil
	}
	if s.heads[i] < 0 {
		return nil
	}
	r := s.storage[s.heads[i]]
	s.storage[s.heads[i]] = nil
	s.heads[i] = r.prev
	res := r.v
	return &res
}

func (s *Stack1[T]) push(i int, v T) {
	if i > len(s.heads) {
		return
	}
	head := s.heads[i]
	e := &Elem[T]{
		v:    v,
		prev: head,
	}
	// Find a place to insert this Elem
	spot := -1
	for i := 0; i < len(s.storage); i++ {
		if s.storage[i] == nil {
			spot = i
			break
		}
	}
	// Save element in Storage
	s.storage[spot] = e
	s.heads[i] = spot
}

func (s *Stack1[T]) isEmpty(i int) bool {
	if i >= len(s.heads) {
		// Generally speaking this is a good place for our specific errors
		return true
	}
	return s.heads[i] == -1
}

//func (s Stack1[T]) push(

/*
	We have an array that will hold all the elements:
	[ .... ]

	Within it we will have different stacks (3)

	If the array is fixed length? Then we might have to "deny" element to be put in, but
	disperse elements at the worst availability.
	Then we might have to mix the elements, hence we'd want to remember where was the previous element and where is the current head of the dedicated stack.
	In an event that we did push and there is no place, we'll scan the array to see if there is more empty place, otherwise deny. That makes:
	peek - O(1)
	pop - O(1)
	push - O(n)
	isEmpty - O(1)

*/

func main() {
	s := newStack[int]()
	s.push(0, 1)
	s.push(0, 2)
	s.push(0, 4)
	s.push(2, 5)
	fmt.Printf("pop(0) = %v, peek(2) = %v, isEmpty(1) == %v\n", *(s.pop(0)), *(s.peek(2)), s.isEmpty(1))
	fmt.Printf("storage == %v\n", s.storage)
}
