package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'balancedForest' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY c
 *  2. 2D_INTEGER_ARRAY edges
 */

func graphToTree(c []int64, edges [][]int64) ([]int64, [][]int64) {
	adj := map[int64][]int64{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	var treeEdges [][]int64
	visited := make([]bool, len(c))
	var dfs func(int64)
	dfs = func(n int64) {
		visited[n-1] = true
		if children, ok := adj[n]; ok {
			for _, child := range children {
				if !visited[child-1] {
					treeEdges = append(treeEdges, []int64{n, child})
					dfs(child)
				}
			}
		}
	}
	dfs(1) // nodes are 1-based
	return c, treeEdges
}

func balancedForest(c []int64, edges [][]int64) int64 {
	//fmt.Printf("Before c: %v\nedges: %v\n", c, edges)
	c, edges = graphToTree(c, edges)
	//fmt.Printf("After c: %v\nedges: %v\n", c, edges)

	// 2/ Re-map edges to map
	adj := map[int64][]int64{}
	parents := map[int64]int64{}
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		parents[v[1]] = v[0]
	}

	// 1/ Calculate sums on all nodes
	sums := map[int64]int64{}
	var dfsSum func(int64) int64
	dfsSum = func(n int64) int64 {
		children, ok := adj[n]
		if !ok {
			// record and return
			sums[n] = c[n-1]
			return c[n-1]
		}
		var r int64
		for _, child := range children {
			r += dfsSum(child)
		}
		r += c[n-1]
		sums[n] = r
		return r
	}
	dfsSum(1)

	fmt.Printf("Sums %v\n", sums)

	//fmt.Printf("Parents %v\n", parents)
	//fmt.Printf("Adj - %v\n", adj)
	//fmt.Printf("Sums %v\n", sums)

	var ans int64
	ans = math.MaxInt64
	// 3/ Iterate over all edges finding matching value,
	// bailing bfs if value too small
	// value of x (as in sums[0] = 2*x + y) cannot be less then sums[0]/2
	for _, edge := range edges {
		// What if we cut it?
		from, to := edge[0], edge[1]
		x := sums[to]
		//fmt.Printf("Trying edge %v, x = %d, root - x = %v\n", edge, x, sums[1]-x)
		// Update sum and adj map and parents
		hasParent := true
		for parent := from; hasParent; parent, hasParent = parents[parent] {
			sums[parent] -= x
		}
		/*
		   We now have x and sums[1] as the two values for amount of nodes
		*/
		var canCutToSumDFS func(node int64, topSum int64, wantSum int64) bool
		visited := make([]bool, len(c)+1)
		visited[to] = true // as we operate as we cut this edge already
		canCutToSumDFS = func(node int64, topSum int64, wantSum int64) bool {
			if children, ok := adj[node]; ok {
				for _, child := range children {
					if visited[child] {
						continue
					}
					visited[child] = true
					if topSum-sums[child] == wantSum || sums[child] == wantSum || canCutToSumDFS(child, topSum, wantSum) {
						return true
					}
				}
			}
			return false
		}

		// Backtrack damage
		hasParent = true
		for parent := from; hasParent; parent, hasParent = parents[parent] {
			sums[parent] += x
		}
	}
	if ans == math.MaxInt64 {
		return int64(-1)
	} else {
		return ans
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int64(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int64(nTemp)

		cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var c []int64

		for i := 0; i < int(n); i++ {
			cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
			checkError(err)
			cItem := int64(cItemTemp)
			c = append(c, cItem)
		}

		var edges [][]int64
		for i := 0; i < int(n)-1; i++ {
			edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var edgesRow []int64
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int64(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != 2 {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		result := balancedForest(c, edges)

		fmt.Fprintf(writer, "%d\n", result)
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
