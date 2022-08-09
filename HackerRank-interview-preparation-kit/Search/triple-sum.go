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

// Complete the triplets function below.
func triplets(a []int32, b []int32, c []int32) int64 {
	/*
	   a - p [1, 5, 2, 6]
	   b - q [4, 3]
	   c - r [5, 1, 6, 2]
	   p <= q >= r

	   4 -> [5, 6] + 4 + [5, 6] {5, 4, 5} {5, 4 , 6} {6, 4, 5} {6, 4, 6}

	   Same for 5

	   So we iterate over all the elements in the mid array and find the amount of elements less then that element in a and b, and multiplication of those would give us the amount of pairs for that element

	*/

	filter := func(arr []int32) []int32 {
		elems := map[int32]bool{}
		for _, v := range arr {
			elems[v] = true
		}
		r := make([]int32, len(elems))
		i := 0
		for k, _ := range elems {
			r[i] = k
			i++
		}
		return r
	}

	a = filter(a)
	b = filter(b)
	c = filter(c)

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	// Will work only on sorted arrays
	/* logn search in a and c
	   countCandidates := func(arr []int32, v int32) int64 {
	       low := 0
	       high := len(arr) - 1
	       for low != high {
	           j := (low + high) / 2
	           if arr[j] <= v && arr[j+1]>v {
	               return int64(j + 1)
	           } else if arr[j] <= v {
	               low = j + 1
	           } else {
	               high = j - 1
	           }
	       }
	       if arr[low] <= v {
	           return int64(low + 1)
	       } else {
	           return 0
	       }
	   }
	*/

	var triplets int64
	var ai int
	var ci int
	for _, v := range b {
		for ; ai < len(a) && a[ai] <= v; ai++ {
		}
		for ; ci < len(c) && c[ci] <= v; ci++ {
		}
		triplets += int64(ai) * int64(ci)
	}

	// O(aloga + blogb + clogc)

	return triplets
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
	checkError(err)
	lena := int32(lenaTemp)

	lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
	checkError(err)
	lenb := int32(lenbTemp)

	lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
	checkError(err)
	lenc := int32(lencTemp)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int32

	for i := 0; i < int(lena); i++ {
		arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
		checkError(err)
		arraItem := int32(arraItemTemp)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int32

	for i := 0; i < int(lenb); i++ {
		arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
		checkError(err)
		arrbItem := int32(arrbItemTemp)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int32

	for i := 0; i < int(lenc); i++ {
		arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
		checkError(err)
		arrcItem := int32(arrcItemTemp)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets(arra, arrb, arrc)

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
