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
 * Complete the 'minimumMoves' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING_ARRAY grid
 *  2. INTEGER startX
 *  3. INTEGER startY
 *  4. INTEGER goalX
 *  5. INTEGER goalY
 */

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	/*
	   BFS till we find the end

	   1. Map grid indo adjacency table
	   2. Have a type with node and steps
	   3. Run BFS from start
	*/

	// 1. Map grid into adjacency table
	adj := make([][]rune, len(grid))
	for i, v := range grid {
		adj[i] = []rune(v)
	}

	// 2. Type to hold node and steps
	type NodeSteps struct {
		x, y  int32
		steps int32
	}

	// 3.1 Where we can move from a point x, y
	// with boundaries check
	maxX, maxY := int32(len(adj)), int32(len(adj[0]))
	canMoveTo := func(x, y int32) [][]int32 {
		var r [][]int32
		for x1 := x - 1; x1 >= 0; x1-- {
			if adj[x1][y] == 'X' {
				break
			}
			r = append(r, []int32{x1, y})
		}
		for x1 := x + 1; x1 < maxX; x1++ {
			if adj[x1][y] == 'X' {
				break
			}
			r = append(r, []int32{x1, y})
		}
		for y1 := y - 1; y1 >= 0; y1-- {
			if adj[x][y1] == 'X' {
				break
			}
			r = append(r, []int32{x, y1})
		}
		for y1 := y + 1; y1 < maxY; y1++ {
			if adj[x][y1] == 'X' {
				break
			}
			r = append(r, []int32{x, y1})
		}
		return r
	}

	// 3. BFS
	q := []NodeSteps{{startX, startY, 0}}
	visited := make([][]bool, len(adj))
	for i := range visited {
		visited[i] = make([]bool, len(adj[i]))
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.x == goalX && n.y == goalY {
			return n.steps
		}

		steps := n.steps + 1

		for _, d := range canMoveTo(n.x, n.y) {
			x, y := d[0], d[1]
			// visited
			if visited[x][y] {
				continue
			}
			// visit and add to queue
			visited[x][y] = true
			q = append(q, NodeSteps{x, y, steps})
		}
	}
	return -1
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

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	startXTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	startX := int32(startXTemp)

	startYTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	startY := int32(startYTemp)

	goalXTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	goalX := int32(goalXTemp)

	goalYTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	goalY := int32(goalYTemp)

	result := minimumMoves(grid, startX, startY, goalX, goalY)

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
