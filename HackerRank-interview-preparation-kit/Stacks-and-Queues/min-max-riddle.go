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

// Complete the riddle function below.
func riddle(arr []int64) []int64 {
	/*
	   Approach:
	   1/ Find closes bigger element for each on left and right
	   2/ With knowledge that particular arr[i] is a potential win for segment right-left[i]-1, fill known answers and then can fill in the rest of missing places by observation that if smaller window win was not found, it means there is no better combination, thus it should be ans[i+1]

	   All in all will get to O(n^2)
	*/

	// 1.1/ Closest smaller element on the left
	left := make([]int, len(arr)) // index of the element on the left that is smaller
	var stack []int               // storing indexes of elements
	for i, v := range arr {
		// pop elements until we find one that is smaller
		for ; len(stack) > 0 && arr[stack[len(stack)-1]] >= v; stack = stack[:len(stack)-1] {
		}

		if len(stack) == 0 {
			left[i] = -1 // there is no element smaller on the left
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	// 1.2/ Closest smaller element on the right
	stack = stack[:0] // keeping the memmory allocated
	right := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		for ; len(stack) > 0 && arr[stack[len(stack)-1]] >= v; stack = stack[:len(stack)-1] {
		}
		if len(stack) == 0 {
			right[i] = len(arr) // there is no element smaller on the right
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 2/ Fill in the winning combinations that we know directly
	ans := make([]int64, len(arr))
	for i := range ans {
		ans[i] = math.MinInt64
	}
	for i, v := range arr {
		interval := right[i] - left[i] - 1
		if ans[interval-1] < v {
			ans[interval-1] = v
		}
	}
	for i := len(ans) - 2; i >= 0; i-- {
		if ans[i] < ans[i+1] {
			ans[i] = ans[i+1]
		}
	}

	return ans
}

func main() {
	// Correction for Test Case 2 that fails with Runtime Error
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*1024)
	//reader := bufio.NewReaderSize(os.Stdin, MAX_INPUT_SIZE*64)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	res := riddle(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
