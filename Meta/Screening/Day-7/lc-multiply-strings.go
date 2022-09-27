func multiply(num1 string, num2 string) string {
	/*

	    Basically implement multiplication element by element

	    Multiplication

	    13238
	      120
	    -----
	    00000+
	   26476
	  13238
	  =======
	  1588560
	  1588560

	  Thus.

	  For every rune in shorter string:
	  Multiply and add to int array
	  Then sum all arrays with overflow

	  1. Find shorter string
	  2. Have an acc of slices to acc numbers to sum
	  3. Multiply values and record in matching arrays
	  4. Sum up arrays and respect overflow

	*/

	numToRune := func(n int) rune {
		return rune('0' + n)
	}

	runeToNum := func(r rune) int {
		return int(r - '0')
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	// num1 is longer now

	var acc [][]int
	rnum1 := []rune(num1)
	rnum2 := []rune(num2)
	var mlength int
	// for all runes of shorter string
	for i := len(rnum1) - 1; i >= 0; i-- {
		v := rnum1[i]
		shift := len(rnum1) - 1 - i
		arr := make([]int, shift) // fill first positions with 0, we will append to it only
		a := runeToNum(v)
		// now we need to run the multiplication of a on all b's in num1
		var overflow int
		for j := len(rnum2) - 1; j >= 0; j-- {
			v2 := rnum2[j]
			b := runeToNum(v2)
			overflow += a * b
			arr = append(arr, overflow%10)
			overflow /= 10
		}
		if overflow > 0 { // can't be > 10
			arr = append(arr, overflow)
		}
		if len(arr) > mlength {
			mlength = len(arr)
		}

		acc = append(acc, arr)
	}

	// now we need to sum all acc's respecting overflow
	var sum []int
	var overflow int
	for i := 0; i < mlength; i++ {
		for _, num := range acc {
			if i < len(num) {
				overflow += num[i]
			}
		}
		// all nums are summed
		sum = append(sum, overflow%10)
		overflow /= 10
	}
	for overflow > 0 {
		sum = append(sum, overflow%10)
		overflow /= 10
	}

	// reverse and convert to string
	var ans []rune
	// skip first 0-roes
	i := len(sum) - 1
	for ; i >= 0 && sum[i] == 0; i-- {
	}
	if i < 0 {
		return "0"
	}
	// append the rest
	for ; i >= 0; i-- {
		ans = append(ans, numToRune(sum[i]))
	}
	return string(ans)
}
