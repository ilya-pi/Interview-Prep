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

// Complete the getMinimumCost function below.
func getMinimumCost(k int32, c []int32) int32 {
	// Have an array of length K to track the amount of flowers everybody bought and continuously pick the next friend that got least flowers
	friends := make([]int32, k)

	sort.Slice(c, func(i, j int) bool { return c[i] > c[j] })

	var price int32
	for i := 0; i < len(c); i++ {
		fI := int32(i) % k
		price += (friends[fI] + 1) * c[i]
		friends[fI]++
	}
	return price
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	minimumCost := getMinimumCost(k, c)

	fmt.Fprintf(writer, "%d\n", minimumCost)

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
