func numDecodings(s string) int {
	/*
	   'A' -> "1"
	   'B' -> "2"
	       ...
	   'Z' -> "26"

	   We will try to build it up bottoms up, look at how many ways are there to construct string so far.

	   Zo.

	   1231027

	   i=0
	   one character -> one way if it is not 0
	   p[0] = 1

	   i=1
	   if 0 -> 1 (and i-1 == 1 or 2)
	   if not 0 -> 1
	   if <= 6 -> 2

	   i
	   if  0 -> amount to get to p[i-2] (and i-1 == 1 or 2)
	   if not 0 -> then +amount to build p[i-1]
	   if <= 6 -> +amount to get to p[i-2] (and i-1 == 1 or 2)

	   And till the end we go

	*/

	dp := make([]int, len(s))
	// Base cases
	dp[0] = 1
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return dp[0]
	}

	dp[1] = 1
	if s[1] == '0' {
		if s[0] > '2' {
			return 0
		}
	} else if (s[0] == '1' && s[1] >= '1') || (s[0] == '2' && s[1] <= '6') {
		dp[1] = 2
	}

	for i := 2; i < len(dp); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else {
			// Take as one
			dp[i] += dp[i-1]
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				dp[i] += dp[i-2]
			}
		}
	}

	//fmt.Printf("%v\n", dp)
	/*
	    2 6 1 1 0 5 5 9 7 1 756562

	    2 6 1 10  5 5 9 7 1
	    26  1 10  5 5 9 7 1


	   [1 2 2 4 2 2 2 2 2 2 2 2 2 2 2 2]

	*/

	return dp[len(s)-1]
}
