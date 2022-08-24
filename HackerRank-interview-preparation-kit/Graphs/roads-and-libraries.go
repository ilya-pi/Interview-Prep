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
 * Complete the 'roadsAndLibraries' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c_lib
 *  3. INTEGER c_road
 *  4. 2D_INTEGER_ARRAY cities
 */

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	/*
	   It seems we need to establish how many enclaves there are and whether it is cheaper to connect them or to build a separate library in there.

	   We will try to run BFS from all nodes checking if we saw them already, counting the amount of enclaves and edges we used to traverse with bfs
	*/
	edges := map[int32][]int32{}
	for _, v := range cities {
		// cities are 1 based
		a, b := v[0]-1, v[1]-1
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}

	var libs int64
	var roads int64
	visited := make([]bool, n)
	for i := int32(0); i < n; i++ {
		if visited[i] { // Already provided with library
			continue
		}

		q := []int32{i}
		visited[i] = true
		libs++ // new enclave center
		for len(q) > 0 {
			k := q[0]
			q = q[1:]
			for _, l := range edges[k] {
				if visited[l] { // Already connected
					continue
				}
				if c_lib > c_road {
					q = append(q, l)
					visited[l] = true
					roads++
				} else { // cheaper to build there a library instead
					visited[l] = true
					libs++
				}
			}
		}
	}

	return libs*int64(c_lib) + roads*int64(c_road)
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

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

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
