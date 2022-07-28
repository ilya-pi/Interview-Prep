package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	/*
	   // 0. We take it that all elemetns that mod r or r^2 are not options to the potential sequence
	   // 1. Create two dictionaries of /r and /(r^2)
	   // 2. Go array step by step and see if we can create those pairs by lookup <- let's begin here as I struggle a bit to fully visualize it in my head, but it feels that it should work

	   rs := map[int64][]int{}
	   rrs := map[int64][]int{}
	   for i, v := range arr {
	       rk := v / r
	       if rk != 0 {
	           rs[rk] = append(rs[rk], i)
	           rrk := rk / r
	           if rrk != 0 {
	               rrs[rrk] = append(rrs[rrk], i)
	           }
	       }
	   }

	   var res int64
	   // Amount of solutions for j as middle
	   jLookup := map[int]int64{}
	   // n * n *
	   for i, v := range arr {
	       // Find first j candidate
	       js := rs[v]
	       for _, j := range js {
	           if i < j {
	               if solutionsAmount, ok := jLookup[j]; ok {
	                   res += solutionsAmount
	               } else {
	                   var solAmount int64
	                   ks := rrs[v]
	                   for _, k := range ks {
	                       if j < k {
	                           solAmount++
	                           // Found a valid triplet (i, j, k)
	                       }
	                   }
	                   jLookup[j] = solAmount
	                   res += solAmount
	               }
	           }
	       }
	   }
	   return res
	*/

	// As we walk through the array, track the amount of potential endings this value can be of and increase those
	r2 := map[int64]int64{}
	r3 := map[int64]int64{}
	var res int64
	for _, v := range arr {
		// Current value is an end to a triplet,
		// since it is an end, then i and j are naturally smaller
		if amount, ok := r3[v]; ok {
			res += amount
		}
		// Current value can be a potential continuation for third value for the
		// amount of geometric pairs that already exist there
		r3[v*r] += r2[v]

		// Current value can be a continuation of a geometric pair
		r2[v*r]++
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
