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
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	/*
	   Approach:
	   Naive approach would be to sum first time n manually and then do the tail of super digit calculation

	   1. Sum all digits in n
	   2. Multiply by k
	   3. Repeat superdigit for new value until resolution
	*/

	arr := []rune(n)
	var v int64
	for _, e := range arr {
		v += int64(e - '0')
	}
	v *= int64(k)

	superDigit := func(v int64) int64 {
		var r int64
		for v > 9 {
			r += v % 10
			v /= 10
		}
		r += v
		return r
	}

	for v >= 10 {
		v = superDigit(v)
	}

	return int32(v)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
