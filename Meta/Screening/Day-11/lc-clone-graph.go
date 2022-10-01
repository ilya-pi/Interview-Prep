/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	/*
	   DFS approach

	   Keep map if seen nodes so far, to add edges to them.
	*/

	if node == nil {
		return nil
	}

	nodes := map[int]*Node{node.Val: &Node{Val: node.Val}}
	//visited := make(map[int]bool) // ? can reused nodes?
	var dfs func(*Node)
	dfs = func(n *Node) {
		nn := nodes[n.Val]
		for _, n2 := range n.Neighbors {
			// Add edge
			if nn2, ok := nodes[n2.Val]; !ok {
				// not visited
				nn2 = &Node{Val: n2.Val}
				nodes[n2.Val] = nn2
				nn.Neighbors = append(nn.Neighbors, nn2)
				dfs(n2)
			} else {
				// visited
				nn.Neighbors = append(nn.Neighbors, nn2)
			}
		}
	}
	dfs(node)
	return nodes[node.Val]
}
