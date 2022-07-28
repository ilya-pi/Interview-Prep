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
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Count all substrings
	anagrams := map[string]int32{}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j <= len(runes); j++ {
			tmp := runes[i:j]
			key := make([]rune, len(tmp))
			copy(key, tmp)
			sort.Slice(key, func(i int, j int) bool { return key[i] < key[j] })
			anagrams[string(key)]++
		}
	}

	// Count possible amount of pairs
	var r int32
	for _, v := range anagrams {
		r += v * (v - 1) / 2
	}

	return r
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
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

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
