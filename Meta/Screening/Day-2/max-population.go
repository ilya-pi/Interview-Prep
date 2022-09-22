package main

import "fmt"

type Person struct {
	born int
	died int
}

func maxPopulation(people []Person) int {
	/*
	   Approach:

	   Keep a fixed array for the years in reasonable interval

	   Map born/died years as increment/decrement on that array

	   Count max year
	*/

	years := make([]int, 3000)

	for _, p := range people {
		years[p.born]++
		years[p.died]--
	}

	var max int
	var maxYear int
	var population int
	for year, delta := range years {
		population += delta
		if max < population {
			max = population
			maxYear = year
		}
	}
	return maxYear
}

// 2000 * 8 b  = 16000 b -> 16kb memmory

func main() {
	people := []Person{{1980, 2034}, {1950, 1960}, {1900, 2000}, {1945, 2020}}
	max := maxPopulation(people)
	fmt.Printf("Result is %d\n", max)
}
