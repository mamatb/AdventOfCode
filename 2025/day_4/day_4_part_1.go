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

func countRollsAccessible(grid [][]bool) int {
	rollsAccessible := 0
	for rowIdx, row := range grid {
		for colIdx, roll := range row {
			if !roll {
				continue
			}
			rollsAdjacent := 0
			for _, pos := range getAdjacent(position{rowIdx, colIdx}) {
				if !outOfBounds(pos, len(grid), len(grid[0])) &&
					grid[pos.row][pos.col] {
					rollsAdjacent++
				}
				if rollsAdjacent >= 4 {
					break
				}
			}
			if rollsAdjacent < 4 {
				rollsAccessible++
			}
		}
	}
	return rollsAccessible
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
	fmt.Println(countRollsAccessible(grid))
}
