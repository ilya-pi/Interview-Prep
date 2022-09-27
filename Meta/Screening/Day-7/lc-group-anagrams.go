import "sort"

func groupAnagrams(strs []string) [][]string {
	/*
	   The brute force approach would be to generate key for each string and save it under that key into a map and then convert map values to string slice

	*/

	toKey := func(s string) string {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		return string(runes)
	}

	acc := make(map[string][]string)
	for _, str := range strs {
		key := toKey(str)
		acc[key] = append(acc[key], str)
	}

	var ans [][]string
	for _, v := range acc {
		ans = append(ans, v)
	}
	return ans
}
