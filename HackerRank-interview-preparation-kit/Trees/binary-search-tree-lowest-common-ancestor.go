package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	v                   int32
	parent, left, right *Tree
}

func (t *Tree) insert(v int32) *Tree {
	n := &Tree{v: v}
	if t == nil {
		return n
	}
	if v < t.v {
		if t.left == nil {
			t.left = n
			n.parent = t
			return t
		}
		t.left.insert(v)
		return t
	}
	if v >= t.v {
		if t.right == nil {
			t.right = n
			n.parent = t
			return t
		}
		t.right.insert(v)
		return t
	}
	return nil
}

func (t *Tree) find(v int32) *Tree {
	if t == nil {
		return nil
	}
	if t.v == v {
		return t
	}
	if v < t.v {
		return t.left.find(v)
	} else {
		return t.right.find(v)
	}
}

func (t *Tree) lowestCommonAncestor(n1, n2 *Tree) *Tree {
	/*
	   Approach:
	   Measure height of both nodes, on the deeper one go up to the same level and then iterate with two pointers up

	   1. Find height for both
	   2. Get to the same height with both pointers
	   3. Go up while comparing references
	*/
	getHeight := func(n *Tree) int {
		var r int
		for k := n; k != nil; k = k.parent {
			r++
		}
		return r
	}
	h1, h2 := getHeight(n1), getHeight(n2)
	if h1 > h2 {
		n1, h1, n2, h2 = n2, h2, n1, h1
	}
	// n2 is deeper then n1
	for ; h2 > h1; n2 = n2.parent {
		h2--
	}
	// both nodes are at the same height now
	for ; n1 != n2; n1, n2 = n1.parent, n2.parent {
	}

	return n1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_ = readLine(reader)
	values := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var n *Tree
	for _, v := range values {
		val, err := strconv.ParseInt(v, 10, 32)
		checkError(err)
		n = n.insert(int32(val))
	}

	nodes := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	n1V, err := strconv.ParseInt(nodes[0], 10, 32)
	checkError(err)
	n2V, err := strconv.ParseInt(nodes[1], 10, 32)
	checkError(err)
	n1, n2 := n.find(int32(n1V)), n.find(int32(n2V))

	result := n.lowestCommonAncestor(n1, n2)

	fmt.Fprintf(writer, "%d\n", result.v)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
