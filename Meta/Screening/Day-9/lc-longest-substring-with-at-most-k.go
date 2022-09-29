func lengthOfLongestSubstringKDistinct(s string, k int) int {
	/*

	   ecebadeeea k = 4

	   Go left to right, countin the amount of distinct characters through map

	   If we have reached k or end => try record max length,
	   Once recorded, if not at the end => move left pointer until we are under k in distinct characters, then move right pointer until we are at k characters again and try to record max lenght, repeat

	*/
	if k == 0 {
		return 0
	}

	/*
	   eceba
	   k = 2

	   e:2 c:1 b:1

	   0, 2
	*/

	distinct := make(map[rune]int)
	add := func(r rune) {
		distinct[r]++
	}
	remove := func(r rune) {
		distinct[r]--
		if distinct[r] == 0 {
			delete(distinct, r)
		}
	}
	canAddMore := func(r rune) bool {
		add(r)
		defer remove(r)
		if len(distinct) <= k {
			return true
		}
		return false
	}
	runes := []rune(s)
	var ans int
	left, right := 0, 0
	for right < len(runes) {
		r := runes[right]
		if canAddMore(r) {
			add(r)
			right++
		} else {
			// Try record current length as max
			if right-left > ans {
				ans = right - left
			}
			for !canAddMore(r) {
				remove(runes[left])
				left++
			}
		}
	}
	// try expand left if we could add more
	if canAddMore(runes[len(runes)-1]) {
		for left >= 0 && canAddMore(runes[left]) {
			add(runes[left])
			left--
		}
		left++
	}
	if right-left > ans {
		ans = right - left
	}
	return ans
}
