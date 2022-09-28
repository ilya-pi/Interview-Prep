func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
	   To this in place, first thing I can think of is to add values of nums2 at the end and bubble them down where they should be

	   We first put them in nums1
	   Then from 0..m-1, so for i := m .. len(nums1) -- bubble elements down,
	   that is a for loop inside of the first loop
	*/

	/*
	   // Copy elements from nums2 to nums1
	   for i, v := range nums2 {
	       nums1[m+i] = v
	   }

	   // Bubble them in place
	   for i := m; i < len(nums1); i++ {
	       for k := i; k > 0 && nums1[k-1] > nums1[k]; k-- {
	           nums1[k-1], nums1[k] = nums1[k], nums1[k-1]
	       }
	   }
	*/

	// This is O(m*n)
	// I can come up with O(m+n) with extra memory, with two pointers
	// But it seems I can also do it in place:

	/* Approach:
	   We will go with 2 pointers from the end of arrays and populate the array nums1 from the end in a cycle; while respecting the limits of p1 and p2
	*/
	/*
	   [1, 2, 4, 0, 0, 0, 0]
	   [2, 3, 6, 7]

	   p1 p2
	   6  3

	   i = 6



	*/
	p1, p2 := m-1, len(nums2)-1
	for i := len(nums1) - 1; i >= 0; i-- {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] > nums2[p2] {
				// then we take nums1[p1]
				nums1[i] = nums1[p1]
				p1--
			} else {
				nums1[i] = nums2[p2]
				p2--
			}
			continue
		}
		if p1 >= 0 {
			nums1[i] = nums1[p1]
			p1--
			continue
		}
		if p2 >= 0 {
			nums1[i] = nums2[p2]
			p2--
			continue
		}
	}

}
