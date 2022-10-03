func permute(nums []int) [][]int {
	/*
	   at each step of the recursion we can take either element to be next in the resulting permutation, building up prefix array

	   if len(nums) == 0 {
	       return [][]int{}
	   }

	   perm(map[int]bool, prefix) ->
	       if map is empty -> copy prefix into result

	       for all elems in map,
	           add to prefix,
	           delete from map and
	           recur
	           add to map
	           delete from prefix

	   return result

	   We will need to use array instead of map and swap elemnts to the "processed part", and then repair it

	*/

	if len(nums) == 0 {
		return [][]int{}
	}

	var ans [][]int

	var perm func([]int, int, *[]int)
	perm = func(nums []int, pos int, prefix *[]int) {
		if pos == len(nums) {
			r := make([]int, len(*prefix))
			copy(r, *prefix)
			ans = append(ans, r)
		}

		for i := pos; i < len(nums); i++ {
			el := nums[i]
			*prefix = append(*prefix, el)
			nums[i], nums[pos] = nums[pos], nums[i]
			perm(nums, pos+1, prefix)
			nums[i], nums[pos] = nums[pos], nums[i]
			*prefix = (*prefix)[:len(*prefix)-1]
		}
	}
	var prefix []int
	perm(nums, 0, &prefix)
	return ans
}
