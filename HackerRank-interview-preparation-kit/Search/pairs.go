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
 * Complete the 'pairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */

func pairs(k int32, arr []int32) int32 {
	/*
	   1. Start with an example
	   2. Brute force
	   3. Have a sketch of an solutions
	   4. Test the input
	   5. Code up
	*/

	/*

	   k = 2
	   arr = [1, 3, 5, 6]

	   ->
	   1-3 = -2
	   3-5 = -2
	*/

	/*
	   abs := func(x int32) int32 {
	       if x < 0 {
	           return -x
	       }
	       return x
	   }
	*/

	var pairs int32
	/*
	   for i := 0; i < len(arr); i++ {
	       for j := i + 1; j < len(arr); j++ {
	           if abs(arr[i] - arr[j]) == k {
	               pairs++
	           }
	       }
	   }
	*/
	// O(n^2)

	// Let's optimize with hashMap
	// We will store all elements in a hashmap,
	// And then lookup the missing element
	// Elements wont be counted twice as we will lookup only the bigger element
	elems := map[int32]int32{}
	for _, v := range arr {
		elems[v]++
	}

	for _, v := range arr {
		if count, ok := elems[v+k]; ok {
			pairs += count
		}
	}

	return pairs
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

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

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
