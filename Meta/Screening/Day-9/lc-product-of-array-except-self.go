func productExceptSelf(nums []int) []int {
	/*
	   [1 10 0 -1 39 20]

	   If we take the array and pass left to right, we can store product of all prefixes to this element (or 1 if there are no prefixed elements)
	   If we run from right to left, then we can multuiply it by the product of suffixes

	   So we have 2 for loops from different sides, starting at 1 elem offset

	   And a special case for 0 or 1 length array
	*/
	if len(nums) < 2 {
		return nums
	}

	ans := make([]int, len(nums))
	// left to right
	ans[0] = 1
	for i, product := 1, nums[0]; i < len(nums); i++ {
		ans[i] = product
		product *= nums[i]
	}

	// right to left
	for i, product := len(nums)-2, nums[len(nums)-1]; i >= 0; i-- {
		ans[i] *= product
		product *= nums[i]
	}

	return ans

}
