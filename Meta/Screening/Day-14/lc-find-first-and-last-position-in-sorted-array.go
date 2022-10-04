import "fmt"

func searchRange(nums []int, target int) []int {
	/*
	   I would guess if it is logN, then we would need to modified binary searches

	   left, right
	   mid = (l + r) / 2
	   if mid >= target -> search left .. mid - 1
	    otherwise mid < target -> mid+1 right


	   0 1 1 2 3 4 _4_ 4 4 9 19 42
	   0 1 1 _4_ 4 4
	   0 _1_ 1
	   1
	   4 -- if left == right -> exit


	   0 1 1 2 _3_ 4 4 4
	   4 _4_ 4
	   _4_
	   l, m-1 --> l is the index we are looking for

	   0 1 1 2 3 _4_ 4 4 4 4 4
	   0 1 _1_ 2 3
	   2 _3_
	         m+1,r and l > r --> new l is the left point

	   l r
	   m
	   if m < 4 -> in m+1, r
	   if m == 4 -> in l, m-1

	   other side:
	   4 4 _9_ 19 42
	   4 _4_ -> r is the index == 4

	   4 4 4 _4_ 9 19 42
	   9 _19_ 42
	   _9_ -> l,m-1 and r is the index we are looking for




	   l r
	   m
	   if m > 4 -> in l,m-1
	   if m == 4 -> in m+1, r
	   ----


	   Let's look at how a binary search for first position might work


	*/

	if len(nums) == 0 {
		return []int{-1, -1}
	}
	/*
	   1, 3
	   1
	*/
	// Find one of value's index
	l, r := 0, len(nums)-1
	ind := -1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			ind = m
			break
		}
		if nums[m] < target {
			l, r = m+1, r
		} else if nums[m] > target {
			l, r = l, m-1
		}
	}
	if ind == -1 {
		return []int{-1, -1}
	}
	fmt.Printf("ind %d\n", ind)

	// Find left side
	l, r = 0, ind
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l, r = m+1, r
		} else if nums[m] == target {
			l, r = l, m-1
		}
	}
	left := l

	/*
	   5,7,7,_8_,8,10
	   3

	   8,_8_,10
	   10

	   1,3
	   1

	   _1_,3

	   1
	*/

	// Find right side
	l, r = ind, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			l, r = m+1, r
		} else if nums[m] > target {
			l, r = l, m-1
		}
	}
	right := r
	return []int{left, right}

}
