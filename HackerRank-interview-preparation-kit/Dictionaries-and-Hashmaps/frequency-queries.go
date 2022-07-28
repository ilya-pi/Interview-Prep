package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	// I don't see the catch, seems like a map of slices?
	// Will try brute force approach
	var r []int32
	data := map[int32]int32{}
	freqs := map[int32]int32{}
	for _, q := range queries {
		switch q[0] {
		case 1: // insert
			if freq, ok := freqs[data[q[1]]]; ok && freq > 0 {
				freqs[data[q[1]]]--
			}
			data[q[1]]++
			freqs[data[q[1]]]++
		case 2: // remove occurance
			if v, ok := data[q[1]]; ok && v > 0 {
				freqs[data[q[1]]]--
				data[q[1]]--
				freqs[data[q[1]]]++
			}
		case 3: // check if integer with frequency is present
			if freq, ok := freqs[q[1]]; ok && freq > 0 {
				r = append(r, 1)
			} else {
				r = append(r, 0)
			}
		}
	}
	return r
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

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
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
