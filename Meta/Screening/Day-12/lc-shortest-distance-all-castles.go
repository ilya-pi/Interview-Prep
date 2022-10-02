import "math"

func shortestDistance(grid [][]int) int {
	/*
	   Approach could be to find shortest distance from all nodes with 1 to all empty ones with 0 and then find the lowest number in that grid
	*/

	if len(grid) == 0 {
		return -1
	}

	// Init dist grid
	dists := make([][]int, len(grid))
	for i, _ := range dists {
		dists[i] = make([]int, len(grid[0]))
	}
	var castleCount int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				dists[i][j] = math.MaxInt
			}
			if grid[i][j] == 1 {
				dists[i][j] = math.MaxInt
				castleCount++
			}
		}
	}
	reachable := make([][]int, len(grid))
	for i := range reachable {
		reachable[i] = make([]int, len(grid[0]))
	}

	adj := func(i, j int) [][2]int {
		var r [][2]int
		for _, opt := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			vert, horiz := i+opt[0], j+opt[1]
			if vert >= 0 && vert < len(grid) && horiz >= 0 && horiz < len(grid[0]) && grid[vert][horiz] != 2 && grid[vert][horiz] != 1 {
				r = append(r, [2]int{vert, horiz})
			}
		}
		return r
	}

	// run bfs and record shortest distances
	bfs := func(i, j int) {
		visited := make([][]bool, len(grid))
		for i, _ := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		visited[i][j] = true

		q := [][3]int{{i, j, 0}}
		for len(q) > 0 {
			el := q[0]
			v, h, d := el[0], el[1], el[2]
			q = q[1:]

			if grid[v][h] == 0 {
				dists[v][h] += d
				reachable[v][h]++
			}
			// walk kids
			for _, ch := range adj(v, h) {
				v1, h1 := ch[0], ch[1]
				if visited[v1][h1] {
					continue
				}
				visited[v1][h1] = true
				q = append(q, [3]int{v1, h1, d + 1})
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
				/*
				   for _, v := range dists {
				       fmt.Printf("%v\n", v)
				   }
				   fmt.Printf("\n")
				*/

			}
		}
	}
	ans := math.MaxInt
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if reachable[i][j] == castleCount && dists[i][j] < ans {
				ans = dists[i][j]
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}
