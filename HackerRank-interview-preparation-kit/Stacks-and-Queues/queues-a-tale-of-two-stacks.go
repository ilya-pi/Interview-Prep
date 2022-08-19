package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Queue []string

func (q Queue) push(v string) Queue {
	q = append(q, v)
	return q
}

func (q Queue) dequeue() (Queue, *string) {
	if len(q) == 0 {
		return q, nil
	}
	return q[1:], &q[0]
}

func (q Queue) peek() *string {
	if len(q) == 0 {
		return nil
	}
	return &q[0]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)
	// defer stdout.Close()
	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	q := Queue{}

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)
		switch s[0] {
		case '1':
			arr := strings.Split(s, " ")
			q = q.push(arr[1])
		case '2':
			q, _ = q.dequeue()
		case '3':
			fmt.Printf("%v\n", *(q.peek()))
		}
	}
	// writer.Flush()
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
