package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) push(v int) {
	*s = append(*s, v)
}

func (s *Stack) peek() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

/*
 * Complete the 'largestRectangle' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY h as parameter.
 */
func largestRectangle(h []int32) int64 {
	/*

	      *     *
	    * **  ***
	   **********
	   ----------

	   We'll implement a bruteforce approach first, which runs in n^2
	   For each ndoe we have the best possible rectangle that begins at that square

	   For each point:
	*/

	/*
	   bestRec := func(k int) int64 {
	       if k == len(h) - 1 {
	           return int64(h[k])
	       }

	       best := int64(h[k])
	       min := h[k]
	       for i := k + 1; i < len(h); i++ {
	           if min < h[i] {
	               min = h[i]
	           }
	           rec := (int64(i) - int64(k) + 1) * int64(min)
	           if rec > best {
	               best = rec
	           }
	       }

	       return best
	   }

	   var overallBest int64
	   for i := 0; i < len(h); i++ {
	       candidate := bestRec(i)
	       if overallBest < candidate {
	           overallBest = candidate
	       }
	   }*/

	/*
	   Solution approach, we want to find left smallest and right smallest elements for each building as a potential min height
	*/
	s := Stack{}
	leftSmaller := make([]int, len(h))
	for i, v := range h {
		for ; !s.isEmpty() && h[s.peek()] >= v; s.pop() {
		}
		if s.isEmpty() {
			leftSmaller[i] = -1
		} else {
			leftSmaller[i] = s.peek()
		}
		s.push(i)
	}
	s = Stack{}
	rightSmaller := make([]int, len(h))
	for i := len(h) - 1; i >= 0; i-- {
		v := h[i]
		for ; !s.isEmpty() && h[s.peek()] >= v; s.pop() {
		}
		if s.isEmpty() {
			rightSmaller[i] = len(h)
		} else {
			rightSmaller[i] = s.peek()
		}
		s.push(i)
	}

	var overallBest int64
	for i := 0; i < len(h); i++ {
		lm := int64(rightSmaller[i]-leftSmaller[i]-1) * int64(h[i])
		if lm > overallBest {
			overallBest = lm
		}
	}

	return overallBest
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	hTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h []int32

	for i := 0; i < int(n); i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	result := largestRectangle(h)

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
