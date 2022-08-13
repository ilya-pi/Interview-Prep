package main

import "fmt"

const N = 4
const M = 3

type matrix [N][M]int

func zeroMatrix(m matrix) matrix {
	cols := map[int]bool{}
	rows := map[int]bool{}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if m[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i, _ := range rows {
		for j := 0; j < M; j++ {
			m[i][j] = 0
		}
	}

	for j, _ := range cols {
		for i := 0; i < N; i++ {
			m[i][j] = 0
		}
	}

	// Can also use the fist elements in the matrix as "0" markers, as we already iterated through those elements and then use them when zeroing the matrix

	return m
}

func main() {
	m := matrix{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Printf("Zeroing %v is\n%v\n", m, zeroMatrix(m))
}
