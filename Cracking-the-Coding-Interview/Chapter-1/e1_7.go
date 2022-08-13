package main

import "fmt"

const N = 4

type matrix [N][N]int

func rotate(image matrix) matrix {
	// Brute force approach
	/*
		n := matrix{}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				n[j][N-1-i] = image[i][j]
			}
		}
		return n
	*/

	/*
	   In place approach:

	   1 2 3 4
	   1 2 3 4
	   1 2 3 4
	   1 2 3 4

	   ->
	       k
	   l 1 x x 1
	     2 x x 2
	     3 x x 3
	     4 4 4 4

	   n = len(matrix[]) -1

	   for l := 0 .. < len(marix[]) / 2
	     for  k:= l .. n - l

	   l, k         ->  k, n - l
	   k, n - l     ->  n - l, n - k
	   n - l, n - k -> n - k, l
	   n - k, l     -> l, k
	*/

	fmt.Printf("len is %d\n", len(image))

	for l := 0; l < N/2; l++ {
		for k := l; k < N-1-l; k++ {
			n := N - 1
			//image[k][N-1-l], image[N-1-l][N-1-k], image[N-1-k][l], image[l][k] = image[l][k], image[k][N-1-l], image[N-1-l][N-1-k], image[N-1-k][l]
			image[k][n-l], image[n-l][n-k], image[n-k][l], image[l][k] = image[l][k], image[k][n-l], image[n-l][n-k], image[n-k][l]
		}
	}
	return image
}

func main() {
	//m := matrix{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	m := matrix{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	fmt.Printf("%v \nrotated is \n%v\n", m, rotate(m))
}
