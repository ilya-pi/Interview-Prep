type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// Initialize sums in NumMatrix
	nm := NumMatrix{}
	nm.dp = make([][]int, len(matrix))
	for row, _ := range nm.dp {
		nm.dp[row] = make([]int, len(matrix[0]))
	}

	// Fill sums in NumMatrix
	for row, _ := range nm.dp {
		var sum int
		for col, _ := range nm.dp[row] {
			sum += matrix[row][col]
			if row > 0 {
				nm.dp[row][col] = sum + nm.dp[row-1][col]
			} else {
				nm.dp[row][col] = sum
			}
		}
	}

	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	/*

	   targetSum == dp[row2][col2] - dp[row2][col1-1] - dp[row1-1][col2] + dp[row1-1][col1-1]

	*/

	/*
	   for _, v := range this.dp {
	       fmt.Printf("%v\n", v)
	   }
	   fmt.Printf("\n")
	*/

	if row2 < row1 {
		return 0
	}
	if col2 < col1 {
		return 0
	}

	var a int
	var b int
	var c int
	if col1 > 0 {
		b = this.dp[row2][col1-1]
	}
	if row1 > 0 {
		c = this.dp[row1-1][col2]
	}
	if col1 > 0 && row1 > 0 {
		a = this.dp[row1-1][col1-1]
	}

	return this.dp[row2][col2] - b - c + a
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
