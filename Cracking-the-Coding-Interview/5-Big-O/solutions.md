
VI.1 O(b)

VI.2 O(b)

VI.3 O(1)

VI.4 O(a/b)

VI.5 O(logn)

VI.6 O(sqrt(n))

VI.7 O(n)

VI.8 O(n)

VI.9 O(n * (sum 0 .. n-1 == n/2)) == O(n^2)

VI.10 O(len(n)) idiotto, yes it is logn

VI.11 O(k * 26^k) 

k, ""
  k - 1, "a"
      k - 2, "aa"
      k - 2, "ab"
      ..
      k - 2, "az"
  k - 1, "b"
  ..
  k - 1, "z"

-> 26^k

VI.12 O(b*logb + a*logb)
