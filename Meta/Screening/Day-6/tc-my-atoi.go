import "math"

func myAtoi(s string) int {
	// 2147483648
	// '  -0032'
	var ans int64
	var negative bool
	var signSet bool
	var readingCifers bool
loop:
	for _, v := range []rune(s) {
		switch {
		case v == ' ' && !readingCifers:
			continue loop
		case v == ' ' && readingCifers:
			break loop
		case v == '+' && !signSet:
			signSet = true
			readingCifers = true
		case v == '+' && signSet:
			break loop
		case v == '-' && !signSet:
			signSet = true
			negative = true
		case v == '-' && signSet:
			break loop
		case v == '0' && !readingCifers:
			signSet = true
			readingCifers = true
			continue loop
		case v == '0' && readingCifers:
			ans *= 10
		case v >= '1' && v <= '9':
			signSet = true
			readingCifers = true
			ans *= 10
			ans += int64(v-'1') + 1
		default:
			break loop
		}
		if ans > int64(math.MaxInt32) {
			if negative {
				return int(math.MinInt32)
			} else {
				return int(math.MaxInt32)
			}
		}
	}
	if negative {
		return -int(ans)
	} else {
		return int(ans)
	}
}
