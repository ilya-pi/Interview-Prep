package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'reverseShuffleMerge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func reverseShuffleMerge(s string) string {
	/*
	   Approach:

	   Since the initial string is reversed in the result given string, we will try to pick the string from the initial string going right to left

	   We can understand what characters we need by calculating the overall amount of characters and /2

	   We can then keep track of the "still" available characters we have in the string to build the string from

	   On each character: we see if we need it in the resulting string, if yes -> add it. If on adding it - we making the string lexicographically bigger (rune is smaller then already added, then we will see if the already added runes can be added later, and if yes -> we will clean them)
	*/

	runes := []rune(s)
	// Calculate all available and needed characters
	available := map[rune]int{}
	for _, v := range runes {
		available[v]++
	}

	var tlen int
	needed := map[rune]int{}
	for k, v := range available {
		needed[k] = v / 2
		tlen += v / 2
	}

	var r []rune
	for i := len(runes) - 1; i >= 0; i-- {
		v := runes[i]
		available[v]--
		if needed[v] > 0 {
			// We need this character, so we'll see if we can add it without compromising other characters for lexicographical order
			// We go from the back and see what characters we can move to still have a lexicographically smaller string
			for j := len(r) - 1; j >= 0; j-- {
				a := r[j]
				if a == '0' {
					continue
				}
				if a > v && available[a] > needed[a] {
					r[j] = '0'
					needed[a]++
				} else {
					// cannot "backtrack" anymore
					break
				}
			}
			needed[v]--
			r = append(r, v)
		}
	}

	cleaned := make([]rune, tlen)
	i := 0
	for _, v := range r {
		if v != '0' {
			cleaned[i] = v
			i++
		}
	}
	return string(cleaned)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := reverseShuffleMerge(s)

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
