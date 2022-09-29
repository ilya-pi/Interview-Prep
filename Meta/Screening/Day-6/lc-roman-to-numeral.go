func romanToInt(s string) int {
	/*
		I             1
		V             5
		X             10
		L             50
		C             100
		D             500
		M             1000

		I can be placed before V (5) and X (10) to make 4 and 9.
		X can be placed before L (50) and C (100) to make 40 and 90.
		C can be placed before D (500) and M (1000) to make 400 and 900.

		2022
		MMXXII

		2019
		XXXIX

		1. We will need an acc, that will be `ans`.

		2. We will scan the string left to right, rune by rune (though here bytes will be the same)

		3. If we see a I, X or C, they need to be placed in a separate acc or size 2, on next rune we decide to either 1) flush the accumulator as a single value of two runes 2) as a single rune and replenish it again or 3) as two independent runes

		4. At the end we flush accumulator as an independent rune if not empty
	*/

	// todo: input validation with regex [IVXLCDM]+

	combos := map[[2]rune]int{
		{'I', 'V'}: 4,
		{'I', 'X'}: 9,
		{'X', 'L'}: 40,
		{'X', 'C'}: 90,
		{'C', 'D'}: 400,
		{'C', 'M'}: 900,
	}

	isCombo := func(r1, r2 rune) bool {
		if _, ok := combos[[2]rune{r1, r2}]; ok {
			return true
		}
		return false
	}

	comboToNum := func(arr []rune) int {
		if len(arr) != 2 {
			// return err otherwise and adjust business logic
			panic("bad input")
		}
		if ans, ok := combos[[2]rune{arr[0], arr[1]}]; ok {
			return ans
		}
		// return err otherwise and adjust business logic
		panic("bad input")
	}

	runes := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	runeToNum := func(r rune) int {
		if ans, ok := runes[r]; ok {
			return ans
		}
		panic("bad input and bad error management by Ilya")
	}

	var ans int    // int 64 most likely
	var acc []rune // we will use slice as we will append and cut it, but the underlying array won't be re-allocated
	for _, v := range []rune(s) {
		switch {
		// can be placed before other runes
		case v == 'I':
			fallthrough
		case v == 'X':
			fallthrough
		case v == 'C':
			// we have case when len(acc) == 1 -> then we need to flush prev and replenish
			// if len(acc) == 0, -> add and continue
			if len(acc) == 1 {
				if isCombo(acc[0], v) {
					acc = append(acc, v)
					ans += comboToNum(acc)
					acc = nil
				} else {
					ans += runeToNum(acc[0])
					acc[0] = v
				}
			} else if len(acc) == 0 {
				acc = append(acc, v)
			}

		// we have case when len(acc) == 1 -> then we need to flush prev and replenish
		// if len(acc) == 0, -> add and continue

		// cannot be placed before other runes
		case v == 'V':
			fallthrough
		case v == 'L':
			fallthrough
		case v == 'D':
			fallthrough
		case v == 'M':
			// we have case when len(acc) == 1 and it makes combo -> add combo, clean acc
			// we have case when len(acc) == 0 -> add value
			if len(acc) == 1 {
				if isCombo(acc[0], v) {
					acc = append(acc, v)
					ans += comboToNum(acc)
					acc = nil
				} else {
					ans += runeToNum(acc[0])
					acc = nil
					ans += runeToNum(v)
				}
			} else if len(acc) == 0 {
				ans += runeToNum(v)
			}
		}
	}
	if len(acc) != 0 {
		ans += runeToNum(acc[0])
	}
	return ans
}
