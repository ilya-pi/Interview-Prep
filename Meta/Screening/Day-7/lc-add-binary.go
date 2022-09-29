func addBinary(a string, b string) string {
	/*
	   Approach:

	   Iterarte over both values and perform addition in line

	   So we will need a pointer to the current position, we'll keep overflow value and populate result

	   In the end we need ot revert result and convert it to string
	*/

	// find longer string
	if len(a) < len(b) {
		a, b = b, a
	}
	// a is longer
	ra, rb := []rune(a), []rune(b)
	var ans []rune
	var overflow int
	for i := len(ra) - 1; i >= 0; i-- {
		overflow += int(ra[i] - '0')
		// map i to index in rb
		ind := len(rb) - 1 - (len(ra) - 1 - i)
		if ind >= 0 {
			overflow += int(rb[ind] - '0')
		}
		ans = append(ans, '0'+rune(overflow%2))
		overflow /= 2
	}
	// last overflow
	if overflow > 0 {
		ans = append(ans, '0'+rune(overflow))
	}
	// reverse result
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
