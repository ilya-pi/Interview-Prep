func intersection(nums1 []int, nums2 []int) []int {
	/*
	   We can probably utilise map in a straightforward way for this.

	   If we put all nums1 in a map1 and while iterating on nums2 put the results in ans, while deleting elements from map1 if we processed them already (this will ensure uniqueness)
	*/
	map1 := make(map[int]bool)
	for _, v := range nums1 {
		map1[v] = true
	}

	var ans []int
	for _, v := range nums2 {
		if _, ok := map1[v]; ok {
			ans = append(ans, v)
			delete(map1, v)
		}
	}
	return ans
}
