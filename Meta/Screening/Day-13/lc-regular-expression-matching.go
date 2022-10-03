func isMatch(s string, p string) bool {
	/*

		    approachy:

		    we have s and p

		    //
		    if s empty and p empth - yes
		    if s empty, p - not - false
		    if s not empth, p - yes = false
		    //

		    if first character of s is matched (first of p is the same or '.')
		    then if second character in p is *, we can match: 1) rest s with p, 2) rest s with -2p, 3) s with -2p
		    if it is not -> we can match 1) rest s with -1p

		    if the first character of s is not matched (not first of p and first p is not '.')
		      if second character in p is *, then we can match s and -2p
		      if it is not -> false

		ab
		.*c

	*/

	var match func([]rune, []rune) bool
	match = func(s []rune, p []rune) bool {
		if len(s) == 0 && len(p) == 0 {
			return true
		}
		if len(s) != 0 && len(p) == 0 {
			return false
		}
		if len(s) == 0 && len(p) != 0 {
			if len(p) > 1 && p[1] == '*' {
				return match(s, p[2:])
			} else {
				return false
			}
		}

		if s[0] == p[0] || p[0] == '.' {
			if len(p) > 1 && p[1] == '*' {
				return match(s[1:], p) || match(s[1:], p[2:]) || match(s, p[2:])
			} else {
				return match(s[1:], p[1:])
			}
		} else {
			if len(p) > 1 && p[1] == '*' {
				return match(s, p[2:])
			} else {
				return false
			}
		}
	}
	return match([]rune(s), []rune(p))

}
