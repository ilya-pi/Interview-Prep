package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'crosswordPuzzle' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY crossword
 *  2. STRING words
 */

func crosswordPuzzle(crossword []string, words string) []string {
	/*
	   Approach:
	   With one word, fit it in all places it can be and try to fit the rest, if not - backtrack.

	   So;
	   1) Take 1st word of words, if empty -> return true, then return sequence
	   2) Fit in all positions we can -> 2 for loops one by row, one by column
	   3) For every fit -> call the self with less words using the same array, if can't return false
	*/

	table := make([][]rune, len(crossword))
	for i := 0; i < len(crossword); i++ {
		table[i] = []rune(crossword[i])
	}

	canFit := func(i, j int, wordS string, vertical bool) bool {
		word := []rune(wordS)
		if vertical {
			for k := 0; k < len(word); k++ {
				if j+k >= len(table) {
					return false
				}
				if !(table[i][j+k] == '-' || table[i][j+k] == word[k]) {
					return false
				}
			}
		}
		if !vertical {
			for k := 0; k < len(word); k++ {
				if i+k >= len(table) {
					return false
				}
				if !(table[i+k][j] == '-' || table[i+k][j] == word[k]) {
					return false
				}
			}
		}
		return true
	}
	fitIn := func(i, j int, wordS string, vertical bool) []rune {
		word := []rune(wordS)
		prev := make([]rune, len(word))
		if vertical {
			for k := 0; k < len(word) && j+k < len(table); k++ {
				prev[k] = table[i][j+k] // for backtracking
				table[i][j+k] = word[k]
			}
		}
		if !vertical {
			for k := 0; k < len(word) && i+k < len(table); k++ {
				prev[k] = table[i+k][j]
				table[i+k][j] = word[k]
			}
		}

		return prev
	}
	unFit := func(i, j int, prev []rune, vertical bool) {
		if vertical {
			for k := 0; k < len(prev); k++ {
				table[i][j+k] = prev[k]
			}
		}
		if !vertical {
			for k := 0; k < len(prev); k++ {
				table[i+k][j] = prev[k]
			}
		}
	}

	var fit func(words []string) bool
	fit = func(words []string) bool {
		if len(words) == 0 {
			return true
		}
		word := words[0]
		words = words[1:]

		for i := 0; i < len(crossword); i++ { // there are cases which we don't have to check in the corners
			for j := 0; j < len(crossword[0]); j++ {
				for _, direction := range []bool{true, false} {
					if canFit(i, j, word, direction) {
						prev := fitIn(i, j, word, direction)
						theRestFits := fit(words)
						if theRestFits {
							return true
						} else {
							unFit(i, j, prev, direction)
						}
					}
				}
			}
		}

		words = append([]string{word}, words...)
		return false
	}

	wordsArr := strings.Split(words, ";")
	if fit(wordsArr) {
		res := make([]string, len(crossword))
		for i := 0; i < len(crossword); i++ {
			res[i] = string(table[i])
		}
		return res
	}
	// If we can fit the thing, this shouldn't happen
	return crossword
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var crossword []string

	for i := 0; i < 10; i++ {
		crosswordItem := readLine(reader)
		crossword = append(crossword, crosswordItem)
	}

	words := readLine(reader)

	result := crosswordPuzzle(crossword, words)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
