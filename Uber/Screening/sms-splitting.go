package main

import (
	"fmt"
	"strings"
)

func split(str string, maxBlockSize int) []string {
	/*
			The problem in this is that we don't know how many segments there would be.

			Thus we will try something akin to a greedy approach with backtracking.

			Edge case: str <= size -> return str
			Edge case 2: We cannot attach a word to aggregation within given limitations at all

			Recursion could look like this:

			spl (words, pos, sz, i, aggr, ans) int
			    if i == sz
			      ans = [:0]
			      pos = 0
			      sz *= 10
			      i = 1
			      return spl(...)
			    if pos == len(words)
			        aggr <- attach(i/%d)
				ans <- append aggr
			    	return i
			    if we can take next word and it fits with postfix size ->
		              pos++
			      aggr += " " + words[pos]
			      return spl(words, pos+1, sz, i, aggr, ans)
			    else
			      pos++
			      aggr <- attach (i/%d)
			      ans <- append aggr
			      return spl(words, pos, sz, i+1, "", ans)
	*/

	// Do we need to split?
	if len(str) <= maxBlockSize {
		return []string{str}
	}

	// We do
	var spl func([]string, int, int, int, *strings.Builder, *[]string) int
	spl = func(words []string, pos, sz, i int, aggr *strings.Builder, ans *[]string) int {
		if i == sz {
			// We cannot complete split with current amount of blocks
			*ans = (*ans)[:0]
			aggr.Reset()
			return spl(words, 0, sz*10, 1, aggr, ans)
		}
		if pos == len(words) {
			// Added all words!
			aggr.WriteString(fmt.Sprintf("(%d/", i))
			aggr.WriteString("%d)")
			*ans = append(*ans, aggr.String())
			return i
		}
		// sz will always be 10, 100, 1000 and etc, so amount of meessages will be max 9, 99, 999 and etc
		// todo(ilya): compute it only when necessary (pass it in the parametres)
		postfixSize := len(fmt.Sprintf("(%d/%d)", i, sz-1))
		if aggr.Len()+1+len(words[pos])+postfixSize <= maxBlockSize {
			// Still can add a word
			aggr.WriteByte(' ')
			aggr.WriteString(words[pos])
			return spl(words, pos+1, sz, i, aggr, ans)
		} else {
			// Close previous block and begin new
			aggr.WriteString(fmt.Sprintf("(%d/", i))
			aggr.WriteString("%d)")
			*ans = append(*ans, aggr.String())
			aggr.Reset()
			return spl(words, pos, sz, i+1, aggr, ans)
		}
	}
	var b strings.Builder
	var ans []string
	words := strings.Split(str, " ")
	total := spl(words, 0, 10, 1, &b, &ans)
	for i, v := range ans {
		ans[i] = fmt.Sprintf(v, total)
	}
	return ans
}

func main() {
	str := `
Tell us a little bit about yourself
I am an aerospace engineer who became a data scientist by accident before the title was popularized and data sets became bigger, faster and more complex. I lived through the (often painful) evolution from data warehouses, to on-prem data lakes, and finally to the cloud while working for a large enterprise. I was a user of the Databricks Unified Data Analytics Platform prior to joining the company and became a believer in the product vision since it solved so many fundamental problems in big data. I’ve been with Databricks for one year as a leader in our Customer Success organization, based in our Midtown Manhattan office. I live in Brooklyn with my wife, baby girl and dog, enjoying the NYC lifestyle.
What were you looking for in your next opportunity, and why did you choose Databricks?
Databricks was a step-change for me both in technology and career growth. I lead a team called Resident Solutions Architects in our Customer Success organization. This is a diverse group of customer-facing big data architects who work across the data engineering and machine learning spectrum to solve our customers’ most challenging use-cases in the field. The variety and complexity of the problems we solve is unparalleled. We are learning and adding value in every business vertical and feeding those insights back into the product continuously. We design massively scalable solutions, write production-quality code and become trusted advisors to the data teams of the most successful enterprises in the world.
`
	res := split(str, 160)
	for _, v := range res {
		fmt.Printf("len == %d -> %s\n", len(v), v)
	}
}
