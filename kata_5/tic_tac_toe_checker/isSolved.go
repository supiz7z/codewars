package tic_tac_toe_checker

import "fmt"

func HasZeros(arr [3][3]int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func MatrixCheck(arr [3][3]int) int {
	for i := 0; i < len(arr); i++ {
		if RowCheck(arr[i]) == 1 {
			return 1
		}
		if RowCheck(arr[i]) == 2 {
			return 2
		}
	}
	return 0
}

func RowCheck(row [3]int) int {
	if row[0] == 1 && row[1] == 1 && row[2] == 1 {
		return 1
	}
	if row[0] == 2 && row[1] == 2 && row[2] == 2 {
		return 2
	}
	return 0
}

func IsSolved(board [3][3]int) int {
	n := len(board)

	diag := [3]int{}
	for i := 0; i < n; i++ {
		diag[i] = board[i][i]
	}
	resDiag1 := RowCheck(diag)
	for i := 0; i < n; i++ {
		diag[i] = board[i][n-1-i]
	}
	resDiag2 := RowCheck(diag)

	rows := MatrixCheck(board)
	columnArr := [3][3]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			columnArr[i][j] = board[j][i]
		}
	}
	columns := MatrixCheck(columnArr)
	if rows == 1 || columns == 1 || resDiag1 == 1 || resDiag2 == 1 {
		return 1
	}
	if rows == 2 || columns == 2 || resDiag1 == 2 || resDiag2 == 2 {
		return 2
	}
	if HasZeros(board) {
		return -1
	}
	return 0
}

func Start() {
	//board := [3][3]int{ //not finished (-1)
	//	{0, 0, 1},
	//	{0, 1, 2},
	//	{2, 1, 0},
	//}
	//board := [3][3]int{ //win "1" in row (1)
	//	{1, 1, 1},
	//	{0, 2, 2},
	//	{0, 0, 0},
	//}
	//board := [3][3]int{ //win "1" in column (1)
	//	{2, 1, 2},
	//	{2, 1, 1},
	//	{1, 1, 2},
	//}
	board := [3][3]int{ //draw (0)
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 1},
	}
	result := IsSolved(board)
	fmt.Println(result)
}
