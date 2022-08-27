package main

import "fmt"

// remove is an order destroying fast way to remove element from a slice
func remove(arr []string, i int) []string {
	if i < 0 || i >= len(arr) {
		return arr
	}
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func order(projects []string, dependencies [][]string) []string {
	/*
	   It seems that if we represented the dependencies as a
	   graph and then did a bfs over all non-visited nodes,
	   while detecting for loops, we'd have our build order in the inverse
	*/

	e := map[string][]string{}
	deps := map[string]map[string]bool{}
	for _, v := range dependencies {
		// v[1] should be built before v[0]
		// v[1] is dependant on v[0]
		e[v[0]] = append(e[v[0]], v[1])
		if _, ok := deps[v[1]]; !ok {
			deps[v[1]] = map[string]bool{}
		}
		deps[v[1]][v[0]] = true
	}

	fmt.Printf("Deps: %v\nEdges: %+v\n", deps, e)

	var res []string
	q := []string{}
	for _, p := range projects {
		if len(deps[p]) == 0 {
			// No nodes are leading to this project to be built
			// We will add only the nodes that have no incoming edges to the q, meaning they can be built
			q = append(q, p)
		}
	}

	fmt.Printf("q is %v\n", q)

	// Start the "build" process
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		fmt.Printf("p == %v\nq == %v\ne[p] == %v\n\n", p, q, e[p])
		res = append(res, p)
		for _, p2 := range e[p] {
			fmt.Printf("Removing edge from %v\n", p2)
			// Remove this dependency, as it was build already
			if _, ok := deps[p2]; ok {
				delete(deps[p2], p)
			}
			// Check if we can build p2
			if len(deps[p2]) == 0 {
				q = append(q, p2)
			}
		}
	}

	fmt.Printf("res is %v\n", res)

	if len(res) == len(projects) {
		return res
	}

	return nil
}

func main() {
	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}}
	// output: f, e, a, b, d, c
	fmt.Printf("Order is %v\n", order(projects, dependencies))
}
