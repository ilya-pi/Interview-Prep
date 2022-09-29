import "strings"

func isPalindrome(s string) bool {
	/*
	   Approach:

	   1/ lowercase string
	   2/ go with a pointer from left and a pointer from right, skipping non-alphanumeric runes
	   3/ if the don't == -> return false
	   4/ while left < right
	   5/ return true at the loops end

	   // Testing 1/3
	   A man, a plan, a canal: Panama
	   a man, a plan, a canal: panama

	   left == 0
	   right == 29

	   left == m
	   right == m

	*/
	s = strings.ToLower(s)

	isAlphanumeric := func(r rune) bool {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
			return true
		}
		return false
	}

	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		for ; left < len(runes) && !isAlphanumeric(runes[left]); left++ {
		}
		for ; right >= 0 && !isAlphanumeric(runes[right]); right-- {
		}
		if left < len(runes) && right >= 0 && runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true

}
