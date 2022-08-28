package main

import "fmt"

type Tree9 struct {
	v           int
	left, right *Tree9
}

func mergeInOrder(a1, a2 []int) [][]int {
	/*
	   We need to merge two arrays while
	   maintaining their internal order
	*/

	// One array is emppty
	if len(a1) == 0 {
		a1, a2 = a2, a1
	}
	if len(a2) == 0 {
		return [][]int{a1}
	}

	b1 := mergeInOrder(a1[1:], a2)
	b2 := mergeInOrder(a1, a2[1:])

	for i, _ := range b1 {
		b1[i] = append([]int{a1[0]}, b1[i]...)
	}

	for i, _ := range b2 {
		b2[i] = append([]int{a2[0]}, b2[i]...)
	}

	res := append(b1, b2...)
	fmt.Printf("merge %v and %v is %v\n", a1, a2, res)

	return res
}

func (t *Tree9) arrVariations() [][]int {

	/*
		The root node is always the first one to be inserted.
		Sub-trees can be inserted in different order as long
		as the order of elements within them is correct
	*/

	if t == nil {
		return nil
	}

	if t.left == nil && t.right == nil {
		fmt.Printf("Returning %v\n", t.v)
		return [][]int{{t.v}}
	}

	fmt.Printf("Node %v\n", t.v)

	lV := t.left.arrVariations()
	rV := t.right.arrVariations()

	fmt.Printf("lV == %v\n", lV)
	fmt.Printf("rV == %v\n", rV)

	var mO [][]int
	if len(lV) == 0 {
		mO = rV
	} else if len(rV) == 0 {
		mO = lV
	} else {
		for _, v1 := range lV {
			for _, v2 := range rV {
				mO = append(mO, mergeInOrder(v1, v2)...)
			}
		}
	}

	fmt.Printf("mO is %v for %v\n", mO, t.v)

	res := make([][]int, len(mO))
	for i, v := range mO {
		res[i] = append([]int{t.v}, v...)
	}
	fmt.Printf("Returning %v for %v\n", res, t.v)
	return res
}

func main() {
	t := &Tree9{v: 10}
	t1 := &Tree9{v: 5}
	t2 := &Tree9{v: 15}
	t3 := &Tree9{v: 2}
	t4 := &Tree9{v: 6}
	t5 := &Tree9{v: 12}
	t.left, t.right = t1, t2
	t1.left, t1.right = t3, t4
	t2.left = t5

	for _, v := range t.arrVariations() {
		fmt.Printf("%v\n", v)
	}
}
