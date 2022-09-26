func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
	   Approach:

	   1/ Merge two sorted arrays into a new sorted array and pick median with middle element or two different ones / 2 if % 2 == 0
	*/

	// 1 merge sorted arrays
	arr := make([]int, len(nums1)+len(nums2))
	i, j1, j2 := 0, 0, 0
	for j1 < len(nums1) && j2 < len(nums2) {
		if nums1[j1] <= nums2[j2] {
			arr[i] = nums1[j1]
			j1++
		} else {
			arr[i] = nums2[j2]
			j2++
		}
		i++
	}

	// bleed the remaining slice
	for ; j1 < len(nums1); j1++ {
		arr[i] = nums1[j1]
		i++
	}

	for ; j2 < len(nums2); j2++ {
		arr[i] = nums2[j2]
		i++
	}

	// 2 find median
	if len(arr)%2 == 1 {
		// len 7 , 0 .. 6 elements, 7 / 2 == 3, 0 1 2 3 4 5 6
		return float64(arr[len(arr)/2])
	} else {
		var ans float64
		// len 8, 0 .. 7, 8 / 2 == 4, 0 1 2 3 4 5 6 7
		mid := len(arr) / 2
		ans += float64(arr[mid])
		ans += float64(arr[mid-1])
		ans /= 2
		return ans
	}
}
