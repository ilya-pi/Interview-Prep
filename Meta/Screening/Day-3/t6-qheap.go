package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type heap []int64

func (h *heap) add(v int64) {
	// panic on *h == nil, technically returning an error could've been better
	*h = append(*h, v)
	h.sift(len(*h) - 1)
}

func (h heap) min() int64 {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h *heap) delete(v int64) {
	//fmt.Printf("Delete %v from %v\n", v, *h)
	var ind int
	for i, v1 := range *h {
		if v == v1 {
			ind = i
			break
		}
	}
	(*h)[ind], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[ind]
	*h = (*h)[:len(*h)-1]
	if ind < len(*h) {
		h.sift(ind)
	}
}

func (hh *heap) sift(ind int) {
	h := *hh
	parent := (ind - 1) / 2
	//fmt.Printf("ind %d parent %d, %v\n", ind, parent, *h)
	if parent >= 0 && h[parent] > h[ind] {
		h[parent], h[ind] = h[ind], h[parent]
		h.sift(parent)
		return
	}
	ch1 := ind*2 + 1
	ch2 := ind*2 + 2
	if ch1 < len(h) && h[ch1] < h[ind] &&
		(ch2 < len(h) && h[ch1] < h[ch2]) {
		h[ch1], h[ind] = h[ind], h[ch1]
		h.sift(ind)
		return
	}
	if ch2 < len(h) && h[ch2] < h[ind] {
		h[ch2], h[ind] = h[ind], h[ch2]
		h.sift(ind)
		return
	}
	*hh = h
}

func process(arr [][]int64, writer *bufio.Writer) {
	var h heap
	for _, v := range arr {
		switch v[0] {
		case 1:
			h.add(v[1])
		case 2:
			h.delete(v[1])
		case 3:
			fmt.Fprintf(writer, "%d\n", h.min())
		}
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
	q := int64(qTemp)

	var arr [][]int64

	for tItr := 0; tItr < int(q); tItr++ {
		s := readLine(reader)

		e := strings.Split(s, " ")
		var ln []int64
		for _, v := range e {
			vi, err := strconv.ParseInt(v, 10, 64)
			checkError(err)
			ln = append(ln, int64(vi))
		}
		arr = append(arr, ln)
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
