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
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

type stack []rune

func (s *stack) push(v rune) {
	*s = append(*s, v)
}

func (s *stack) peek() rune {
	if len(*s) == 0 {
		return ' '
	}
	return (*s)[len(*s)-1]
}

func (s *stack) pop() {
	if len(*s) == 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s *stack) length() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

func isBalanced(s string) string {
	/*
	   Continuously add brackets to the stack and pop matching
	*/
	// optimization
	// if len(s) % 2 == 1 {
	//     return "NO"
	// }

	var st stack
	for _, v := range []rune(s) {
		switch {
		case v == '(', v == '[', v == '{':
			st.push(v)
		case v == ')' && st.peek() == '(',
			v == ']' && st.peek() == '[',
			v == '}' && st.peek() == '{':
			st.pop()
		default:
			st.push(v)
		}
	}
	if st.length() == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

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
