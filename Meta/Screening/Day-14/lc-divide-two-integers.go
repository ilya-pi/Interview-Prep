import "math"

func divide(dividend int, divisor int) int {
	/*
	   if divident > 0 {
	   - divisor until divident < divisor and add ans each time
	   }
	   if either < 0 {
	   make them positive and add - sign
	   }

	*/

	var negative bool
	if dividend < 0 {
		dividend = -dividend
		negative = !negative
	}
	if divisor < 0 {
		divisor = -divisor
		negative = !negative
	}

	var ans int
	if divisor == 1 {
		ans = dividend
	} else {
		for dividend >= divisor {
			dividend -= divisor
			ans++
		}
	}
	if negative {
		return -ans
	} else {
		if ans > int(math.MaxInt32) {
			return int(math.MaxInt32)
		}
		return ans
	}

}
