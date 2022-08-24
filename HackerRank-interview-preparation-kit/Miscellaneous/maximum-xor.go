package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Trie struct {
	nodes [2]*Trie
}

func (t *Trie) add(v int32) {
	current := t
	// We should add elements from the beginning as that is a higher bit
	for i := 31; i >= 0; i-- {
		mask := int32(1) << i
		bit := (mask & v) >> i
		if current.nodes[bit] == nil {
			current.nodes[bit] = &Trie{}
		}
		current = current.nodes[bit]
	}
}

func (t *Trie) maxXor(v int32) int32 {
	current := t
	var res int32
	for i := 31; i >= 0; i-- {
		mask := int32(1) << i
		bit := (v & mask) >> i
		// Can we go to an opposite bit to get higher xor?
		if current.nodes[1-bit] != nil {
			//yes
			res += 1 << i
			current = current.nodes[1-bit]
		} else {
			//no
			current = current.nodes[bit]
		}
	}
	return res
}

// Complete the maxXor function below.
func maxXor(arr []int32, queries []int32) []int32 {
	/*
	   Brute force calculation
	   res := make([]int32, len(queries))
	   for i, q := range queries {
	       max := arr[0]^q
	       for i := 1; i < len(arr); i++ {
	           v := arr[i]
	           r := q^v
	           if r > max {
	               max = r
	           }
	       }
	       res[i] = max
	   }
	   return res
	*/

	/*
	   My dumbzy dumb - I need to use a trie of the inversed numbers from 32 to pick the "best" highest bits
	*/

	trie := &Trie{}
	res := make([]int32, len(queries))
	for _, v := range arr {
		trie.add(v)
	}
	for i, q := range queries {
		res[i] = trie.maxXor(q)
	}
	return res
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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries []int32

	for i := 0; i < int(m); i++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := maxXor(arr, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
