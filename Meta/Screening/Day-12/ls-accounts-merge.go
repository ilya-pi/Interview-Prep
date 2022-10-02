import "sort"

func accountsMerge(accounts [][]string) [][]string {
	/*
	   I kind of want to solve this with maps, but may be it is not the best idea

	   But let's explore it:

	   Name1 -> email11, email12, email13
	   Name2 -> email21, email22
	   Name3 -> email31, email12

	   email11 -> 1
	   email12 -> 1
	   email13 -> 1
	   email21 -> 2
	   email22 -> 2
	   email31 -> 3
	   email12 -> 3, 1

	   Then for each emailxy we will add all e-mails to the smallest nameid and set their link to the smallest nameid, then if somebody would want to merge into that user, it will merge in the smaller nameid. We also need to keep track of merged names, to skip them on final iteration

	   O(n*m + n*m *(klogk + m) + n*m * m + n*mlogm) = o(nm^2 + n*mklokg)
	   n - amount of people
	   m - longest email list
	   k - amount of matching users

	   Of with graphs potentially:

	   We create nodes for all entries, marking name node as name node and merging correct e-mails in the graph

	   Then we bfs all nodes, returning all email nodes as slice and one (only) named node, then we can sort slice. We run it for all unvisited nodes

	   Thus: O(m*n + n*mlogm)
	*/

	type Node struct {
		val    string
		isName bool
	}
	adj := make(map[*Node][]*Node)
	emailNodes := make(map[string]*Node)
	for _, account := range accounts {
		name := account[0]
		nameNode := &Node{val: name, isName: true}
		for i := 1; i < len(account); i++ {
			email := account[i]
			emailNode := emailNodes[email]
			if emailNode == nil {
				emailNodes[email] = &Node{val: email}
				emailNode = emailNodes[email]
			}
			// connect them to name
			adj[nameNode] = append(adj[nameNode], emailNode)
			adj[emailNode] = append(adj[emailNode], nameNode)
		}
	}

	// We have adjacency table now
	visited := make(map[*Node]bool)
	bfs := func(n *Node) (string, []string) {
		var name *string
		emailsMap := make(map[string]bool)

		q := []*Node{n}
		visited[n] = true
		for len(q) > 0 {
			el := q[0]
			q = q[1:]

			if el.isName && name == nil {
				name = &el.val
				// can also skip the nil check, as it should be the same
			}
			if !el.isName {
				// record e-mail
				emailsMap[el.val] = true
			}
			for _, ch := range adj[el] {
				if visited[ch] {
					continue
				}
				visited[ch] = true
				q = append(q, ch)
			}
		}
		var emails []string
		for k, _ := range emailsMap {
			emails = append(emails, k)
		}
		return *name, emails

	}

	var ans [][]string
	// Core part
	for node, _ := range adj {
		if !visited[node] {
			name, emails := bfs(node)
			sort.Strings(emails)
			merged := append([]string{name}, emails...)
			ans = append(ans, merged)
		}
	}
	return ans
}
