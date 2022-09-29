func isOneEditDistance(s string, t string) bool {
	if len(s) < len(t) {
		s, t = t, s
	}

	rs, rt := []rune(s), []rune(t)

	if len(s) == len(t)+1 {
		token := true
		for p1, p2 := 0, 0; p1 < len(rs); {
			if p2 < len(rt) && rs[p1] != rt[p2] {
				if !token {
					return false
				}
				token = false
				p1++
			} else {
				p1++
				p2++
			}
		}
		return true
	} else if len(s) == len(t) {
		token := true
		for i, vs := range rs {
			if vs != rt[i] {
				if !token {
					return false
				}
				token = false
			}
		}
		return !token // exactly 1 edit should be present
	}
	return false

}
