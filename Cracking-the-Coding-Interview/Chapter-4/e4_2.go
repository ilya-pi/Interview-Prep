package main

import "fmt"

type Tree2 struct {
	v     int
	left  *Tree2
	right *Tree2
}

func (t *Tree2) Prn(ident int) string {
	res := ""
	if t == nil {
		for i := 0; i < ident; i++ {
			res += " "
		}
		res += "nil"
		return res
	}
	res += t.left.Prn(ident + 2)
	res += t.right.Prn(ident + 2)
	for i := 0; i < ident; i++ {
		res += " "
	}
	res += fmt.Sprintf("%v\n", t.v)
	return res
}

func (t *Tree2) String() string {
	return t.Prn(0)
}

func toMBST(arr []int) *Tree2 {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &Tree2{v: arr[0], left: nil, right: nil}
	}
	mid := arr[len(arr)/2]
	fmt.Printf("arr is %v\n", arr)
	fmt.Printf("mid is %v\n", mid)
	r := Tree2{
		v:     mid,
		left:  toMBST(arr[:len(arr)/2]),
		right: toMBST(arr[len(arr)/2+1:]),
	}

	return &r
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("%v\n", toMBST(arr))
}
