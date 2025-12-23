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

func countAccessible(rolls map[position]bool) int {
	accessible := 0
	for roll := range rolls {
		adjacent := 0
		for _, pos := range getAdjacent(roll) {
			if rolls[pos] {
				adjacent++
			}
			if adjacent >= 4 {
				break
			}
		}
		if adjacent < 4 {
			accessible++
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
	rowIdx, rolls := 0, map[position]bool{}
	for inputScanner.Scan() {
		for colIdx, symbol := range inputScanner.Text() {
			if symbol == '@' {
				rolls[position{row: rowIdx, col: colIdx}] = true
			}
		}
		rowIdx++
	}
	result, accessible := 0, countAccessible(rolls)
	for accessible > 0 {
		result += accessible
		accessible = countAccessible(rolls)
	}
	fmt.Println(result)
}
