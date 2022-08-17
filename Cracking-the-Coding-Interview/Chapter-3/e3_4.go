package main

import "fmt"

type Stack6[T any] []T

func (s Stack6[T]) push(v T) Stack6[T] {
	s = append(s, v)
	return s
}

func (s Stack6[T]) pop() (*T, Stack6[T]) {
	if len(s) == 0 {
		return nil, s
	}
	return &s[len(s)-1], s[:len(s)-1]
}

func (s Stack6[T]) peek() *T {
	if len(s) == 0 {
		return nil
	}
	return &s[len(s)-1]
}

type Queue[T any] struct {
	left, right Stack6[T]
}

func (q *Queue[T]) push(v T) {
	if len(q.right) != 0 {
		q.toLeft()
	}
	q.left = q.left.push(v)
}

func (q *Queue[T]) pop() *T {
	if len(q.right) == 0 {
		q.toRight()
	}
	var r *T
	r, q.right = q.right.pop()
	return r
}

func (q *Queue[T]) peek() *T {
	if len(q.right) == 0 {
		q.toRight()
	}
	return q.right.peek()
}

func (q *Queue[T]) toRight() {
	var v *T
	for q.left.peek() != nil {
		v, q.left = q.left.pop()
		q.right = q.right.push(*v)
	}
}

func (q *Queue[T]) toLeft() {
	var v *T
	for q.right.peek() != nil {
		v, q.right = q.right.pop()
		q.left = q.left.push(*v)
	}
}

func main() {
	q := &Queue[int]{}
	q.push(1)
	q.push(2)
	q.push(3)
	v := q.pop()
	fmt.Printf("Pop %v\n", *v)
	q.push(4)
	v = q.pop()
	fmt.Printf("Pop %v\n", *v)

}
