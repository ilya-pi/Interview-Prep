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
 * Complete the 'stepPerms' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

const modul = 10000000007

func stepPerms(n int64) int64 {
	// Edge cases
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}

	// Filling up the recursion call stack
	memo := make([]int64, n+1)
	memo[1] = 1
	memo[2] = 2
	memo[3] = 4
	for i := int64(4); i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2] + memo[i-3]
	}

	return memo[n] % modul
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := nTemp

		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
	}

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
