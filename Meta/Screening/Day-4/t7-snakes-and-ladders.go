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
 * Complete the 'quickestWayUp' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY ladders
 *  2. 2D_INTEGER_ARRAY snakes
 */

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	size := 100
	// 22:27
	// 22:37 impl
	// 22:41 error - we don't have to climb the ladders and snakes
	/*
	   1/ Build adj matrix, from each node we can go to the next 6, add ladders and add snakes
	   2/ Run BFS and exit on the node 100, tracking amount of steps
	*/
	adj := make([][]int32, size)
	teleports := make(map[int32]bool)
	// create roll the die connections
	for i := 0; i < len(adj); i++ {
		for j := 1; j <= 6 && i+j < len(adj); j++ {
			adj[i] = append(adj[i], int32(i+j))
		}
	}
	// add ladders to adj
	for _, ladder := range ladders {
		from, to := ladder[0]-1, ladder[1]-1 // they are 1-based
		// erase previous connections
		adj[from] = []int32{to}
		// add fast step
		teleports[from] = true
	}
	// add snakes
	for _, snake := range snakes {
		from, to := snake[0]-1, snake[1]-1
		adj[from] = []int32{to}
		teleports[from] = true
	}

	// run bfs till we hit elem size
	type NodeDistance struct {
		node     int32
		distance int32
	}
	q := []NodeDistance{{0, 0}}
	visited := make([]bool, size)
	for len(q) > 0 {
		a := q[0]
		q = q[1:]

		if a.node == int32(size-1) {
			return a.distance
		}
		distance := a.distance + 1
		for _, child := range adj[a.node] {
			if visited[child] {
				continue
			}
			visited[child] = true
			if teleports[a.node] {
				q = append(q, NodeDistance{child, distance - 1})
			} else {
				q = append(q, NodeDistance{child, distance})
			}
		}
	}

	return int32(-1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var ladders [][]int32
		for i := 0; i < int(n); i++ {
			laddersRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var laddersRow []int32
			for _, laddersRowItem := range laddersRowTemp {
				laddersItemTemp, err := strconv.ParseInt(laddersRowItem, 10, 64)
				checkError(err)
				laddersItem := int32(laddersItemTemp)
				laddersRow = append(laddersRow, laddersItem)
			}

			if len(laddersRow) != 2 {
				panic("Bad input")
			}

			ladders = append(ladders, laddersRow)
		}

		mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		m := int32(mTemp)

		var snakes [][]int32
		for i := 0; i < int(m); i++ {
			snakesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var snakesRow []int32
			for _, snakesRowItem := range snakesRowTemp {
				snakesItemTemp, err := strconv.ParseInt(snakesRowItem, 10, 64)
				checkError(err)
				snakesItem := int32(snakesItemTemp)
				snakesRow = append(snakesRow, snakesItem)
			}

			if len(snakesRow) != 2 {
				panic("Bad input")
			}

			snakes = append(snakes, snakesRow)
		}

		result := quickestWayUp(ladders, snakes)

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
