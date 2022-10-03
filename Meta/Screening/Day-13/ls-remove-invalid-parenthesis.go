func removeInvalidParentheses(s string) []string {
	/*

	   ()())()

	   (())) ) (

	   1. Count the number of parenthesis to be deleted
	   2. In a recursion take only if the stack is valid and can still take

	   To count:

	   stack
	   push '(', pop on ')' if have matching. If cannot -> +1 errorCount
	   at the end pop stack untill it is empty and count errorCount

	   To recurse:

	   rec(src []rune, stack *[]rune, toTake int, prefix *[]rune)
	     if toTake == 0 and stack is empty -> add result
	       return

	     if r == '(' and toTake>1 -> yes
	     if r == ')' and stack.peek == '(', pop stack and -> yes
	     add to prefix
	     rec(src, stack, toTake -1, prefix)
	     remove from prefix
	     if r == '(' -> pop stack
	     if r == ')' -> push stack '('
	     rec(src, stack, toTake -1, prefix)

	*/

	if len(s) == 0 {
		return []string{}
	}

	runes := []rune(s)

	// Count the number of parenthesis to delete
	var stack []rune
	var toDelete int
	for _, v := range runes {
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1] // pop
			} else {
				toDelete++
			}
		}
		// ignore other runes
	}
	// count the rest of errors
	for len(stack) > 0 {
		toDelete++
		stack = stack[:len(stack)-1]
	}

	toKeep := len(runes) - toDelete

	acc := make(map[string]bool)
	var generate func([]rune, *[]rune, int, int, *[]rune)
	generate = func(src []rune, stack *[]rune, pos int, toTake int, prefix *[]rune) {
		if toTake == 0 && len(*stack) == 0 {
			// save prefix - good combo
			acc[string(*prefix)] = true
			return
		}
		if pos == len(src) {
			return
		}
		// Try to take current position
		if src[pos] != '(' && src[pos] != ')' {
			*prefix = append(*prefix, src[pos])
			generate(src, stack, pos+1, toTake-1, prefix)
			*prefix = (*prefix)[:len(*prefix)-1] // backtrack
			return
		}
		if src[pos] == '(' && toTake > len(*stack) { // still can satisfy the correctnes of sequence
			// take
			*prefix = append(*prefix, src[pos])
			*stack = append(*stack, src[pos])
			generate(src, stack, pos+1, toTake-1, prefix)
			*prefix = (*prefix)[:len(*prefix)-1]
			*stack = (*stack)[:len(*stack)-1]
		}
		if src[pos] == ')' && len(*stack) > 0 {
			// take and pop stack
			*prefix = append(*prefix, src[pos])
			*stack = (*stack)[:len(*stack)-1]
			generate(src, stack, pos+1, toTake-1, prefix)
			*prefix = (*prefix)[:len(*prefix)-1]
			*stack = append(*stack, '(')
		}
		// not take
		// need to take 5 and altogether 8 and looking at
		// 3 (0, 1, 2, 3) (still to look at len(src) - pos - 1)
		if toTake <= len(src)-pos-1 {
			// not take and recurse
			generate(src, stack, pos+1, toTake, prefix)
		}
	}
	var prefix []rune
	generate(runes, &stack, 0, toKeep, &prefix)
	var ans []string
	for k, _ := range acc {
		ans = append(ans, k)
	}
	return ans
}
