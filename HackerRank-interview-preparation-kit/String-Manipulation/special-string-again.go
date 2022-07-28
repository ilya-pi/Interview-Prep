package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//[m n o non n ono o p opo o]

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	// Go rune by rune and examine every string for potentiallity of being special and keep appending to it untill we cannot build longer special string
	isSpecialString := func(s string) bool {
		if len(s) == 0 {
			return false
		}

		if len(s) == 1 {
			return true
		}

		if len(s)%2 == 1 {
			s = s[:len(s)/2] + s[len(s)/2+1:]
		}

		r := []rune(s)[0]
		for _, v := range s {
			if r != v {
				return false
			}
		}
		return true
	}

	canBeSpecialString := func(s string) bool {
		if len(s) == 0 {
			return true
		}

		if len(s) == 1 {
			return true
		}

		var midRunePosition int
		r := []rune(s)[0]
		for i := 1; i < len([]rune(s)); i++ {
			switch {
			case midRunePosition == 0 && r != []rune(s)[i]:
				// found exceptional mid char
				midRunePosition = i
			case midRunePosition != 0 && r != []rune(s)[i]:
				// a second exception -> string is not special
				return false
			case midRunePosition != 0 && r == []rune(s)[i] && i > 2*midRunePosition:
				// exception char is not in the middle anymore
				return false
			}
		}
		return true
	}

	// Go rune by rune and examine every string for potentiallity of being special and keep appending to it untill we cannot build longer special string

	var special []string
	var candidates []string
	for _, v := range s {
		// Filter through current candidates
		tmp := candidates[:0]
		for _, candidate := range candidates {
			if isSpecialString(candidate) {
				special = append(special, candidate)
			}

			if canBeSpecialString(candidate + string(v)) {
				tmp = append(tmp, candidate+string(v))
			}
		}
		candidates = tmp

		// Add this char to the list of candidates
		candidates = append(candidates, string(v))
	}

	for _, candidate := range candidates {
		if isSpecialString(candidate) {
			special = append(special, candidate)
		}
	}

	fmt.Printf("%v\n", special)

	return int64(len(special))

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

	s := readLine(reader)

	result := substrCount(n, s)

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
