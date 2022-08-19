package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack []rune

func (s Stack) push(v rune) Stack {
	s = append(s, v)
	return s
}

func (s Stack) pop() (Stack, *rune) {
	if len(s) == 0 {
		return s, nil
	}
	return s[:len(s)-1], &s[len(s)-1]
}

func (s Stack) peek() *rune {
	if len(s) == 0 {
		return nil
	}
	return &s[len(s)-1]
}

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func isBalanced(s string) string {
	/*
	   Approach: I beleive we can use stack and pop the last
	   element if we meet "oppostite" bracket at the top of the
	   current stack
	*/
	st := Stack{}

	matching := func(v rune) rune {
		switch v {
		case ')':
			return '('
		case '}':
			return '{'
		case ']':
			return '['
		}
		return ' '
	}

	for _, v := range []rune(s) {
		switch v {
		case '(', '[', '{':
			st = st.push(v)
		case ')', ']', '}':
			if top := st.peek(); top != nil && *top == matching(v) {
				st, _ = st.pop()
				continue
			} else {
				return "NO"
			}
		default:
			// Unexpected rune altogether
			return "NO"
		}
	}

	if st.peek() == nil {
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
