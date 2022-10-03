func findStrobogrammatic(n int) []string {
	/*

	   I would think that symmetry rules are covered by the example in 1:
	   11
	   69
	   88
	   96

	   While possible mid point is covered by example 1:
	   0
	   1
	   8

	   Basically, if we are given number n and:

	   n%2 == 1, then we have 3 options for mid point 0, 1, 8

	   And then we recursively can take for all positions (on one side, the other side is dictated by them: 1, 6, 8 or 9)

	   We will need to backtrack after taking an element and then fill the other side in a for loop, before converting to string

	   in recursion, we would want:

	   toFill, prefix

	   we'll populate results to acc [][]rune

	*/

	var acc [][]rune
	var generate func(int, *[]rune)
	generate = func(toFill int, prefix *[]rune) {
		if toFill == 0 {
			r := make([]rune, len(*prefix))
			copy(r, *prefix)
			acc = append(acc, r)
			return
		}

		for _, v := range []rune{'1', '6', '8', '9', '0'} {
			if toFill == 1 && v == '0' {
				continue
			}
			*prefix = append(*prefix, v)
			generate(toFill-1, prefix)
			// backtrack
			*prefix = (*prefix)[:len(*prefix)-1]
		}
	}
	// alloc memory
	var prefix []rune
	if n%2 == 1 {
		for _, v := range []rune{'0', '1', '8'} {
			prefix = append(prefix, v)
			generate(n/2, &prefix)
			//prefix = prefix[:len(prefix)-1]
			prefix = prefix[:0]
		}
	} else {
		generate(n/2, &prefix)
	}
	// now we need to glue the other parts for each and record to ans
	matchingLeft := func(arr []rune) string {
		ans := make([]rune, len(arr))
		for i := len(arr) - 1; i >= 0; i-- {
			el := arr[i]
			ind := len(arr) - 1 - i
			switch el {
			case '1':
				ans[ind] = '1'
			case '6':
				ans[ind] = '9'
			case '8':
				ans[ind] = '8'
			case '9':
				ans[ind] = '6'
			case '0':
				ans[ind] = '0'
			default:
				// can't happen
			}
		}
		return string(ans)
	}
	var ans []string
	for _, v := range acc {
		var leftPart string
		if n%2 == 1 {
			leftPart = matchingLeft(v[1:])
		} else {
			leftPart = matchingLeft(v)
		}
		ans = append(ans, leftPart+string(v))
	}
	return ans

}
