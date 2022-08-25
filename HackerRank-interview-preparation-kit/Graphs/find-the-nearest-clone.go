package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the findShortest function below.

/*
 * For the unweighted graph, <name>:
 *
 * 1. The number of nodes is <name>Nodes.
 * 2. The number of edges is <name>Edges.
 * 3. An edge exists between <name>From[i] to <name>To[i].
 *
 */
func findShortest(graphNodes int32, graphFrom []int32, graphTo []int32, ids []int64, val int32) int32 {
	/*
	   I believe we can do a loop over all nodes that have the desired color and run BFS from that node to the desired color, marking nodes as visited if we saw them.
	   The shortest of all these bfs traversals will be the shortest path (there might be different enclaves)
	   In a nutshell we'll see each node only once making it O(n) algorithm
	*/
	adj := map[int32][]int32{}
	// Map input into more convenient form to iterate over adjacent nodes
	// Hoping len(graphFrom) == len(graphTo)
	for i := 0; i < len(graphFrom); i++ {
		adj[graphFrom[i]-1] = append(adj[graphFrom[i]-1], graphTo[i]-1)
		adj[graphTo[i]-1] = append(adj[graphTo[i]-1], graphFrom[i]-1)
	}
	visited := make([]bool, graphNodes+1)
	shortestPath := int32(-1)
	for i := int32(0); i < graphNodes; i++ {
		if ids[i] != int64(val) || visited[i] {
			// This node is not of an interesting colour to us or it was visited already
			continue
		}
		path := int32(-1)
		type node struct {
			id int32
			d  int32
		}
		q := []node{{i, 0}}
		visited[i] = true
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			newDist := n.d + 1
			if ids[n.id] == int64(val) && n.d != 0 {
				if path < 0 || path > n.d { // This is the new shortest path
					path = n.d
				}
				newDist = 0
			}
			for _, v := range adj[n.id] {
				if !visited[v] {
					visited[v] = true
					q = append(q, node{id: v, d: newDist})
				}
			}
		}
		if (shortestPath == -1 && path > 0) || (shortestPath > path) {
			shortestPath = path
		}
	}
	return shortestPath
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000*32)
	/*
		f, err := os.Open("./input05.txt")
		checkError(err)
		reader := bufio.NewReader(f)
	*/

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	graphNodesEdges := strings.Split(readLine(reader), " ")
	graphNodes, err := strconv.ParseInt(graphNodesEdges[0], 10, 64)
	checkError(err)

	graphEdges, err := strconv.ParseInt(graphNodesEdges[1], 10, 64)
	checkError(err)

	var graphFrom, graphTo []int32
	for i := 0; i < int(graphEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")
		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		graphFrom = append(graphFrom, int32(edgeFrom))
		graphTo = append(graphTo, int32(edgeTo))
	}

	idsTemp := strings.Split(readLine(reader), " ")

	var ids []int64

	for i := 0; i < int(graphNodes); i++ {
		idsItem, err := strconv.ParseInt(idsTemp[i], 10, 64)
		checkError(err)
		ids = append(ids, idsItem)
	}

	valTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	val := int32(valTemp)

	ans := findShortest(int32(graphNodes), graphFrom, graphTo, ids, val)

	fmt.Fprintf(writer, "%d\n", ans)

	writer.Flush()
}

func main2() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	f, err := os.Open("./input05.txt")
	checkError(err)
	reader := bufio.NewReader(f)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	graphNodesEdges := strings.Split(readLine(reader), " ")
	graphNodes, err := strconv.ParseInt(graphNodesEdges[0], 10, 64)
	checkError(err)

	graphEdges, err := strconv.ParseInt(graphNodesEdges[1], 10, 64)
	checkError(err)

	var graphFrom, graphTo []int32
	for i := 0; i < int(graphEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")
		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		graphFrom = append(graphFrom, int32(edgeFrom))
		graphTo = append(graphTo, int32(edgeTo))
	}

	idsTemp := strings.Split(readLine(reader), " ")

	var ids []int64

	for i := 0; i < int(graphNodes); i++ {
		idsItem, err := strconv.ParseInt(idsTemp[i], 10, 64)
		checkError(err)
		ids = append(ids, idsItem)
	}

	valTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	val := int32(valTemp)

	ans := findShortest(int32(graphNodes), graphFrom, graphTo, ids, val)

	fmt.Fprintf(writer, "%d\n", ans)

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
