package main

import "fmt"

/*
Have trie, seed it with 10 words and then ask all words with prefix car
*/

type Trie struct {
	terminates bool
	children   map[rune]*Trie
}

func (t *Trie) add(word string) {
	if t.children == nil {
		t.children = make(map[rune]*Trie)
	}
	// Walk down and add
	at := t
	for _, v := range []rune(word) {
		// Do we have this child at at?
		if ch, ok := at.children[v]; ok {
			// Yes -> just step in that branch
			at = ch
		} else {
			// No -> create that branch and step in it
			n := &Trie{children: make(map[rune]*Trie)}
			at.children[v] = n
			at = at.children[v]
		}
	}
	// We need to add "terminates" var to signal that it is a word
	at.terminates = true
}

func (t *Trie) all(prefix string) []string {
	if t.children == nil {
		return nil
	}
	at := t
	// Walk in the trie till prefix
	for _, v := range []rune(prefix) {
		if ch, ok := at.children[v]; ok {
			at = ch
		} else {
			// No further words found :-(
			return nil
		}
	}

	// Return all terminating nodes in that sub-trie
	var ans []string
	// Dfs sub-trie and append to ans on terminating nodes
	var dfs func(*Trie, string)
	dfs = func(n *Trie, prefix string) {
		// Check all that terminates
		if n.terminates {
			ans = append(ans, prefix)
		}

		// Walk children
		// Children will not be nil as we always initialize it
		for k, ch := range n.children {
			dfs(ch, prefix+string(k))
		}
	}
	dfs(at, prefix)
	return ans
}

func main() {
	var t Trie
	t.add("car")
	t.add("carpet")
	t.add("volvo")
	t.add("castle")
	t.add("candy")
	t.add("country")
	prefix := "ca"
	fmt.Printf("All starting with %q: %v\n", prefix, t.all(prefix))
}
