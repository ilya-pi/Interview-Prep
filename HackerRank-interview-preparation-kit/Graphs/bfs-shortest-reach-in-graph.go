package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func distances(edges []map[int64]bool, n int64, vertices int64) []int {
	type VD struct {
		e int64
		d int
	}
	visited := make([]bool, vertices)
	visited[n] = true
	q := []VD{{n, 0}}
	distances := make([]int, vertices)
	for i := int64(0); i < vertices; i++ {
		distances[i] = -1
	}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		// Save node to distances
		distances[node.e] = node.d
		// Add all children
		for adj, _ := range edges[node.e] {
			if visited[adj] {
				continue
			}
			visited[adj] = true
			q = append(q, VD{adj, node.d + 6})
		}
	}
	return append(distances[:n], distances[n+1:]...)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	querries, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	for i := querries; i > 0; i-- {
		base := strings.Split(readLine(reader), " ")
		vertices, err := strconv.ParseInt(base[0], 10, 64)
		checkError(err)
		edgesCount, err := strconv.ParseInt(base[1], 10, 64)
		checkError(err)
		edges := make([]map[int64]bool, vertices)
		for k := int64(0); k < vertices; k++ {
			edges[k] = map[int64]bool{}
		}
		for k := edgesCount; k > 0; k-- {
			ab := strings.Split(readLine(reader), " ")
			aV, err := strconv.ParseInt(ab[0], 10, 64)
			checkError(err)
			bV, err := strconv.ParseInt(ab[1], 10, 64)
			checkError(err)
			edges[aV-1][bV-1] = true
			edges[bV-1][aV-1] = true
		}
		start, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		// Do the algo and output
		ans := distances(edges, start-1, vertices)
		if len(ans) == 1 {
			fmt.Fprintf(writer, "%d\n", ans)
			continue
		}
		fmt.Fprintf(writer, "%d", ans[0])
		for i := 1; i < len(ans); i++ {
			fmt.Fprintf(writer, " %d", ans[i])
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
