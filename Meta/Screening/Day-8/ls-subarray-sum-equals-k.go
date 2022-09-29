func subarraySum(nums []int, k int) int {
	/*

	   6 23 -29 8 29 1 3 5

	   We can track the sums we saw so far and at i-th element,
	   the number of subArrays with a given sum k would be:
	     if sum_i == k => +1
	     if have sum_j (j<i) that sum_i - sum_j == k => +1 (+x, amount of j's so that sum_i - sum_j == k)

	   We will need a map to track the sum's we saw so far

	*/
	var ans int
	sums := make(map[int]int)
	var sum int
	for _, v := range nums {
		sum += v
		if sum == k {
			ans++
		}
		if extra, ok := sums[sum-k]; ok {
			ans += extra
		}
		sums[sum]++
	}
	return ans
}
