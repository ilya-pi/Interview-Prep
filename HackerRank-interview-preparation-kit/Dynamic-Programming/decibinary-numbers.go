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
 * Complete the 'decibinaryNumbers' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER x as parameter.
 */

// const DECIBINARY_LIMIT = 1000000000000
// 10000000000000000
// 10000000000000000
const DECIBINARY_LIMIT = 10000000
const DECIMAL_LIMIT = 100000

var numbers [][]int64
var xarr []int64

/*

 Optimizations:

 [ ] 2^n values
 [ ] Memo smaller decibinary values

*/

func init() {
	// Gen all decibinary till DECIBINARY_LIMIT
	numbers = make([][]int64, DECIMAL_LIMIT)
	xarr = make([]int64, DECIMAL_LIMIT)
	db2Dc := func(prev int64, db int64) int64 {
		if prev != 0 && (db-1)%10 != 9 {
			return prev + 1
		}
		if db == 0 {
			return 0
		}
		var r int64
		n := int64(1)
		for db > 0 {
			r += (db % 10) * n
			db /= 10
			n *= 2
		}
		return r
	}
	var prev int64
	for i := int64(0); i < DECIBINARY_LIMIT; i++ {
		dc := db2Dc(prev, i)
		prev = dc
		numbers[dc] = append(numbers[dc], i)
	}
	var sum int64
	for i := int64(0); i < int64(len(xarr)); i++ {
		sum += int64(len(numbers[i]))
		xarr[i] = sum
	}
}

// find returns index of the first xarr that is > x
func find(x int64) int64 {
	left, right := int64(0), int64(len(xarr))-1
	for right-left > 1 {
		mid := left + (right-left)/2
		if xarr[mid] < x {
			left = mid
		} else if xarr[mid] > x {
			right = mid
		} else if xarr[mid] == x {
			return mid
		}
	}
	return left
}

func decibinaryNumbers(x int64) int64 {
	/*
	   Approach:

	   Generate decibinary numbers till high enough level and then run sum throught them.

	   1/ Done in init, as it is shared between querries
	   2/ Can be optimized with some form of tree

	*/

	// Binary search the value in array
	ind := find(x) // x is 1 based
	if xarr[ind] == x {
		dcs := numbers[ind]
		return dcs[len(dcs)-1]
	} else {
		ind2 := x - xarr[ind] - 1
		return numbers[ind+1][ind2]
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
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		x, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := decibinaryNumbers(x)

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
