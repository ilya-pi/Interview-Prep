func permuteUnique(nums []int) [][]int {
	/*

	   Unique permutations

	   > Persist them in the tree and then just dfs all results

	   So we generate permutations recursively in a similar way, but then we store them to the tree

	   In a tree that has root as empty node
	*/

	if len(nums) == 0 {
		return [][]int{}
	}

	type Tree struct {
		children map[int]*Tree // we will out numbers on edges
	}

	root := &Tree{children: make(map[int]*Tree)}
	add := func(arr []int) {
		cur := root
		for _, v := range arr {
			if next, ok := cur.children[v]; ok {
				cur = next
			} else {
				cur.children[v] = &Tree{children: make(map[int]*Tree)}
				cur = cur.children[v]
			}
		}
	}
	var ans [][]int
	var dfs func(*Tree, *[]int)
	dfs = func(t *Tree, prefix *[]int) {
		if len(t.children) == 0 {
			r := make([]int, len(*prefix))
			copy(r, *prefix)
			ans = append(ans, r)
			return
		}
		for v, n := range t.children {
			*prefix = append(*prefix, v)
			dfs(n, prefix)
			*prefix = (*prefix)[:len(*prefix)-1]
		}
	}

	var perm func([]int, int, *[]int, func([]int))
	perm = func(nums []int, pos int, prefix *[]int, add func([]int)) {
		if pos == len(nums) {
			add(*prefix)
			return
		}
		for i := pos; i < len(nums); i++ {
			*prefix = append(*prefix, nums[i])
			nums[i], nums[pos] = nums[pos], nums[i]
			perm(nums, pos+1, prefix, add)
			nums[i], nums[pos] = nums[pos], nums[i]
			*prefix = (*prefix)[:len(*prefix)-1]
		}
	}
	var prefix []int
	perm(nums, 0, &prefix, add)
	dfs(root, &prefix)
	return ans
}
