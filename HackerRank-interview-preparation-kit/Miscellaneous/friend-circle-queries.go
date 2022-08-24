package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maxCircle function below.
func maxCircle2(queries [][]int32) []int32 {
	/*
	   Essentially we are building a graph and adding new edges at each step.
	   We can run a BFS from the edge and count it's group size, if it bigger then the previous - then it is the new biggest group and we should see that as new max size frined group
	*/

	adj := map[int32]map[int32]bool{} // Adjacency map

	bfsCount := func(i int32) int32 {
		visited := map[int32]bool{}
		q := []int32{}
		q = append(q, i)
		var cnt int32
		for len(q) != 0 {
			v := q[0]
			q = q[1:]
			if yes, ok := visited[v]; yes && ok {
				continue
			}
			cnt++
			visited[v] = true
			for v2, _ := range adj[v] {
				if _, ok := visited[v2]; !ok {
					q = append(q, v2)
				}
			}
		}
		return cnt
	}

	/*
	   On each step we add adjacency to both nodes and perform BFS from first node to find new max group size
	*/

	var maxGroupSize int32
	res := make([]int32, len(queries))
	for i, v := range queries {
		v1 := v[0]
		v2 := v[1]
		if _, ok := adj[v1]; !ok {
			adj[v1] = map[int32]bool{}
		}
		if _, ok := adj[v2]; !ok {
			adj[v2] = map[int32]bool{}
		}
		adj[v1][v2] = true
		adj[v2][v1] = true
		newMax := bfsCount(v1)
		if newMax >= maxGroupSize {
			maxGroupSize = newMax
		}
		res[i] = maxGroupSize
	}
	return res
}

// Complete the maxCircle function below.
func maxCircle3(queries [][]int32) []int32 {
	/*

	   Situations:

	   a-b is added:

	   a is a part of a group already & b part of teh group -> merge groups a and b (merge links)
	   a is not part of a group, b is -> add a to b's group (update links)
	   b is not part of a group, a is -> add b to a's group (update links)

	   On each step maintain maxSize and print if it bigger

	*/

	groups := map[int32]*map[int32]bool{}
	var max int32
	res := make([]int32, len(queries))
	for i, v := range queries {
		a, b := v[0], v[1]
		gA, aOk := groups[a]
		gB, bOk := groups[b]

		switch {
		case !aOk && !bOk:
			gab := map[int32]bool{a: true, b: true}
			groups[a] = &gab
			groups[b] = &gab
		case !aOk && bOk:
			(*gB)[a] = true
			groups[a] = gB
		case aOk && !bOk:
			(*gA)[b] = true
			groups[b] = gA
		case aOk && bOk:
			//merge
			for k, v := range *gB {
				// Merge all the elements
				(*gA)[k] = v
				// Update all the links in the merged group
				groups[k] = gA
			}
			groups[b] = gA
		}

		newMax := int32(len((*groups[b])))
		if newMax > max {
			max = newMax
		}
		res[i] = max
	}
	return res
}

type DSU struct {
	sizes   map[int32]int32
	parents map[int32]int32
}

func newDSU() *DSU {
	r := DSU{
		sizes:   map[int32]int32{},
		parents: map[int32]int32{},
	}
	return &r
}

func (d *DSU) add(v int32) int32 {
	d.parents[v] = v
	d.sizes[v] = 1
	return d.parents[v]
}

func (d *DSU) parent(v int32) int32 {
	p := d.parents[v]
	if p == v || p == 0 {
		return p
	} else {
		d.parents[v] = d.parent(p)
		return d.parents[v]
	}
}

func (d *DSU) union(a int32, b int32) {
	pA, pB := d.parent(a), d.parent(b)
	if pA == pB {
		return
	}
	if d.sizes[pA] < d.sizes[pB] {
		pA, pB, a, b = pB, pA, b, a
	}
	d.parents[pB] = pA
	d.sizes[pA] += d.sizes[pB]
}

func (d *DSU) size(v int32) int32 {
	return d.sizes[d.parent(v)]
}

func maxCircle(queries [][]int32) []int32 {

	/*
	   We will use Disjoint Set Union data structure that I had to read about and it performs most required operations in ammrotized O(1)
	*/

	var max int32
	dsu := newDSU()
	res := make([]int32, len(queries))
	for i, v := range queries {
		a, b := v[0], v[1]
		aP, bP := dsu.parent(a), dsu.parent(b)
		if aP == 0 {
			aP = dsu.add(a)
		}
		if bP == 0 {
			bP = dsu.add(b)
		}
		if aP != bP {
			dsu.union(a, b)
		}

		newSetSize := dsu.size(a)
		if newSetSize > max {
			max = newSetSize
		}
		res[i] = max
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := maxCircle(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
