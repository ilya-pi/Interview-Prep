package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'commonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func commonChild(s1 string, s2 string) int32 {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	if len(s1) == 1 || len(s2) == 1 {
		if s1[0] == s2[0] {
			return 1
		} else {
			return 0
		}
	}

	common := make([][]int32, len(s1))
	for i := 0; i < len(common); i++ {
		common[i] = make([]int32, len(s2))
	}

	max := func(a int32, b int32) int32 {
		if a > b {
			return a
		} else {
			return b
		}
	}

	if s1[0] == s2[0] {
		common[0][0] = 1
	} else {
		common[0][0] = 0
	}

	for i := 1; i < len(s1); i++ {
		if s2[0] == s1[i] {
			common[i][0] = 1
		} else {
			common[i][0] = common[i-1][0]
		}
	}

	for j := 1; j < len(s2); j++ {
		if s1[0] == s2[j] {
			common[0][j] = 1
		} else {
			common[0][j] = common[0][j-1]
		}
	}

	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if s1[i] == s2[j] {
				common[i][j] = common[i-1][j-1] + 1
				//fmt.Printf("%v", string(s1[i]))
			} else {
				common[i][j] = max(common[i-1][j], common[i][j-1])
			}
		}
	}

	//fmt.Printf("\n\n%v\n\n", common)
	/*
	   [1 0 0]
	   [0 0 1]
	   [0 0 1]
	*/

	return common[len(s1)-1][len(s2)-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

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
