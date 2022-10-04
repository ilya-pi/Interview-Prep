import "math"

func findPeakElement(nums []int) int {
	/*

	   Well, brute force is linear.

	   But logn...

	   if start < end then:

	   And end is not peak

	   then there is an index somewhere between them, so that it first goes up and then goes down

	   To try to find it, we will assume...:


	   smaller .................bigger

	   up up up up up down down down

	   if we split in half and mid
	       is smaller then bigger then atleast there is a peak on the right
	       is bigger then bigger

	   2 3 4 5 6 5 4
	   2 3 4 _5_ 6 5 4 ... but

	   1,2,1,_3_,5,6,4

	   212 -> peak!
	   123 - > there is a peak on the right
	   321 -> there is a peak on the left
	   212 -> there can be peak left and right

	*/

	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]
	}

	/*
	   1 2 1 3 5 6 4
	   1 2 1 _3_ 5 6 4
	   5_6_4 -> exit

	   1 2 3 6 5
	   6 5
	   6

	*/

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		//121
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		//123
		if get(mid)-1 < get(mid) && get(mid) < get(mid+1) {
			left, right = mid+1, right
			continue
		}
		// 321 & 212
		left, right = left, mid-1
		continue
	}
	return left
}
