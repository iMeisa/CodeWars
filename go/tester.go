package main

import (
	"CodeWars/ticTacToeChecker"
	"fmt"
)

func main() {
	fmt.Println(ticTacToeChecker.IsSolved([3][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{2, 1, 0},
	}))
	fmt.Println(ticTacToeChecker.IsSolved([3][3]int{
		{1, 1, 1},
		{0, 2, 2},
		{0, 0, 0},
	}))
	fmt.Println(ticTacToeChecker.IsSolved([3][3]int{
		{1, 2, 1},
		{2, 2, 1},
		{1, 2, 2},
	}))
	fmt.Println(ticTacToeChecker.IsSolved([3][3]int{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 1},
	}))
}
