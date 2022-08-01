package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func activityNotifications(expenditure []int32, d int32) int32 {
	if int32(len(expenditure)) < d {
		return 0
	}

	spent := make([]int32, d)
	copy(spent, expenditure[:d])
	sort.Slice(spent, func(i, j int) bool { return spent[i] < spent[j] })

	doubleMean := func(arr []int32) int32 {
		if len(arr)%2 == 1 {
			return 2 * arr[len(arr)/2]
		} else {
			return arr[len(arr)/2] + arr[len(arr)/2-1]
		}
	}

	var bs func([]int32, int32) int32
	bs = func(arr []int32, v int32) int32 {
		if len(arr) == 1 {
			if arr[0] == v {
				return 0
			} else {
				// Factually can't happen
				return -1
			}
		}

		if v < arr[len(arr)/2] {
			return bs(arr[:len(arr)/2], v)
		} else {
			return int32(len(arr)/2) + bs(arr[len(arr)/2:], v)
		}
	}

	var res int32
	for i := d; i < int32(len(expenditure)); i++ {
		fmt.Sprintf("%d\n", i)
		if expenditure[i] >= doubleMean(spent) {
			res++
		}
		// bs element, change for new one and bubble it in place
		j := bs(spent, expenditure[i-d])
		// fmt.Printf("j is %d\n", j)
		// replace that element
		spent[j] = expenditure[i]
		// bubble it now
	bubble_loop:
		for {
			switch {
			case j == 0 && spent[j] <= spent[j+1]:
				break bubble_loop
			case j == int32(len(spent)-1) && spent[j-1] <= spent[j]:
				break bubble_loop
			case j < int32(len(spent))-1 && spent[j] > spent[j+1]:
				spent[j], spent[j+1] = spent[j+1], spent[j]
				j++
			case j > 0 && spent[j] < spent[j-1]:
				spent[j], spent[j-1] = spent[j-1], spent[j]
				j--
			default:
				break bubble_loop
			}
		}
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	stdout := os.Stdout

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
