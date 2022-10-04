/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	/*

	   So the badness of version goes like this:

	   0 0 0 0 0 0 0 0 1 1 1 1 1 1 1

	   We need to find first 1

	   binary search:

	   l r
	   m = (l+r) / 2
	   if mid is good -> look on the right mid+1
	   if mid is bad -> look on thh left mid-1

	   0 0 0 0 0 0 0 _0_ 1 1 1 1 1 1 1
	   1 1 1 _1_ 1 1 1
	   1 _1_ 1
	   1 ! l < l -> outcome

	   0 0 _1_ 1
	   0 _0_
	   mid + 1 -> l is outcome


	*/

	/*
	   0 1 1
	   1 2 3
	     2
	   1 1
	*/

	l, r := 1, n
	for l <= r {
		m := (l + r) / 2
		if isBadVersion(m) {
			l, r = l, m-1
		} else {
			l, r = m+1, r
		}
	}
	return l

}
