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
	var adjacent []position
	adjacent = append(adjacent, position{pos.row - 1, pos.col - 1})
	adjacent = append(adjacent, position{pos.row - 1, pos.col})
	adjacent = append(adjacent, position{pos.row - 1, pos.col + 1})
	adjacent = append(adjacent, position{pos.row, pos.col - 1})
	adjacent = append(adjacent, position{pos.row, pos.col + 1})
	adjacent = append(adjacent, position{pos.row + 1, pos.col - 1})
	adjacent = append(adjacent, position{pos.row + 1, pos.col})
	adjacent = append(adjacent, position{pos.row + 1, pos.col + 1})
	return adjacent
}

func outOfBounds(pos position, rows int, cols int) bool {
	return pos.row < 0 || pos.row >= rows || pos.col < 0 || pos.col >= cols
}

func countAccessible(grid [][]bool, rolls map[position]bool) int {
	accessible := 0
	for roll := range rolls {
		adjacent := 0
		for _, pos := range getAdjacent(position{roll.row, roll.col}) {
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
			grid[roll.row][roll.col] = false
			delete(rolls, roll)
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
	rolls := map[position]bool{}
	for inputScanner.Scan() {
		var row []bool
		for colIdx, symbol := range inputScanner.Text() {
			row = append(row, symbol == '@')
			if row[len(row)-1] {
				rolls[position{len(grid), colIdx}] = true
			}
		}
		grid = append(grid, row)
	}
	result, accessible := 0, countAccessible(grid, rolls)
	for accessible > 0 {
		result += accessible
		accessible = countAccessible(grid, rolls)
	}
	fmt.Println(result)
}
