func wordBreak2(s string, wordDict []string) bool {
	/*

	   I could try to solve this just recursively.
	   At each step I produce all possible prefixes that can be "consumed" and check if we can break the rest of the string.
	   If either of them works -> return true

	   We could use trie for the dictionary

	   trie
	       edges map[rune]*trie
	       terminates bool

	   getOpt = s ->
	       trie of dictionary
	       for every letter in s
	           if cant step
	               return options
	           trie = step trie
	           if terminates trie
	               add option

	   canBreak = s ->
	       if len(s) == 0
	           return true

	       options := getOpt(s)
	       for opt in options // []rune slicing the same array
	           if canBreak(opt)
	               return true

	       if there are no options that work
	           return false

	*/

	runes := []rune(s)

	// trie
	type Trie struct {
		edges      map[rune]*Trie
		terminates bool
	}
	trie := &Trie{edges: make(map[rune]*Trie)}
	canStep := func(t *Trie, v rune) bool {
		if t == nil {
			return false
		}
		if _, ok := t.edges[v]; ok {
			return true
		}
		return false
	}
	step := func(t *Trie, v rune) *Trie {
		if t == nil {
			return nil
		}
		if ans, ok := t.edges[v]; ok {
			return ans
		} else {
			return nil
		}
	}
	populate := func(t *Trie, dict []string) {
		for _, word := range dict {
			cur := t
			//fmt.Printf("Adding word %s\n", word)
			for _, v := range []rune(word) {
				//fmt.Printf("Added %s\n", string(v))
				if next, ok := cur.edges[v]; ok {
					cur = next
				} else {
					cur.edges[v] = &Trie{edges: make(map[rune]*Trie)}
					cur = cur.edges[v]
				}
			}
			cur.terminates = true
		}
	}
	populate(trie, wordDict)

	memo := make(map[int][]int)
	getOpts := func(src []rune, start int, dict *Trie) []int {
		if memoized, ok := memo[start]; ok {
			return memoized
		}
		var opts []int
		t := dict
		for i, v := range src[start:] {
			//fmt.Printf("Can step is %v for %s, trie = %v\n", canStep(t, v), string(v), t)
			if !canStep(t, v) {
				break
			}
			t = step(t, v)
			if t.terminates {
				opts = append(opts, start+i+1)
			}
		}
		// Let's reverse opts, to first try the longer options
		for i, j := 0, len(opts)-1; i < j; i, j = i+1, j-1 {
			opts[i], opts[j] = opts[j], opts[i]
		}
		memo[start] = opts
		return opts
	}

	//fmt.Printf("Trie is %v\n", trie)

	var canBreak func([]rune, int) bool
	canBreak = func(src []rune, start int) bool {
		if start >= len(src) {
			return true
		}

		options := getOpts(src, start, trie)
		//fmt.Printf("Options for %s at %d are %v\n", string(src), start, options)
		for _, opt := range options {
			if canBreak(src, opt) {
				return true
			}
		}
		// Tried all options, non work
		return false
	}
	return canBreak(runes, 0)

}

func wordBreak(s string, wordDict []string) bool {
	/*

	   Let's put it the other way around.

	   We will have an array dp[0..len(s)]
	   For every position we will see if it _reachable_, going left to right.
	   To understand that position is reachable we will check for every word in dictionary that
	       1) it fits
	       2) the letters are matching
	       3) the position where it would've begin is true

	   Then the worst case would be O(len(dict)*len(average word)*len(s))

	   leet
	   i = 4

	   leet_c_ode

	*/
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
	dict_loop:
		for _, word := range wordDict {
			pos := i - len(word)
			if pos >= 0 && dp[pos] {
				//fmt.Printf("Comparing %s and %s\n", word, s[pos:i])
				if word == s[pos:i] {
					dp[i] = true
					break dict_loop
				}
			}
		}
	}
	return dp[len(s)]

}
