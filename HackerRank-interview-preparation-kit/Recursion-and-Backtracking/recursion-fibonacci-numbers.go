package main

import "fmt"

func fibonacci(n int) int {
	/*
	   	Approach:

	   We will use memo to remember the value
	   of previous fibonacci numbers
	*/
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
