package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	/*
	   Approach: expand the array length until it
	   reaches full arr length in reverse approach
	*/
	if len(arr) == 0 {
		return 0
	}
	memo := make([]int32, len(arr))
	for k := len(arr) - 1; k >= 0; k-- {
		switch k {
		case len(arr) - 1:
			if arr[k] > 0 {
				memo[k] = arr[k]
			} else {
				memo[k] = 0
			}
		case len(arr) - 2:
			var max int32
			for _, v := range []int32{
				arr[len(arr)-2],
				arr[len(arr)-1]} {
				if max < v {
					max = v
				}
			}
			memo[k] = max
		case len(arr) - 3:
			var max int32
			for _, v := range []int32{
				arr[len(arr)-3],
				arr[len(arr)-2],
				arr[len(arr)-1],
				arr[len(arr)-3] + arr[len(arr)-1]} {
				if max < v {
					max = v
				}
			}
			memo[k] = max
		default:
			var max int32
			for _, v := range []int32{
				arr[k] + memo[k+2],
				memo[k+1]} {
				if max < v {
					max = v
				}
			}
			memo[k] = max
		}
	}
	return memo[0]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := maxSubsetSum(arr)

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
