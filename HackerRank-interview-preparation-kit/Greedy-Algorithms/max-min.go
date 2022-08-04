package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'maxMin' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */

func maxMin(k int32, arr []int32) int32 {
	// Approach: go backwards and maintain unfairness measurement while removing elements from the array, either at the beginning or at the end.
	// Would it always yield correct result?  -> No
	// But, the resulting array will be a window in the sorted array, as it will have max and min elements that will contribute to unfairness and the rest will be inbetween them. Thus ->
	// 1. Sort array
	// 2. Slide a window of k elements while calculating unfairness and taking the lowest one
	// So it will get to O(nlogn) due to sorting
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	minMax := arr[k-1] - arr[0]
	for i := int32(1); i <= int32(len(arr))-k; i++ {
		if minMax > arr[i+k-1]-arr[i] {
			minMax = arr[i+k-1] - arr[i]
		}
	}
	return minMax
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

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := maxMin(k, arr)

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
