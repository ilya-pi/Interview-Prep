package main

import "fmt"

type Dog string
type Cat string

/*
Keep a linked list and remove the correctly type
element from the head depending on the type.
Keep end link to have O(1) on queue and dequeue.
*/

/*
enqueue
dequeueAny
dequeueDog
dequeueCat
*/

type Animal struct {
	typ  int
	name string
}

type Node6 struct {
	v    Animal
	next *Node6
}

type Q6 struct {
	head *Node6
	tail *Node6
}

func (q *Q6) enqueue(v Animal) {
	n := &Node6{v: v}
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.next = n
	q.tail = n
}

func (q *Q6) dequeueAny() *Animal {
	if q == nil && q.head == nil {
		return nil
	}
	res := q.head.v
	q.head = q.head.next
	return &res
}

func (q *Q6) dequeue(typ int) *Animal {
	if q == nil || q.head == nil {
		return nil
	}
	var prev *Node6
	k := q.head
	for ; k != nil; prev, k = k, k.next {
		if k.v.typ == typ {
			//fmt.Printf("Found %v\n", k.v.name)
			break
		}
	}
	// No dog was found in queue
	if k == nil {
		return nil
	}
	// Remove node K
	res := k.v

	if prev != nil {
		prev.next = k.next
	} else {
		q.head = k.next
	}
	return &res
}

func main() {
	q := &Q6{}

	q.enqueue(Animal{1, "Bob"})
	q.enqueue(Animal{2, "Alice"})
	q.enqueue(Animal{1, "Nina"})
	q.enqueue(Animal{1, "Arla"})
	a := q.dequeue(1)
	fmt.Printf("Got %v\n", a)
	a = q.dequeueAny()
	fmt.Printf("Got %v\n", a)
	a = q.dequeue(1)
	fmt.Printf("Got %v\n", a)
}
