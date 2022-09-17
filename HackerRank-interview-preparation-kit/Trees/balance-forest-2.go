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
	adj := map[int64][]int64{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// 1/ Calculate sums on all nodes
	sums := map[int64]int64{}
	v := make([]bool, len(c)+1)
	var dfsSum func(int64) int64
	dfsSum = func(n int64) int64 {
		v[n] = true
		children, ok := adj[n]
		if !ok {
			// record and return
			sums[n] = c[n-1]
			return c[n-1]
		}
		var r int64
		for _, child := range children {
			if v[child] {
				continue
			}
			r += dfsSum(child)
		}
		r += c[n-1]
		sums[n] = r
		return r
	}
	dfsSum(1)

	//fmt.Printf("Sums %v\n", sums)

	var ans int64
	ans = math.MaxInt64

	total := sums[1]

	seenSums := map[int64]bool{}
	leftSums := map[int64]bool{}

	v2 := make([]bool, len(c)+1)
	var solve func(n int64)
	solve = func(n int64) {
		v2[n] = true
		sum := sums[n]

		// scenarios

		// scenario: 1 - it is a smaller tree, so we have 2 times x in the rest
		if (total-sum)%2 == 0 {
			x := (total - sum) / 2
			min := x - sum
			if leftSums[x] ||
				seenSums[x+sum] {
				if min >= 0 && min < ans {
					ans = min
				}
			}
		}

		// scenario: 2 - it is one of the equal trees of size x
		x := sum
		min := x - (total - 2*x)
		if leftSums[x] ||
			leftSums[total-2*x] ||
			seenSums[x*2] ||
			seenSums[total-x] {
			if min >= 0 && min < ans {
				ans = min
			}
		}

		seenSums[sum] = true
		for _, ch := range adj[n] {
			if v2[ch] {
				continue
			}
			solve(ch)
		}
		delete(seenSums, sum)
		leftSums[sum] = true
	}
	solve(1)

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
