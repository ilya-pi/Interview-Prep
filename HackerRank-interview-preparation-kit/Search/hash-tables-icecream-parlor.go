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
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) {
	/*

	   It is somewhat similar to pairing the elements when we need them

	   We'll use a map to see if there is a matching second flavour that sums up to the price

	*/

	costs := map[int32]int{}
	for i, v := range cost {
		costs[v] = i // Technically, this would overwrite the existing costs if theyu are the sam,e but it should be okay
	}

	for i, v := range cost {
		// Potentially might be a problem that we overwrite flavours costs,
		// though shouldn't be as we overwrite with the later index in the
		// original cost array
		if j, ok := costs[money-v]; ok && i != j {
			fmt.Printf("%d %d\n", i+1, j+1)
			return
		}
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
