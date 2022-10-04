func longestValidParentheses(s string) int {
	/*
	   Brute force:
	   The longest valid parenthesis sequence has to begin somewhere, hence for every rune in the string we can check what is the longest sequence that is valid
	   O(n^2)
	*/

	if len(s) == 0 || len(s) == 1 {
		return 0
	}

	runes := []rune(s)

	// ")()())"

	bestParenthesisSequence := func(src []rune, ind int) int {
		var stack []rune
		var ans int
		var popped int
		for i := ind; i < len(src); i++ {
			//fmt.Printf("Stack: %s, i %d\n", string(stack), i)
			if src[i] == '(' {
				ans++
				stack = append(stack, '(')
			} else if src[i] == ')' {
				if len(stack) > 0 && stack[len(stack)-1] == '(' {
					ans++
					stack = stack[:len(stack)-1]
					if len(stack) == 0 {
						popped = 0
					} else {
						popped += 2
					}
				} else {
					return ans
				}
			}
		}
		return ans - popped - len(stack)
	}

	var ans int
	for i, _ := range runes {
		option := bestParenthesisSequence(runes, i)
		//fmt.Printf("%d: %d\n", i, option)
		if option > ans {
			ans = option
		}
	}
	return ans

}
