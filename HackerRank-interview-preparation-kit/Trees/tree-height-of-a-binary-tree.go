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
	v           int64
	left, right *Tree
}

func (t *Tree) insert(v int64) *Tree {
	n := &Tree{v: v}
	if t == nil {
		return n
	}

	if t.v > v {
		if t.left == nil {
			t.left = n
		} else {
			t.left.insert(v)
		}
	} else {
		if t.right == nil {
			t.right = n
		} else {
			t.right.insert(v)
		}
	}
	return t
}

func (t *Tree) height() int {
	/*
	   DFS from root with tracking max height
	*/
	var height int
	var dfs func(int, *Tree)
	dfs = func(h int, n *Tree) {
		if h > height {
			height = h
		}
		if n.left != nil {
			dfs(h+1, n.left)
		}
		if n.right != nil {
			dfs(h+1, n.right)
		}
	}
	dfs(0, t)
	return height
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
		val, err := strconv.ParseInt(v, 10, 64)
		checkError(err)
		n = n.insert(val)
	}

	result := n.height()

	fmt.Fprintf(writer, "%d\n", result)

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
