package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minTime function below.
func minTime(machines []int64, goal int64) int64 {
	/*
	   Approach:
	   Find prod in Max days
	   Then binary search to find required days
	*/

	prodInDays := func(days int64) int64 {
		var res int64
		for _, v := range machines {
			res += days / v
		}
		return res
	}

	var maxDaysMachine int64
	for _, v := range machines {
		if v > maxDaysMachine {
			maxDaysMachine = v
		}
	}
	left, right := int64(0), maxDaysMachine
	produce := prodInDays(right)
	for goal > produce {
		left = right
		right *= 2
		produce = prodInDays(right)
	}

	// minDays is between left and right
	for right-left > 1 {
		mid := (right + left) / 2
		produce = prodInDays(mid)
		if goal > produce {
			left = mid + 1
		} else if goal < produce {
			right = mid // I feel unease about this, as one day can tip it over the goal
		} else if goal == produce {
			right = mid
		}
	}

	produce = prodInDays(right)
	// Correction, as machine production is non-linear
	// and the same produce can be reached in multiple ways
	for produce >= goal {
		right--
		produce = prodInDays(right)
	}
	return right + 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nGoal := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nGoal[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	goal, err := strconv.ParseInt(nGoal[1], 10, 64)
	checkError(err)

	machinesTemp := strings.Split(readLine(reader), " ")

	var machines []int64

	for i := 0; i < int(n); i++ {
		machinesItem, err := strconv.ParseInt(machinesTemp[i], 10, 64)
		checkError(err)
		machines = append(machines, machinesItem)
	}

	ans := minTime(machines, goal)

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
