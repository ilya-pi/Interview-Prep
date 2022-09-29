import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	/*
	   Check if it is ipv4 or ipv6 by splitting on ".", if not split -> ipv6, if not split on ":" -> neither.

	   for ipv4 check there are 4 elements and neither start from 0, unless length == 1, and that they are < 255

	   for ipv6, check that all runes in subsequences are between 0 - 9 a-z A-Z and len == 4
	*/

	split := strings.Split(queryIP, ".")
	if len(split) == 4 {
		// Try for IPv4
		for _, block := range split {
			rb := []rune(block)
			if len(rb) > 1 && rb[0] == '0' {
				return "Neither"
			}
			v, err := strconv.ParseInt(block, 10, 32)
			if err != nil || v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		split = strings.Split(queryIP, ":")
		if len(split) == 8 {
			// Try for IPv6
			for _, block := range split {
				rb := []rune(block)
				if len(rb) < 1 || len(rb) > 4 {
					return "Neither"
				}
				_, err := strconv.ParseInt(block, 16, 64)
				if err != nil {
					return "Neither"
				}
			}
			return "IPv6"
		}
	}
	return "Neither"

}
