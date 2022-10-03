func subsets(nums []int) [][]int {
	/*
	   appraoch:

	   at every step of the recursion, we can either take the element, or not, if we are at the end -> add result. we know we are at the end if we are looking at the pos greater then we have in array

	   ans [][]int

	   src []int, pos int, prefix *[]int

	   1. check if complete
	   2. take the element
	   3. backtrack
	   4. not take the element

	*/

	var ans [][]int

	var genSubsets func([]int, int, *[]int)
	genSubsets = func(src []int, pos int, prefix *[]int) {
		if pos == len(src) {
			// have a set!
			r := make([]int, len(*prefix))
			copy(r, *prefix)
			ans = append(ans, r)
			return
		}

		// not take the element
		genSubsets(src, pos+1, prefix)
		// take the element
		*prefix = append(*prefix, src[pos])
		genSubsets(src, pos+1, prefix)
		// backtrack
		*prefix = (*prefix)[:len(*prefix)-1]
	}
	var prefix []int
	genSubsets(nums, 0, &prefix)
	return ans
}
