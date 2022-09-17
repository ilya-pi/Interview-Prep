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
 * Complete the 'minTime' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY roads
 *  2. INTEGER_ARRAY machines
 */

func minTime(roads [][]int32, machines []int32) int32 {
	/*
	   Approach:

	   Run dfs from first node, while choosing whether it is better to
	   cut at path of cut at a later seen moment on exiting node, depending on
	   whether it is infested or not and whether it is cheaper to cut later for either of the infested paths
	*/

	// 1. Preparing information
	adj := map[int32]map[int32]int32{} // map "from" -> map "to" -> "weight"
	for _, road := range roads {
		a, b, weight := road[0], road[1], road[2]
		if adj[a] == nil {
			adj[a] = map[int32]int32{}
		}
		if adj[b] == nil {
			adj[b] = map[int32]int32{}
		}

		adj[a][b] = weight
		adj[b][a] = weight
	}
	isMachine := map[int32]bool{}
	for _, machine := range machines {
		isMachine[machine] = true
	}
	// 2. Dfs with picking the paths to count
	visited := make([]bool, len(roads)+1) // there will be one less road then cities
	var ans int32
	var dfs func(int32, int32) int32
	dfs = func(n int32, timeToComeToN int32) int32 {
		visited[n] = true
		var timeToBeSafeFromInfestedNeighbours int32
		// as we can chose to cut the link later in the graph
		var mostExpensiveInfestedNeighbourToGetRidOf int32

		for to, timeToComeToTo := range adj[n] {
			if visited[to] {
				continue
			}

			timeToBeSafeFromIfInfested := dfs(to, timeToComeToTo)
			timeToBeSafeFromInfestedNeighbours += timeToBeSafeFromIfInfested
			if timeToBeSafeFromIfInfested > mostExpensiveInfestedNeighbourToGetRidOf {
				mostExpensiveInfestedNeighbourToGetRidOf = timeToBeSafeFromIfInfested
			}
		}
		// Are we exiting from an infested node?
		if isMachine[n] {
			// Then we need to get rid of all infested neighbours
			ans += timeToBeSafeFromInfestedNeighbours
			// What is time to get rid of us?
			return timeToComeToN
		} else {
			// If we are exiting from a "regular" node on dfs,
			// we could keep the most expensive path to cut it potentially later
			// when we exit from an infected node
			ans += timeToBeSafeFromInfestedNeighbours - mostExpensiveInfestedNeighbourToGetRidOf
			// As we are returning into a regular node, we have an option to either cut it or cut the other worst connection (if exists)
			if mostExpensiveInfestedNeighbourToGetRidOf < timeToComeToN {
				return mostExpensiveInfestedNeighbourToGetRidOf
			} else {
				return timeToComeToN
			}
		}
	}
	// Run dfs from fist machine with 0 cost to come into it
	dfs(0, 0)
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var roads [][]int32
	for i := 0; i < int(n)-1; i++ {
		roadsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var roadsRow []int32
		for _, roadsRowItem := range roadsRowTemp {
			roadsItemTemp, err := strconv.ParseInt(roadsRowItem, 10, 64)
			checkError(err)
			roadsItem := int32(roadsItemTemp)
			roadsRow = append(roadsRow, roadsItem)
		}

		if len(roadsRow) != 3 {
			panic("Bad input")
		}

		roads = append(roads, roadsRow)
	}

	var machines []int32

	for i := 0; i < int(k); i++ {
		machinesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		machinesItem := int32(machinesItemTemp)
		machines = append(machines, machinesItem)
	}

	result := minTime(roads, machines)

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
