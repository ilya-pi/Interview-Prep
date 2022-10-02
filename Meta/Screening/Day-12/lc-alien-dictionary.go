func alienOrder(words []string) string {
	/*
	   In every pair, the first different rune is smaller then the second in two words

	   That create an adjacency table. Then we need to do topoligical sorting of that and that would be the desired word

	   A cycle makes producing a solution impossible
	*/

	/*
	   if len(words) == 1 {
	       // return first character
	       return string(words[0][:1])
	   }
	*/

	inspectOrder := func(s1, s2 string) (rune, rune, bool) {
		rs1, rs2 := []rune(s1), []rune(s2)
		i := 0
		for ; i < len(rs1) && i < len(rs2); i++ {
			if rs1[i] != rs2[i] {
				return rs1[i], rs2[i], true
			}
		}
		return ' ', ' ', false
	}

	// Technically this seems like wrong logic, as we don't know about their ordering with other characters
	adj := make(map[rune][]rune)
	for _, word := range words {
		for _, v := range []rune(word) {
			adj[v] = []rune{}
		}
	}
	for i := 0; i < len(words)-1; i++ {
		aw, bw := words[i], words[i+1]

		// Wrong input that is not caught with a cycle in a graph
		minLen := len(aw)
		if len(bw) < minLen {
			minLen = len(bw)
		}
		if len(aw) > len(bw) && aw[:minLen] == bw[:minLen] {
			return ""
		}

		a, b, ok := inspectOrder(aw, bw)
		if !ok {
			continue
		}
		// it should be a directed acyclic graph
		adj[a] = append(adj[a], b)
	}

	var st []rune
	visited := make(map[rune]bool)
	processing := make(map[rune]bool)
	var topologicalSort func(rune) bool
	topologicalSort = func(r rune) bool {
		visited[r] = true
		processing[r] = true
		for _, n := range adj[r] {
			if processing[n] {
				return true
			}
			if visited[n] {
				continue
			}
			hasLoop := topologicalSort(n)
			if hasLoop {
				return true
			}
		}
		processing[r] = false
		st = append(st, r)
		return false
	}
	for t, _ := range adj {
		if visited[t] {
			continue
		}
		hasLoop := topologicalSort(t)
		if hasLoop {
			return ""
		}
	}
	var ans []rune
	for i := len(st) - 1; i >= 0; i-- {
		ans = append(ans, st[i])
	}
	return string(ans)
}
