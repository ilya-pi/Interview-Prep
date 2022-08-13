package main

import "fmt"

type LL5 struct {
	data int
	next *LL5
}

func (l *LL5) append(v int) *LL5 {
	if l == nil {
		return &LL5{data: v}
	}
	k := l
	for ; k.next != nil; k = k.next {
	}
	n := &LL5{data: v}
	k.next = n
	return l
}

func (l *LL5) String() string {
	r := ""
	for k := l; k != nil; k = k.next {
		r += fmt.Sprintf("%v -> ", k.data)
	}
	r += "nil"
	return r
}

func (l *LL5) add(l2 *LL5) *LL5 {
	/*
		Go number by number with two pointers maintaining overflow
		And then "draining" the list we didn't iterate over
	*/
	var overflow int
	i := l
	j := l2

	var r *LL5
	for i != nil && j != nil {
		v := (i.data + j.data + overflow) % 10
		overflow = (i.data + j.data + overflow) / 10
		r = r.append(v)
		i = i.next
		j = j.next
	}

	if j == nil {
		i, j = j, i
	}

	// Drain j
	for k := j; k != nil; k = k.next {
		v := (k.data + overflow) % 10
		overflow = (k.data + overflow) / 10
		r = r.append(v)
	}
	if overflow != 0 {
		r = r.append(overflow)
	}
	return r
}

func (l *LL5) reverse() *LL5 {
	head := l
	var prev *LL5
	for current := head; current != nil; {
		t := current.next
		current.next = prev
		prev = current
		head = current
		current = t
	}
	return head
}

func main() {
	l1 := &LL5{data: 9}
	l1 = l1.append(7)
	l1 = l1.append(8)
	fmt.Printf("Orig = %v\n", l1)
	l1 = l1.reverse()
	fmt.Printf("Reverse = %v\n", l1)
	l2 := &LL5{data: 6}
	l2 = l2.append(8)
	l2 = l2.append(5)
	l := l1.add(l2)
	fmt.Printf("%v\n + %v\n = %v\n", l1, l2, l)
}
