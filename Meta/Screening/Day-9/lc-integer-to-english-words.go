import "strings"

func numberToWords(num int) string {
	/*
	   2^32 ~= 10^10

	   1 - 9
	   One
	   Two
	   Three
	   Four
	   Five
	   Six
	   Seven
	   Eight
	   Nine


	   10 - Ten
	   11 - Eleven
	   12 - Twelve
	   13 - Thirteen
	   14 - Fourteen
	   15 - Fifteen
	   16 - Sixteen
	   17 - Seventeen
	   18 - Eightteen
	   19 - Nineteen

	   20 - Twenty
	   30 - Thirty
	   40 - Fourty
	   50 - Fifty
	   60 - Sixty
	   70 - Seventy
	   80 - Eighty
	   90 - Ninety

	   100 Hundred
	   200 - Two Hundred ...

	   1000 - Thousand
	   2000 - Two Thousand

	   12 000 Twelve Thousand

	   1 000 000 One Million
	   2 ... Two Million

	   999 000 000 Nine Hundred Ninety Nine Million
	   1 000 000 000 One Billion

	   abc def ghi jkl

	   abc => toEng() + Billion
	   def => toEnd() + Million
	   ghi => toEng() + Thousand
	   jkl => toEng()

	   toEng() []string

	   ans []string

	   len(s) / 3 == blocks
	   if len(s) > blocks *3 => deal with top and then the rest

	   toEng() -> have hundreds? and !=0 -> numToEng(1..9) + "Hundred"
	   2-nd is 0 -> continue
	   2-nd is 1-> special cases
	   2-nd 2..9 -> twenty..ninety + numToEng(1..9)

	   return ans.Join(" ")

	*/

	// Also potential validation for 2^31 boundaries

	if num == 0 {
		return "Zero"
	}

	numToEng := func(n int) string {
		if n < 1 || n > 9 {
			return ""
		}
		switch n {
		case 1:
			return "One"
		case 2:
			return "Two"
		case 3:
			return "Three"
		case 4:
			return "Four"
		case 5:
			return "Five"
		case 6:
			return "Six"
		case 7:
			return "Seven"
		case 8:
			return "Eight"
		case 9:
			return "Nine"
		}
		return ""
	}

	threeToEng := func(n int) []string {
		/*
		   1 - 9
		   One
		   Two
		   Three
		   Four
		   Five
		   Six
		   Seven
		   Eight
		   Nine


		   10 - Ten
		   11 - Eleven
		   12 - Twelve
		   13 - Thirteen
		   14 - Fourteen
		   15 - Fifteen
		   16 - Sixteen
		   17 - Seventeen
		   18 - Eightteen
		   19 - Nineteen

		   20 - Twenty
		   30 - Thirty
		   40 - Fourty
		   50 - Fifty
		   60 - Sixty
		   70 - Seventy
		   80 - Eighty
		   90 - Ninety

		   100 Hundred
		   200 - Two Hundred ...
		*/

		var acc []string

		hundreds := n / 100
		n %= 100
		if hundreds > 0 {
			acc = append(acc, numToEng(hundreds))
			acc = append(acc, "Hundred")
		}

		if n > 0 && n < 10 {
			acc = append(acc, numToEng(n))
		} else if n >= 10 && n < 20 {
			var str string
			switch n {
			case 10:
				str = "Ten"
			case 11:
				str = "Eleven"
			case 12:
				str = "Twelve"
			case 13:
				str = "Thirteen"
			case 14:
				str = "Fourteen"
			case 15:
				str = "Fifteen"
			case 16:
				str = "Sixteen"
			case 17:
				str = "Seventeen"
			case 18:
				str = "Eighteen"
			case 19:
				str = "Nineteen"
			}
			acc = append(acc, str)
		} else if n >= 20 {
			by10 := n / 10
			var str string
			switch by10 {
			case 2:
				str = "Twenty"
			case 3:
				str = "Thirty"
			case 4:
				str = "Forty"
			case 5:
				str = "Fifty"
			case 6:
				str = "Sixty"
			case 7:
				str = "Seventy"
			case 8:
				str = "Eighty"
			case 9:
				str = "Ninety"
			}
			acc = append(acc, str)
			left := n % 10
			if left > 0 {
				acc = append(acc, numToEng(left))
			}
		}
		return acc
	}

	var acc [][]string
	blocksMap := map[int]string{1: "Thousand", 2: "Million", 3: "Billion"}
	block := 0
	for num != 0 {
		last3 := num % 1000
		num /= 1000
		cur := threeToEng(last3)
		if blockName, ok := blocksMap[block]; ok && len(cur) > 0 {
			cur = append(cur, blockName)
		}
		acc = append(acc, cur)
		block++
	}

	// reverse acc and join
	var ans []string
	for i := len(acc) - 1; i >= 0; i-- {
		ans = append(ans, acc[i]...)
	}
	return strings.Join(ans, " ")

}
