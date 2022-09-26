func lengthOfLongestSubstring(s string) int {
	/*
	   Approach:

	   We will iterate through the string and at each position assess if we can use this string as the longest one without repeating characters, if no -> backtrack that element, backtrack length and continue

	   We will need:

	   1. current max length ("ans")
	   2. set of see characters so far
	   3. running length


	*/
	var ans int
	seen := make(map[rune]int)
	var runningMax int
	var start int
	for i, v := range []rune(s) {
		// have we seen it before?
		if pos, saw := seen[v]; saw && pos >= start {
			// backtrack
			runningMax = i - pos
			start = pos
			seen[v] = i
		} else {
			seen[v] = i
			runningMax++
		}
		if runningMax > ans {
			ans = runningMax
		}
	}
	return ans
}
