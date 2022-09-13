package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'swapNodes' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY indexes
 *  2. INTEGER_ARRAY queries
 */

type Tree struct {
	v           int32
	left, right *Tree
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	// Write your code here
	// 1 Convert indexes to a regular tree
	nodes := map[int32]*Tree{}
	for i, _ := range indexes {
		nV := int32(i + 1)
		n := Tree{v: nV}
		nodes[nV] = &n
	}
	for i, v := range indexes {
		nV := int32(i + 1)
		n := nodes[nV]
		left := v[0]
		right := v[1]

		if left != -1 {
			lN := nodes[left]
			n.left = lN
		}
		if right != -1 {
			rN := nodes[right]
			n.right = rN
		}
	}

	swapKs := func(n *Tree, k int32) *Tree {
		/*
		   Approach: bfs top to down with swap
		*/
		type NodeAndDepth struct {
			n     *Tree
			depth int32
		}
		q := []NodeAndDepth{{n, 1}}
		for len(q) > 0 {
			cn := q[0]
			q = q[1:]
			if cn.n == nil {
				continue
			}
			if cn.depth%k == 0 {
				cn.n.left, cn.n.right = cn.n.right, cn.n.left
			}
			q = append(q, NodeAndDepth{cn.n.left, cn.depth + 1})
			q = append(q, NodeAndDepth{cn.n.right, cn.depth + 1})
		}
		return n
	}

	inOrderArr := func(n *Tree) []int32 {
		var r []int32
		visit := func(n Tree) {
			r = append(r, n.v)
		}
		var inOrder func(*Tree)
		inOrder = func(n *Tree) {
			if n == nil {
				return
			}
			inOrder(n.left)
			visit(*n)
			inOrder(n.right)
		}
		inOrder(n)
		return r
	}

	n := nodes[1]
	res := make([][]int32, len(queries))
	for i, k := range queries {
		n = swapKs(n, k)
		res[i] = inOrderArr(n)
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for i := 0; i < int(n); i++ {
		indexesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != 2 {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for i, rowItem := range result {
		for j, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if j != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
