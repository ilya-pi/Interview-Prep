package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumPasses' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. LONG_INTEGER m
 *  2. LONG_INTEGER w
 *  3. LONG_INTEGER p
 *  4. LONG_INTEGER n
 */

func minimumPasses(m int64, w int64, p int64, n int64) int64 {
	/*
	   Approach:

	   Count days one by one while understanding the remaining days. If remaining days is less then the countDays -> return the smaller one
	*/

	var pastDays int64
	var candies int64
	var remainingDays int64
	production := m * w
	if m > math.MaxInt64/w {
		return int64(1) // Handling int64 overflow in a rather crude way
	}

	remainingDays = n / production
	if n%production > 0 {
		remainingDays++
	}
	for pastDays < remainingDays {
		// Produced for 1 day
		pastDays++
		candies += production

		// Is enough to invest?
		if candies < p {
			extraDays := (p - candies) / production
			if (p-candies)%production > 0 {
				extraDays++
			}

			// Work extra days
			pastDays += extraDays
			candies += extraDays * production
		}

		// Investment decisions
		toBuy := candies / p
		candies = candies % p
		if m < w {
			if w-m > toBuy {
				m += toBuy
				toBuy = 0
			} else {
				toBuy -= w - m
				m = w
			}
		} else if m > w {
			if m-w > toBuy {
				w += toBuy
				toBuy = 0
			} else {
				toBuy -= m - w
				w = m
			}
		}
		// Invest the rest
		m += toBuy / 2
		w += toBuy - toBuy/2

		// Production updates
		production = m * w

		// Remaining Days update
		newRemainingDays := pastDays + (n-candies)/production
		if (n-candies)%production > 0 {
			newRemainingDays++
		}
		if newRemainingDays < remainingDays {
			remainingDays = newRemainingDays
		}
	}
	if pastDays > remainingDays {
		return remainingDays
	} else {
		return pastDays
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	m, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)

	w, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)

	p, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)

	n, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)

	result := minimumPasses(m, w, p, n)

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
