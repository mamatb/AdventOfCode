package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row int
	col int
}

func getAdjacent(pos position) []position {
	return []position{
		{row: pos.row - 1, col: pos.col - 1},
		{row: pos.row - 1, col: pos.col},
		{row: pos.row - 1, col: pos.col + 1},
		{row: pos.row, col: pos.col - 1},
		{row: pos.row, col: pos.col + 1},
		{row: pos.row + 1, col: pos.col - 1},
		{row: pos.row + 1, col: pos.col},
		{row: pos.row + 1, col: pos.col + 1},
	}
}

func outOfBounds(pos position, rows int, cols int) bool {
	return pos.row < 0 || pos.row >= rows || pos.col < 0 || pos.col >= cols
}

func countAccessible(grid [][]bool) int {
	accessible := 0
	for rowIdx, row := range grid {
		for colIdx, roll := range row {
			if !roll {
				continue
			}
			adjacent := 0
			for _, pos := range getAdjacent(position{row: rowIdx, col: colIdx}) {
				if !outOfBounds(pos, len(grid), len(grid[0])) &&
					grid[pos.row][pos.col] {
					adjacent++
				}
				if adjacent >= 4 {
					break
				}
			}
			if adjacent < 4 {
				accessible++
			}
		}
	}
	return accessible
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_4.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	var grid [][]bool
	for inputScanner.Scan() {
		var row []bool
		for _, symbol := range inputScanner.Text() {
			row = append(row, symbol == '@')
		}
		grid = append(grid, row)
	}
	fmt.Println(countAccessible(grid))
}
