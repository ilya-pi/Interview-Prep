func nextPermutation(nums []int) {
	/*

	   7 8 2 3 92 60 21 12 3
	   7 8 2 12 3 3 21 60 92

	   7 8 2 3 92 60 21 12 3



	   So, we need to find first index from right, where the element is smaller then either of the "options"

	   Then find the smallest option, otherwise it is not "next"; and then reverse the tail after found element, to make it as small as possible to be the "next array"

	*/

	reverse := func(arr []int, start, end int) {
		// passing slice by value, as we are not re-alligning underlying array
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	k := len(nums) - 2 // need at least one element to swap
	for ; k >= 0 && nums[k] >= nums[k+1]; k-- {
	}

	if k == -1 {
		// no working option, reverse all arr and return
		reverse(nums, 0, len(nums)-1)
		return
	}

	// we have elem at k, that is smaller then previous
	// let's find the best option
	j := len(nums) - 1
	for ; j > k && nums[k] >= nums[j]; j-- {
	}
	// we have in j smallest element of options that is bigger then k
	// swap them
	nums[k], nums[j] = nums[j], nums[k]
	// minimize the end after k to be the very next permutation
	reverse(nums, k+1, len(nums)-1)
	return
}
