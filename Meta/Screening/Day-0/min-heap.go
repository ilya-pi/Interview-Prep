package main

import "fmt"

type MinHeap []int

func (h MinHeap) hasLeft(ind int) bool {
	left := ind*2 + 1
	if left < len(h) {
		return true
	}
	return false
}

func (h MinHeap) hasRight(ind int) bool {
	right := ind*2 + 2
	if right < len(h) {
		return true
	}
	return false
}

func (h MinHeap) sift(ind int) {
	// check it is bigger then parent
	parent := (ind - 1) / 2
	if ind != 0 && h[parent] > h[ind] {
		// swap and check heap properties again
		h[parent], h[ind] = h[ind], h[parent]
		h.sift(parent)
		return
	}
	// check it is smaller then children
	left, right := ind*2+1, ind*2+2
	switch {
	case h.hasLeft(ind) && h[ind] > h[left] && ((h.hasRight(ind) && h[left] < h[right]) || !h.hasRight(ind)):
		// swap left
		h[ind], h[left] = h[left], h[ind]
		h.sift(left)
		return
	case h.hasRight(ind) && h[ind] > h[right]:
		// swap right
		h[ind], h[right] = h[right], h[ind]
		h.sift(right)
		return
	}
}

func (h MinHeap) Add(v int) MinHeap {
	h = append(h, v)
	h.sift(len(h) - 1)
	return h
}

func (h MinHeap) Pop() (*int, MinHeap) {
	if len(h) == 0 {
		return nil, h
	}
	ans := h[0]
	h[0], h[len(h)-1], h = h[len(h)-1], h[0], h[:len(h)-1]
	h.sift(0)
	return &ans, h
}

func (h MinHeap) Peek() *int {
	if len(h) == 0 {
		return nil
	}
	return &h[0]
}

func main() {
	minHeap := MinHeap{}
	minHeap = minHeap.Add(3)
	minHeap = minHeap.Add(10)
	minHeap = minHeap.Add(2)
	minHeap = minHeap.Add(8)
	minHeap = minHeap.Add(32)
	fmt.Printf("%v\n\n", minHeap)
	//fmt.Printf("min - %v\n", *(minHeap.Peek()))
	for v, minHeap := minHeap.Pop(); v != nil; v, minHeap = minHeap.Pop() {
		if v != nil {
			fmt.Printf("Popped - %v\n", *v)
			//fmt.Printf("%v\n", minHeap)
			//fmt.Printf("Min - %v\n", *(minHeap.Peek()))
		}
	}
}
