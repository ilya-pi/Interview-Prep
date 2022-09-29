func validPalindrome(s string) bool {
	/*
	   Approach:

	   Check from left and from right that it can be a palindrome, if found diff in runes -> check for both options to be a palindrome

	*/

	var palindrome func(string, bool) bool
	palindrome = func(s string, token bool) bool {
		runes := []rune(s)
		for left, right := 0, len(runes)-1; left < right; {
			if runes[left] != runes[right] {
				if !token {
					return false
				}
				token = false
				return palindrome(s[left:right], token) || palindrome(s[left+1:right+1], token)
			} else {
				left++
				right--
			}
		}
		return true
	}
	return palindrome(s, true)
}
