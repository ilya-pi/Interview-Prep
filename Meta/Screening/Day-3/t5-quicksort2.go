package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func print(arr []int32) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func quickSort2(n int32, arr []int32) {
	partition := func(arr []int32, low, high int32) int32 {
		pivot := arr[low]
		i := high + 1 // last partitioned element
		for j := high; j > low; j-- {
			if arr[j] > pivot {
				i--
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		// partition the pivot element
		arr[i-1], arr[low] = arr[low], arr[i-1]
		return i - 1
	}

	var qs func([]int32, int32, int32)
	qs = func(arr []int32, low, high int32) {
		if low >= high {
			return
		}
		pi := partition(arr, low, high)
		qs(arr, low, pi-1)
		qs(arr, pi+1, high)
		print(arr[low : high+1])
	}
	qs(arr, 0, int32(len(arr))-1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

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

	quickSort2(n, arr)
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
