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

This still doesn't pass, as the tree becomes unbalanced :-/ But there are no balanced trees in Go std
And I am a bit too lazy to implement AVL or Red Black trees

*/

/*
 * Complete the 'maximumSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. LONG_INTEGER_ARRAY a
 *  2. LONG_INTEGER m
 */

type BST struct {
	v           int64
	left, right *BST
}

func (b *BST) add(v int64) *BST {
	if b == nil {
		n := &BST{
			v: v,
		}
		return n
	}
	if v < b.v {
		if b.left == nil {
			n := &BST{v: v}
			b.left = n
		} else {
			b.left.add(v)
		}
	} else if v > b.v {
		if b.right == nil {
			n := &BST{v: v}
			b.right = n
		} else {
			b.right.add(v)
		}
	} // else it is ==, then ignore

	return b
}

func (b *BST) higher(v int64, higher int64) int64 {
	if b == nil {
		return higher
	}
	if b.v == v {
		return v
	}
	if b.v > v {
		if b.v < higher || higher == 0 {
			higher = b.v
		}
		return b.left.higher(v, higher)
	} else {
		return b.right.higher(v, higher)
	}
}

func maximumSum(a []int64, m int64) int64 {
	/*
	   Brute force! :-D
	*/
	/*
	   sums := make([]int64, len(a))
	   sums[0] = a[0] % m
	   for i := 1; i < len(a); i++ {
	       sums[i] = (sums[i-1] + a[i]) % m
	   }

	   var max int64
	   lookup := map[int64]bool{} // existing sum modulo values
	   // todo can be merged up
	   for i := 0; i < len(a); i++ {
	       v := sums[i]
	       if v == m - 1 {
	           return m - 1 // cannot be better!
	       }
	       //Trying to find best match
	       for j := v + 1; j < m && max < (m + v - j)%m; j++ {
	           if _, ok := lookup[j]; ok {
	               nMax := (m + v - j)%m
	               if nMax > max {
	                   max = nMax
	               }
	           }
	       }
	       lookup[v] = true
	   }
	   return max
	*/

	/*
	   Approach with a structure that tries to give us higher bound in logn
	*/
	var bst *BST
	sums := make([]int64, len(a))
	v0 := a[0] % m
	bst = bst.add(v0)
	sums[0] = v0
	max := v0
	for i := 1; i < len(a); i++ {
		v := (sums[i-1] + a[i]) % m
		// Try to find best match
		// We want previous sum to end up at v+1, so that it -1 the m and we get max sum, otherwise it will be lower, so (v+1), (v+2), (v+3),... will be good candidates for getting m-1, m-2, m-3,... as max
		higher := bst.higher(v+1, 0)
		var sum int64
		if higher > 0 {
			sum = (m + v - higher) % m
		} else {
			sum = v
		}
		if sum == m-1 {
			return sum
		}
		if sum > max {
			max = sum
		}

		// Add new values to supporting structures
		sums[i] = v
		bst = bst.add(v)
	}
	return max
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

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		m, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)

		aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var a []int64

		for i := 0; i < int(n); i++ {
			aItem, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			a = append(a, aItem)
		}

		result := maximumSum(a, m)

		fmt.Fprintf(writer, "%d\n", result)
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
