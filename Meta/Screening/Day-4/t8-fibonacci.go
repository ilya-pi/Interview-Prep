package main

import "fmt"

func fibonacci(n int64) int64 {
	memo := make([]int64, n+1)

	var fib func(int64) int64
	fib = func(n int64) int64 {
		if n == 0 || n == 1 {
			return 1
		} // how do we see the first ones?

		if memo[n] != 0 {
			return memo[n]
		}

		f1 := fib(n - 1)
		f2 := fib(n - 2)
		memo[n-1] = f1
		memo[n-2] = f2
		return f1 + f2
	}

	return fib(n)
}

func main() {
	for i := 0; i < 40; i++ {
		fmt.Printf("%d\n", fibonacci(int64(i)))
	}
}
