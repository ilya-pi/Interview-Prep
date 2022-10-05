func checkSubarraySum(nums []int, k int) bool {
	/*
	   This sum up to a multiple of k if % k == 0

	   So if I will go from left to right, and keep track of all %k I've seen so far,
	   at each element I can decide whether it is a good match for a continuous sequence

	   Neet to watch not to count myself (when I divide by k), but if there is a 0 in the seen part -> it is a match!

	   Nothing removes elements from the seen part


	   ...... % k
	   ................. %k


	*/

	if len(nums) < 2 {
		return false
	}

	seen := make(map[int][]int)
	sums := make([]int, len(nums))
	var sum int
	for i, v := range nums {
		sum += v
		sums[i] = sum
		sumk := sum % k
		if opts, ok := seen[sumk]; ok {
			for _, ind := range opts {
				if ind == i-1 && nums[ind] == 0 && nums[i] == 0 {
					return true
				}
				if i-ind > 1 && sum-sums[ind] >= k {
					return true
				}
			}
		}
		if sumk == 0 && sum > 0 && i > 0 {
			return true
		}
		seen[sumk] = append(seen[sumk], i)
	}
	return false

}
