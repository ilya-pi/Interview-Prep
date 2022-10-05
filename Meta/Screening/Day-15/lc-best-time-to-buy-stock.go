import "math"

func maxProfit(prices []int) int {
	/*
	   The best profit seem to be if you bought at the lowest price and sold at the highest after

	   So on each day we "try" to sell against the current minimum, if that yields more profit -> go for it
	   At the same time each day presents new opportunity to buy at the minimum, so we try that too
	   So we need:

	   min
	   bestProfit

	   And then scan array left to right

	   7 1 5 3 6 4

	   1
	   5
	*/

	min := math.MaxInt
	var bestProfit int
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > bestProfit {
			bestProfit = v - min
		}
	}
	return bestProfit

}
