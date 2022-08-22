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
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */

type tracker [201]int32

func (e *tracker) spend(v int32) {
	if v <= 200 {
		(*e)[v]++
	}
}

func (e *tracker) unSpend(v int32) {
	if (*e)[v] > 0 {
		(*e)[v]--
	}
}

// This was painful to write
func (e *tracker) medianX2(d int32) int32 {
	if d%2 == 1 {
		// We need to find d / 2 + 1 - the element
		i := 0
		for k := int32(0); k < d/2+1; k, i = k+(*e)[i], i+1 {
		}
		i--
		return int32(2 * i)
	} else {
		// We need to find d / 2 and d / 2 + 1 elements
		i := 0
		k := int32(0)
		for ; k < d/2; k, i = k+(*e)[i], i+1 {
		}
		i--
		//fmt.Printf("i is %v\n", i)
		//fmt.Printf("e is %v\n", e)
		// now k >= d / 2
		if k == d/2 {
			j := i + 1
			for ; (*e)[j] == 0; j++ {
			}
			return int32(i + j)
		} else {
			return int32(2 * i)
		}
	}
}

func activityNotifications(expenditure []int32, d int32) int32 {
	e := &tracker{}
	for i := 0; i < int(d); i++ {
		e.spend(expenditure[i])
	}

	var r int32
	for i := int(d); i < len(expenditure); i++ {
		if e.medianX2(d) <= expenditure[i] {
			r++
		}
		e.unSpend(expenditure[i-int(d)])
		e.spend(expenditure[i])
	}
	return r
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
