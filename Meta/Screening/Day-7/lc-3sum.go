import "sort"

func threeSum(nums []int) [][]int {
	/*

	   Brute force:

	   Iterate with i, j, k from 0 to len-1, i+1 to len-1, j+1 to len-1 and find all triplets, since all indices are not intersecting, they will be unique

	   O(n^3), space to store answers

	   Iteration over it:

	   We don't need the last cycle, if we have a map to find the matching element, as long as it is >i and j; and there might be multiple options for k, hence we need map of int to []int

	   O(n^2) and O(n) space

	   -1,0,1,2,-1,-4

	   lookup:
	   -1 -> 0, 4
	   0 -> 1
	   1 -> 2
	   2 -> 3
	   -4 -> 5

	   -1 0 ? we have +1? -> -1 0 1
	   -1 1 ? we have 0? yes at 1 < j - > no
	   -1 2 ? we have -1? yes at 0 and 4 -> -1 2 -1
	   -1 0 1

	*/

	var ans [][]int
	lookup := make(map[int]int) // value to indices of that value, keeping only the last one
	for i, v := range nums {
		lookup[v] = i
	}

	filter := make(map[[3]int]bool)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)-1; j++ { // j can't be last, as we also need k
			target := -(nums[i] + nums[j])
			if k, ok := lookup[target]; ok && k > j {
				opt := [3]int{nums[i], nums[j], nums[k]}
				sort.Slice(opt[:], func(i1, i2 int) bool { return opt[i1] < opt[i2] })
				if _, saw := filter[opt]; saw {
					continue
					// saw this option already, no need to
					//consider other indices for the same sum
				}
				filter[opt] = true
				ans = append(ans, opt[:])
			}
		}
	}
	return ans
}
