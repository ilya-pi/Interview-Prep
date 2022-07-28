package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	if len(s) == 0 {
		return "YES"
	}

	// Count frequencies of chars
	freqs := make(map[rune]int32)
	for _, v := range s {
		freqs[v]++
	}

	// Reduce frequencies
	f := make(map[int32]int32)
	for _, v := range freqs {
		f[v]++
	}

	// More then 2 different frequencies
	if len(f) > 2 {
		return "NO"
	}

	// Just right
	if len(f) == 1 {
		return "YES"
	}

	// See if we can "fix" string with removing 1 char
	elems := [][]int32{}
	for k, v := range f {
		elems = append(elems, []int32{k, v})
	}

	if !(elems[0][1] == 1 && (elems[0][0]-elems[1][0])*(elems[0][0]-elems[1][0]) == 1 ||
		elems[1][1] == 1 && (elems[0][0]-elems[1][0])*(elems[0][0]-elems[1][0]) == 1 ||
		elems[0][0] == 1 && elems[0][1] == 1 ||
		elems[1][0] == 1 && elems[1][1] == 1) {
		return "NO"
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
