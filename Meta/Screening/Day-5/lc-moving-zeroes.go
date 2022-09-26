func moveZeroes(nums []int) {
	/*
	   Approach: in order to minimize amount of swaps, we will use two pointers
	   It is a bit similar to partition in quicksort

	*/

	for p1, p2 := 0, 0; p1 < len(nums) && p2 < len(nums); {
		if nums[p1] == 0 && nums[p2] != 0 && p1 < p2 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
		}
		if nums[p1] != 0 {
			p1++
		}
		if nums[p2] == 0 || p2 < p1 {
			p2++
		}
	}

}
