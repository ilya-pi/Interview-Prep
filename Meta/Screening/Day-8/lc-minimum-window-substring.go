func minWindow(s string, t string) string {
	/*
	 GVADOBECODEBANC
	 ABC
	 BANC
	 GVADOBECODEBANC

	 ---DOBECODEBA--

	 ---DOBECODEBA--

	 Balance map:
	 A - 0
	 B - -1
	 C - 0

	 canSkip(checks that rune's value is in required map is balance greater or equal to required)
	 hasAll(checks that map is 0 or <0)

	 once cant' skip -> record the shortest result so far

	 and continue step with left pointer while maintaining the balance

	 A - 1
	 B - 1
	 C - 1

	 ----

	 We will go from left to right with two pointers, left and right.

	 While going with right we will add runes to map count of runes checking at every step that we cover other string (map key-value comparison)

	 Once cover, we will move left pointer untill we still cover (canSkip method that checks balance of the desired rune)

	 Once can't skip -> record this result as option if it is shorter then result so far (or if so far was empty) and make step with left pointer respecting balances, then the right pointer cycle continues.

	 We continue till right pointer is at the end and then we do last cycle outside of loop

	*/

	// Build required map
	required := make(map[rune]int)
	for _, r := range t {
		required[r]++
	}

	balance := make(map[rune]int)

	covers := func() bool {
		for r, count := range required {
			if c2, ok := balance[r]; !ok || c2 < count {
				return false
			}
		}
		return true
	}

	canSkip := func(r rune) bool {
		if requiredCount, ok := required[r]; ok {
			if currentCount, ok1 := balance[r]; ok1 && currentCount > requiredCount {
				return true
			} else {
				return false
			}
		}
		return true // not a required rune
	}

	left, right := 0, 0
	rs := []rune(s)
	var ans []rune
	for {
		if covers() {
			for canSkip(rs[left]) {
				balance[rs[left]]--
				left++
			}
			// we have a local minimum of covering string between left and right
			if len(ans) == 0 || len(ans) > right-left {
				ans = rs[left:right]
			}
			// make a step with a left that will turn covers() into false
			balance[rs[left]]--
			left++
		}
		if right == len(rs) {
			// reached the end of the string and did all the checks
			break
		}
		balance[rs[right]]++
		right++
	}
	return string(ans)
}
