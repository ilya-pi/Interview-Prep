package main

import "fmt"

type Tree3 struct {
	v     int
	left  *Tree3
	right *Tree3
}

type LL3 struct {
	v    int
	next *LL3
}

func (l *LL3) String() string {
	res := ""
	for k := l; k != nil; k = k.next {
		res += fmt.Sprintf("%v -> ", k.v)
	}
	res += "nil"
	return res
}

func listOfDepths(t *Tree3) []*LL3 {
	/*
		Essentially it is a DFS with adding a list to the results once we hit nil node
	*/
	r := map[*LL3]bool{}
	dfs3(t, nil, &r)
	res := []*LL3{}
	for k, _ := range r {
		res = append(res, k)
	}
	return res
}

func dfs3(t *Tree3, l *LL3, r *map[*LL3]bool) {
	// Visit
	if t == nil {
		// Save the accumulated list in a map (set)
		(*r)[l] = true
		return
	}

	// Process left and right
	n := &LL3{v: t.v, next: l}
	dfs3(t.left, n, r)
	dfs3(t.right, n, r)
}

func main() {
	t := &Tree3{v: 1}
	t1 := &Tree3{v: 2}
	t2 := &Tree3{v: 3}
	t3 := &Tree3{v: 4}
	t4 := &Tree3{v: 5}
	t5 := &Tree3{v: 6}
	t6 := &Tree3{v: 7}
	t.left = t1
	t.right = t2
	t1.left = t3
	t1.right = t4
	t3.left = t5
	t3.right = t6
	for _, v := range listOfDepths(t) {
		fmt.Printf("%v\n", v)
	}
}
