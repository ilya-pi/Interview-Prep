func removeDuplicates(nums []int) int {
	/*
	   Approach:

	   Run through the array with two pointers, one at the last unique element, another looking for next unique element

	   If we find unique element -> swap to the next after last unique one

	   Loops exits when second pointer reached end of the array

	   Tests:

	   0 1 1 2 2 2 2 2 3 4 7 88 9 32

	   0 0
	   0

	   0 1

	   i == 1 j == 2 ls == 1
	   0 1 (1)

	   i == 1 j == 3 ls == 1
	   0 1 2 (2)

	   ls = 2
	*/
	// len(nums) >= 1 in input, skipping explicit validation
	i := 0 // last unique element placed
	lastSeen := nums[0]
	for j := 1; j < len(nums); j++ {
		if lastSeen != nums[j] {
			i++ // last correctly placed element
			// don't care about swapping, as overwrite that memmory
			nums[i] = nums[j]
			lastSeen = nums[j]
		}
	}
	return i + 1 // amount of unique elements

}
