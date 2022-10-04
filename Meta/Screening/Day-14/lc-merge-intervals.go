import "sort"

func merge(intervals [][]int) [][]int {
	/*
	   Lets draw how things overlap:

	   |.......|
	      |.......|

	   |.............|
	      |....|

	   If they are sorted in non-decresing order on the begin of interval, then we can merge them one by one from left.

	   We will use two pointers (or linked list with next)
	   p1 looks at the current element (starts at 0)
	   p2 looks at the currently evaluated element to merge (starts at 1)
	     if p2 can be merge into p1 (p2.start <= p1.end) -> merge into p1 and move p2++
	     if not, record p1 to ans and p1 = p2, p2++
	   loop is till p2 < len(intervals)
	   record last element in p1 to ans

	*/

	if len(intervals) < 2 {
		return intervals
	}

	// Sort intervals
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	p1, p2 := 0, 1
	var ans [][]int
	for p2 < len(intervals) {
		if intervals[p2][0] <= intervals[p1][1] {
			if intervals[p2][1] > intervals[p1][1] {
				intervals[p1][1] = intervals[p2][1]
			}
			p2++
		} else {
			ans = append(ans, intervals[p1])
			p1 = p2
			p2++
		}
	}
	ans = append(ans, intervals[p1])
	return ans
}
