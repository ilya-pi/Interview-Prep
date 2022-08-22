package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
 * Complete the 'abbreviation' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func abbreviation(a string, b string) string {
	/*
		    Example:
		    AbcDE
		    ABDE

		    Every lowercase letter is either present or not, uppercase - fixed

		for every i from end:
		    if A[0] == B[0]
		    CM == CM(a[1:], b[1:])

		       - a b c D E
		    -  1 1 1 1 0 0
		    A  0 1 1 1 0 0
		    B  0 0 1 1 0 0
		    D  0 0 0 0 1 0
		    E  0 0 0 0 0 1

	*/

	ar := []rune(a)
	br := []rune(b)
	dp := make([][]bool, len(ar)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(br)+1)
	}

	dp[0][0] = true
	for i := 0; i < len(ar); i++ {
		if unicode.IsLower(ar[i]) {
			dp[i+1][0] = dp[i][0]
		} else {
			dp[i+1][0] = false
		}
	}
	for j := 0; j < len(br); j++ {
		dp[0][j+1] = false
	}

	/*


	     - A b c D E
	   - 1 0 0 0 0 0
	   A
	   B
	   D
	   E


	     - d a B c d
	   - 1 1 1 0 0 0
	   A 0 0 1 0 0 0
	   B 0 0 0 1 1 1
	   C 0 0 0 0 1 1

	*/

	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(br); j++ {
			switch {
			case unicode.IsUpper(ar[i]) && ar[i] == br[j]:
				dp[i+1][j+1] = dp[i][j]
			case unicode.IsUpper(ar[i]) && ar[i] != br[j]:
				dp[i+1][j+1] = false
			case unicode.IsLower(ar[i]) && unicode.ToUpper(ar[i]) == br[j]:
				dp[i+1][j+1] = dp[i][j] || dp[i][j+1]
			case unicode.IsLower(ar[i]) && unicode.ToUpper(ar[i]) != br[j]:
				dp[i+1][j+1] = dp[i][j+1]
			}
		}
	}
	if dp[len(ar)][len(br)] {
		return "YES"
	} else {
		return "NO"
	}
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
		a := readLine(reader)

		b := readLine(reader)

		result := abbreviation(a, b)

		fmt.Fprintf(writer, "%s\n", result)
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
