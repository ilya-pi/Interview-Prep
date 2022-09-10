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
 * Complete the 'poisonousPlants' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func poisonousPlants(p []int32) int32 {
	/*
	   Naive approach:

	   Kill plants day by day, can't be worse then O(n^2)
	*/
	/*
	   var days int32
	   killed := -1
	   for ;killed != 0; {
	       killed = 0
	       toKill := map[int]bool{}
	       for i := 1; i < len(p); i++ {
	           if p[i] > p[i-1] {
	               toKill[i] = true
	               killed++
	           }
	       }
	       newP := make([]int32, len(p) - killed)
	       for k, i := 0, 0; i < len(p); i++ {
	           if ok, _ := toKill[i]; !ok {
	               newP[k] = p[i]
	               k++
	           }
	       }
	       p = newP
	       days++
	   }
	   return days - 1
	*/
	/* Approach #2:

	   a1 a2 a3 a4 a4

	   We will go from the back and add elements to stack.

	   If we see an element that is bigger then the current stack head,
	   then we add it, if it is smaller then we continue killing the head
	   until it is bigger, while counting amount of "kills". Then we max
	   amount of kills with the current max

	*/
	/*
		var stack []int32
		var maxKills int32
		for i := len(p) - 1; i >= 0; i-- {
			if len(stack) == 0 {
				stack = append(stack, p[i])
				continue
			}
			// peek
			head := stack[len(stack)-1]
			if p[i] >= head {
				stack = append(stack, p[i])
				continue
			}
			// p[i] is less then head, so we start "killing" heads
			var kills int32
			for p[i] < head {
				fmt.Printf("%v killed %v\n", p[i], head)
				kills++
				// pop
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					break
				}
				// peek
				head = stack[len(stack)-1]
			}
			fmt.Printf("%v killed %v overall\n", p[i], kills)
			if kills > maxKills {
				maxKills = kills
			}
			stack = append(stack, p[i])
		}
		return maxKills
	*/
	/*
		Correction on approach#2:
		At times it will take longer to kill all the plants, as there might be
		overlapping ascending sequences that cut in the middle, then it will take
		at the max of "killing" days for continue the killing if such a plant was met,
		this we need to record the amount of killed plants so far and do corrections if
		we create new killing ascending sequence.
	*/
	type pair struct {
		killed int32
		v      int32
	}
	var stack []pair
	var maxKills int32
	for i := len(p) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, pair{v: p[i]})
			continue
		}
		// peek
		head := stack[len(stack)-1]
		if p[i] >= head.v {
			stack = append(stack, pair{v: p[i]})
			continue
		}
		// p[i] is less then head, so we start "killing" heads
		var kills int32
		for p[i] < head.v {
			kills++
			if kills < head.killed {
				kills = head.killed
			}
			// pop
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			// peek
			head = stack[len(stack)-1]
		}
		if kills > maxKills {
			maxKills = kills
		}
		stack = append(stack, pair{v: p[i], killed: kills})
	}
	return maxKills
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}

	result := poisonousPlants(p)

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
