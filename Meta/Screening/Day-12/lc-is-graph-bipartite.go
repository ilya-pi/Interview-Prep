func isBipartite(graph [][]int) bool {
	/*
	   To understand that graph is bipartite we will paint nodes in two colors while walking the graph

	   If we happen to walk into a node that is painted differently then we'd want -> it is not bipartite

	   If we were able to walk the whole graph with painting nodes -> it is bipartite

	   We will need
	   colors array
	   bfs walk

	   0 - 1
	   | \ |
	   3 - 2

	*/

	colors := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		// for all nodes
		if colors[i] > 0 { // visited this node already
			continue
		}

		q := []int{i}
		colors[i] = 1
		for len(q) > 0 {
			el := q[0]
			q = q[1:]

			col := colors[el]
			nextCol := 2 // two colors: 1 and 2, both > 0
			if col == 2 {
				nextCol = 1
			}

			for _, n := range graph[el] {
				if colors[n] > 0 {
					if colors[n] != nextCol {
						return false
					}
					continue
				}
				colors[n] = nextCol
				q = append(q, n)
			}
		}
	}
	return true
}
