package main

import "fmt"

func main() {
	var board = [][]int{
		{0, 0, 8, 0, 7, 1, 0, 5, 0},
		{0, 0, 0, 0, 0, 4, 0, 0, 0},
		{7, 1, 0, 2, 0, 0, 3, 0, 0},
		{0, 7, 5, 0, 0, 0, 0, 8, 0},
		{1, 0, 0, 0, 3, 0, 6, 2, 0},
		{0, 2, 3, 0, 0, 0, 0, 1, 0},
		{8, 6, 0, 4, 0, 0, 9, 0, 0},
		{0, 0, 0, 0, 0, 8, 0, 0, 0},
		{0, 0, 7, 0, 2, 9, 0, 3, 0},
	}
	sudoku(board)
	OutputSudoku(board)
}
func sudoku(board [][]int) bool {
	//遍历行和列
	for i := 0; i < 9; i++ { //行
		for j := 0; j < 9; j++ { //列
			if board[i][j] != 0 { //跳过原始数字
				continue
			}
			for k := 1; k <= 9; k++ {
				if backT(i, j, k, board) == true { //放置k是否可以
					board[i][j] = k
					if sudoku(board) == true {
						return true //满足校验条件返回
					}
					board[i][j] = 0 //不满足回溯为0
				}
			}
			return false //返回
		}
	}
	return true
}

func backT(row int, col int, val int, board [][]int) bool {
	//同行是否重复
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}
	//同列是否重复
	for j := 0; j < 9; j++ {
		if board[j][col] == val {
			return false
		}
	}
	//9宫格是否重复
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for t := startRow; t < startRow+3; t++ {
		for x := startCol; x < startCol+3; x++ {
			if board[t][x] == val {
				return false
			}
		}
	}
	return true //都满足返回true
}
func OutputSudoku(sudoku [][]int) {
	for i := 0; i < 81; i++ {
		fmt.Printf("%2d ", sudoku[i/9][i%9])
		if i != 0 && (i+1)%9 == 0 {
			fmt.Println("")
		}
	}
}
