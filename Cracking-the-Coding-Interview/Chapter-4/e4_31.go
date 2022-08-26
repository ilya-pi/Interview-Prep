package main

import "fmt"

type Tree31 struct {
	v           int
	left, right *Tree31
}

type LL31 struct {
	v    int
	next *LL31
}

func (l *LL31) add(v int) *LL31 {
	n := &LL31{v: v}
	if l == nil {
		return n
	}
	k := l
	for ; k.next != nil; k = k.next {
	}
	k.next = n
	return l
}

func (l *LL31) String() string {
	res := ""
	for k := l; k != nil; k = k.next {
		res += fmt.Sprintf("%v -> ", k.v)
	}
	res += "nil"
	return res
}

func (t *Tree31) listOfDepths() []*LL31 {
	r := map[int]*LL31{}

	var recordDepths func(*Tree31, int)
	recordDepths = func(n *Tree31, d int) {
		r[d] = r[d].add(n.v)
		if n.left != nil {
			recordDepths(n.left, d+1)
		}
		if n.right != nil {
			recordDepths(n.right, d+1)
		}
	}
	recordDepths(t, 0)
	res := make([]*LL31, len(r))
	for k, v := range r {
		res[k] = v
	}
	return res
}

func main() {
	t := &Tree31{v: 1}
	t1 := &Tree31{v: 2}
	t2 := &Tree31{v: 3}
	t3 := &Tree31{v: 4}
	t4 := &Tree31{v: 5}
	t5 := &Tree31{v: 6}

	t.left = t1
	t.right = t2
	t1.left = t5
	t1.right = t4
	t4.right = t3
	for _, v := range t.listOfDepths() {
		fmt.Printf("%v\n", v)
	}
}
