package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func (s *stack) pop() *string {
	if len(*s) == 0 {
		return nil
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &top
}

func (s *stack) peek() *string {
	if len(*s) == 0 {
		return nil
	}
	r := (*s)[len(*s)-1]
	return &r
}

func (s *stack) length() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

type queue struct {
	st1 stack
	st2 stack
}

func (q *queue) push(v string) {
	if q.st2.length() != 0 {
		q.flipLeft()
	}
	q.st1.push(v)
}

func (q *queue) dequeue() string {
	if q.st1.length() != 0 {
		q.flipRight()
	}
	return *(q.st2.pop())
}

func (q *queue) peek() string {
	q.flipRight()
	return *(q.st2.peek())
}

func (q *queue) flipLeft() {
	for q.st2.length() > 0 {
		q.st1.push(*(q.st2.pop()))
	}
}

func (q *queue) flipRight() {
	for q.st1.length() > 0 {
		q.st2.push(*(q.st1.pop()))
	}
}

func process(arr [][]string, writer *bufio.Writer) { // []string {
	//var ans []string

	var q queue
	for _, v := range arr {
		switch v[0] {
		case "1":
			q.push(v[1])
		case "2":
			q.dequeue()
		case "3":
			fmt.Fprintf(writer, "%v\n", q.peek())
			//ans = append(ans, q.peek())
		}
	}
	//return ans
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

	var arr [][]string

	for tItr := 0; tItr < int(q); tItr++ {
		s := readLine(reader)

		e := strings.Split(s, " ")
		arr = append(arr, e)
	}

	process(arr, writer)
	// for _, v := range r {
	//     fmt.Fprintf(writer, "%v\n", v)
	// }

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
