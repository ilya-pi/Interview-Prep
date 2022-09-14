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
 * Complete the 'maxRegion' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY grid as parameter.
 */

func maxRegion(grid [][]int32) int32 {
	/*
	   Approach:
	   For all nodes we run dfs counting region size with visited and store max
	*/
	visited := make([][]bool, len(grid))
	for i, v := range grid {
		visited[i] = make([]bool, len(v))
	}
	n := int32(len(grid))
	m := int32(len(grid[0])) // hoping input is correct

	var max int32
	var dfs func(i, j int32, acc *int32)
	dfs = func(i, j int32, acc *int32) {
		if grid[i][j] != 1 {
			return
		}
		visited[i][j] = true
		*acc += 1
		for _, adj := range [][]int32{{i - 1, j - 1}, {i, j - 1}, {i + 1, j - 1},
			{i - 1, j}, {i + 1, j},
			{i - 1, j + 1}, {i, j + 1}, {i + 1, j + 1}} {
			x, y := adj[0], adj[1]
			if x < 0 || x >= n || y < 0 || y >= m {
				// no node
				continue
			}
			if visited[x][y] {
				continue
			}
			// visit logic happens on entry, otherwise we count twice
			//visited[x][y] = true
			if grid[x][y] == 1 {
				//*acc += 1
				dfs(x, y, acc)
			}
		}
	}
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < m; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				var regionSize int32
				dfs(i, j, &regionSize)
				if regionSize > max {
					max = regionSize
				}
			}
		}
	}
	return max
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

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var grid [][]int32
	for i := 0; i < int(n); i++ {
		gridRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var gridRow []int32
		for _, gridRowItem := range gridRowTemp {
			gridItemTemp, err := strconv.ParseInt(gridRowItem, 10, 64)
			checkError(err)
			gridItem := int32(gridItemTemp)
			gridRow = append(gridRow, gridItem)
		}

		if len(gridRow) != int(m) {
			panic("Bad input")
		}

		grid = append(grid, gridRow)
	}

	res := maxRegion(grid)

	fmt.Fprintf(writer, "%d\n", res)

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
