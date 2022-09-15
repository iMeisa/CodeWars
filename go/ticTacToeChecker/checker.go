package ticTacToeChecker

func IsSolved(board [3][3]int) int {

	var yRows [3][3]int
	diagRows := [2][3]int{
		{board[0][0], board[1][1], board[2][2]},
		{board[0][2], board[1][1], board[2][0]},
	}

	incomplete := false

	// Check X row
	for x, row := range board {
		if isWin(row) {
			return row[0]
		}

		// Make Y rows
		for y := range board {
			if board[x][y] == 0 {
				incomplete = true
			}
			yRows[y][x] = board[x][y]
		}
	}

	// Check Y rows
	for _, row := range yRows {
		if isWin(row) {
			return row[0]
		}
	}

	// Check diag rows
	for _, row := range diagRows {
		if isWin(row) {
			return row[0]
		}
	}

	if incomplete {
		return -1
	}

	return 0
}

func isWin(a [3]int) bool {
	v := a[0]
	if v == 0 {
		return false
	}

	for _, i := range a[1:] {
		if i != v {
			return false
		}
	}

	return true
}
