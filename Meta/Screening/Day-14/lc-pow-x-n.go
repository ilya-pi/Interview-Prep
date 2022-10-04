func myPow(x float64, n int) float64 {
	/*
		    To pow x we need to * n times
		    To make it faster we can go in steps of x^2

		    x^1 = x
		    x^2 = x * x
		    x^4 = x^2 * x^2
		    x^8 = x^4 * x^4
		    ..
		    x^k that k > n and then fill the rest

		Find first bit higher then n:
		for n > 0 {
		bits++
		n >> 1
		}

		1023
		0111111111
		1
		2 1
		4 2
		...
		1023 will have highest degree 9
		1024 10

		0 ->9, and I will have 2^9

		1. Find top bit of n
		2. Populate array till top bit of n
		3. From the end of array, continuosly multiply res with val in array, while reducing n if n == 0 -> break

		We got x^n if n > 0
		if n < 0 0> 1/x^n

		--

	*/

	var negative bool
	if n < 0 {
		negative = true
		n = -n
	}
	if n == 0 {
		return float64(1.0)
	}

	// Find top bit of n
	k := n
	var bits int
	for k > 0 {
		bits++
		k >>= 1
	}

	// Populate look up array
	lookup := make([]float64, bits+1)
	lookup[0] = x // 2^0 == 1 and would give us 1 x
	/*
	   2^0
	   2^1
	   2^3
	   2^4
	   ..
	   2^8 = 256

	*/
	for i := 1; i < len(lookup); i++ {
		lookup[i] = lookup[i-1] * lookup[i-1]
	}
	//fmt.Printf("%v\n", lookup)

	var ans float64
	ans = 1.0
	for k, i := n, len(lookup)-1; k > 0; {
		if k >= (1 << i) {
			k -= 1 << i
			ans *= lookup[i]
		} else {
			i--
		}
	}
	if negative {
		return float64(1.0) / ans
	} else {
		return ans
	}
}
