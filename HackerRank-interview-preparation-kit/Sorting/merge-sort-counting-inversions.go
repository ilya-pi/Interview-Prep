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
 * Complete the 'countInversions' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func countInversions(arr []int32) int64 {
	// Write your code here
	/*
	   Naive approach first:

	   1. For all elements, if there is a smaller element in the rest -> add the inversions and swap elements
	*/
	/*
	   var swaps int
	   for i := 0; i < len(arr); i++ {
	       // Find lesser element
	       minIndex := i
	       for k := i + 1; k < len(arr); k++ {
	           if arr[k] < arr[minIndex] {
	               minIndex = k
	           }
	       }
	       // Move all elements
	       min := arr[minIndex]
	       for j := minIndex; j > i; j-- {
	           arr[j] = arr[j-1]
	       }
	       arr[i] = min
	       swaps += minIndex - i
	   }
	   return int64(swaps)
	*/
	/*
	   var swaps int
	   for i := 0; i < len(arr); i++ {
	       for j := i+1; j < len(arr); j++ {
	           if arr[i] > arr[j] {
	               swaps++
	           }
	       }
	   }
	   return int64(swaps)
	*/

	var count int

	merge := func(a, b []int32) []int32 {
		i, j, k := 0, 0, 0

		r := make([]int32, len(a)+len(b))
		for i < len(a) && j < len(b) {
			if a[i] <= b[j] {
				r[k] = a[i]
				i++
			} else {
				r[k] = b[j]
				count += len(a) - i
				j++
			}
			k++
		}

		for ; i < len(a); i, k = i+1, k+1 {
			r[k] = a[i]
		}
		for ; j < len(b); j, k = j+1, k+1 {
			r[k] = b[j]
		}
		return r
	}

	var mergeSort func([]int32) []int32
	mergeSort = func(a []int32) []int32 {
		if len(a) <= 1 {
			return a
		}
		// Split
		mid := len(a) / 2
		left := mergeSort(a[:mid])
		right := mergeSort(a[mid:])
		// Merge
		return merge(left, right)
	}
	_ = mergeSort(arr)
	return int64(count)
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

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := countInversions(arr)

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
