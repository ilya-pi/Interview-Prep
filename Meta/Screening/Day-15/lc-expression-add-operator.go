import "strconv"

func addOperators(num string, target int) []string {
	/*
	   A very ugly recursion?

	   3456237490

	   At each character we try to insert either of +-* to get to the target value, if at the end we have it -> great

	   We will try to reverse and memoise it once we have it


	   getSequences(src string, acc int, pos int, target int)
	       if pos == len(src) {
	           if acc == target {
	               record it
	           }
	           return
	       }

	       num at pos
	       getSequences(src, acc + nap, pos+1, target)
	       getSequences(src, acc - nap, pos+1, target)
	       getSequences(src, acc * nap, pos+1, target)

	*/

	var ans []string
	var prefix []rune
	var generate func([]rune, int, int, int, int, int, *[]rune)
	generate = func(src []rune, target, from, to, acc int, prev int, prefix *[]rune) {
		//fmt.Printf("acc = %d, target = %d, prefix %s, from-to: %s\n", acc, target, string(*prefix), string(src[from:to]))
		if from == len(src) {
			if acc == target {
				r := make([]rune, len(*prefix))
				copy(r, *prefix)
				ans = append(ans, string(r))
			}
			return
		}

		v := src[from:to]
		// todo(ilya): error handling
		vInt64, _ := strconv.ParseInt(string(v), 10, 64)
		vInt := int(vInt64)

		// empty case
		if to < len(src) && src[from] != '0' {
			generate(src, target, from, to+1, acc, prev, prefix)
		}

		// -
		*prefix = append(*prefix, '-')
		*prefix = append(*prefix, v...)
		generate(src, target, to, to+1, acc-vInt, -vInt, prefix)
		*prefix = (*prefix)[:len(*prefix)-len(v)-1]

		// +
		*prefix = append(*prefix, '+')
		*prefix = append(*prefix, v...)
		generate(src, target, to, to+1, acc+vInt, +vInt, prefix)
		*prefix = (*prefix)[:len(*prefix)-len(v)-1]

		// *
		*prefix = append(*prefix, '*')
		*prefix = append(*prefix, v...)
		acc -= prev
		acc += prev * vInt
		generate(src, target, to, to+1, acc, prev*vInt, prefix)
		*prefix = (*prefix)[:len(*prefix)-len(v)-1]
	}
	runes := []rune(num)
	if len(runes) == 1 {
		if int(num[0]-'0') == target {
			return []string{num}
		} else {
			return []string{}
		}
	}

	for i := 1; i <= len(runes); i++ {
		if i > 1 && runes[0] == '0' {
			break
		}
		firstVal, _ := strconv.ParseInt(string(runes[0:i]), 10, 64)
		fv := int(firstVal)
		prefix = append(prefix, runes[0:i]...)
		generate(runes, target, i, i+1, fv, fv, &prefix)
		prefix = prefix[:0]
	}
	return ans
}
