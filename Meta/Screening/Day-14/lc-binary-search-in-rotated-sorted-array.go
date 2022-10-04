func search(nums []int, target int) int {
	/*
	   binary search where it was rotated
	   and then binary search in the correct part

	   [4,5,6,7,0,1,2]
	          7,0,1,2
	          7,0,1
	          7,0


	   if start < end
	       there was rotation
	       bs start end
	           for start+1 < end {
	             mid = between them
	             if mid < start - choose left mid
	             if mid > start - choose mid right
	           }
	       we hace end of left in start
	       and start of right in end

	       if el <= end of arr - search right
	       else searrch - left

	   bs arr, left, right
	       if left >= right {
	           return -1
	       }
	     mid
	     if el == mid -> return
	     if el < mid -> bs left
	     if el > mid -> bs right

	*/
	// 4,5,6,7,0,1,2
	/*
	   4 6
	   mid = 5
	   4 4
	*/
	// Binary search on array
	var bs func([]int, int, int, int) int
	bs = func(arr []int, target, left int, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			return bs(arr, target, mid+1, right)
		} else {
			return bs(arr, target, left, mid-1)
		}
		return -1 // unreachable
	}

	// Find if it was rotated
	if nums[0] > nums[len(nums)-1] {
		// It was rotated!
		left, right := 0, len(nums)-1
		for left+1 < right {
			mid := (left + right) / 2
			if nums[mid] < nums[left] {
				left, right = left, mid
			} else if nums[mid] > nums[left] {
				left, right = mid, right
			}
		}
		// Now pick the side to search
		if target <= nums[len(nums)-1] {
			// Search right part
			//fmt.Printf("Search right: %d - %d\n", nums[right], nums[len(nums)-1])
			return bs(nums, target, right, len(nums)-1)
		} else {
			// Search left part
			//fmt.Printf("Search left: %d - %d\n", nums[0], nums[left])
			return bs(nums, target, 0, left)
		}
	} else {
		return bs(nums, target, 0, len(nums)-1)
	}
}
