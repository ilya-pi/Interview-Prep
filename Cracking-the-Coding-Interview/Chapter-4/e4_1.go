package main

import "fmt"

type Node1[T comparable] struct {
	v     T
	nodes []*Node1[T]
}

func (n *Node1[T]) add(n2 *Node1[T]) {
	if n2 != nil {
		n.nodes = append(n.nodes, n2)
	}
}

type Queue1[T comparable] []*Node1[T]

func (q *Queue1[T]) enqueue(n *Node1[T]) {
	(*q) = append((*q), n)
}

func (q *Queue1[T]) dequeue() *Node1[T] {
	if len(*q) == 0 {
		return nil
	}
	r := (*q)[0]
	(*q) = (*q)[1:]
	return r
}

func (n *Node1[T]) havePath(n2 *Node1[T]) bool {
	/*
		Normally we should be able to figure the path out with BFS

		So we visit the node and then process all the children in FIFO, continuously expanding
	*/
	q := &Queue1[T]{}
	visited := map[*Node1[T]]bool{}
	for k := n; k != nil; k = q.dequeue() {
		//fmt.Printf("Queue = %v\n", q)
		// Check condition
		if k == n2 {
			return true
		}
		//fmt.Printf("Checked %v\n", k.v)
		// Process children
		for _, v := range k.nodes {
			if _, ok := visited[v]; !ok {
				//fmt.Printf("Enqueued %v\n", v.v)
				q.enqueue(v)
			} else {
				//fmt.Printf("Seen %v\n", v.v)
			}
		}
		visited[k] = true
	}

	return false
}

func main() {
	n := &Node1[int]{v: 1}
	n1 := &Node1[int]{v: 2}
	n2 := &Node1[int]{v: 3}
	n3 := &Node1[int]{v: 4}
	n4 := &Node1[int]{v: 5}
	n.add(n1)
	n1.add(n2)
	n.add(n3)
	fmt.Printf("Have path == %v\n", n.havePath(n2))
	fmt.Printf("Have path == %v\n", n.havePath(n4))
}
