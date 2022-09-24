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
 * Complete the 'bfs' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. 2D_INTEGER_ARRAY edges
 *  4. INTEGER s
 */

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// 22:00
	// 22:08 - impl
	// 22:15 - passed
	/*
	   Approach:

	   1/ create adj map
	   2/ run bfs with node + distance, recording distance in arr
	   3/ cut out the starting element
	*/
	// 1
	adj := map[int32][]int32{}
	for _, v := range edges {
		a, b := v[0], v[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// 2
	ans := make([]int32, n)
	type NodeDistance struct {
		node     int32
		distance int32
	}
	q := []NodeDistance{{s, 0}}
	visited := make([]bool, n)
	for len(q) > 0 {
		a := q[0]
		q = q[1:]

		// record distance
		ans[a.node-1] = a.distance
		// run for children
		distance := a.distance + 6
		for _, child := range adj[a.node] {
			if visited[child-1] {
				continue
			}
			visited[child-1] = true // nodes are 1-based
			q = append(q, NodeDistance{child, distance})
		}
	}

	// 3 cut out and re-map
	ans = append(ans[:s-1], ans[s:]...)
	for i, v := range ans {
		if v == 0 {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != 2 {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
