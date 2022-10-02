func letterCombinations(digits string) []string {
	/*
	   2 - abc
	   3 - def
	   4 - ghi
	   5 - jkl
	   6 - mno
	   7 - pqrs
	   8 - tuv
	   9 - wxyz

	   acc []string
	   combos(src []rune, i int, prefix []rune)
	*/

	if len(digits) == 0 {
		return []string{}
	}

	keyboard := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'}}

	var ans []string
	var combos func([]rune, int, *[]rune)
	combos = func(src []rune, pos int, prefix *[]rune) {
		if pos >= len(src) {
			ans = append(ans, string(*prefix))
			return
		}

		for _, v := range keyboard[src[pos]] {
			*prefix = append(*prefix, v)
			combos(src, pos+1, prefix)
			*prefix = (*prefix)[:pos]
		}
	}
	var prefix []rune
	combos([]rune(digits), 0, &prefix)
	return ans
}
