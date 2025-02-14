package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

func xmasHorizontal(posX position, matrix [][]string) bool {
	return posX.col < len(matrix[posX.row])-3 &&
		matrix[posX.row][posX.col+1] == "M" &&
		matrix[posX.row][posX.col+2] == "A" &&
		matrix[posX.row][posX.col+3] == "S"
}

func xmasHorizontalRev(posX position, matrix [][]string) bool {
	return posX.col > 2 &&
		matrix[posX.row][posX.col-1] == "M" &&
		matrix[posX.row][posX.col-2] == "A" &&
		matrix[posX.row][posX.col-3] == "S"
}

func xmasVertical(posX position, matrix [][]string) bool {
	return posX.row < len(matrix)-3 &&
		matrix[posX.row+1][posX.col] == "M" &&
		matrix[posX.row+2][posX.col] == "A" &&
		matrix[posX.row+3][posX.col] == "S"
}

func xmasVerticalRev(posX position, matrix [][]string) bool {
	return posX.row > 2 &&
		matrix[posX.row-1][posX.col] == "M" &&
		matrix[posX.row-2][posX.col] == "A" &&
		matrix[posX.row-3][posX.col] == "S"
}

func xmasDiagonal(posX position, matrix [][]string) bool {
	return posX.col < len(matrix[posX.row])-3 && posX.row < len(matrix)-3 &&
		matrix[posX.row+1][posX.col+1] == "M" &&
		matrix[posX.row+2][posX.col+2] == "A" &&
		matrix[posX.row+3][posX.col+3] == "S"
}

func xmasDiagonalRev(posX position, matrix [][]string) bool {
	return posX.col > 2 && posX.row > 2 &&
		matrix[posX.row-1][posX.col-1] == "M" &&
		matrix[posX.row-2][posX.col-2] == "A" &&
		matrix[posX.row-3][posX.col-3] == "S"
}

func xmasAntidiag(posX position, matrix [][]string) bool {
	return posX.col < len(matrix[posX.row])-3 && posX.row > 2 &&
		matrix[posX.row-1][posX.col+1] == "M" &&
		matrix[posX.row-2][posX.col+2] == "A" &&
		matrix[posX.row-3][posX.col+3] == "S"
}

func xmasAntidiagRev(posX position, matrix [][]string) bool {
	return posX.col > 2 && posX.row < len(matrix)-3 &&
		matrix[posX.row+1][posX.col-1] == "M" &&
		matrix[posX.row+2][posX.col-2] == "A" &&
		matrix[posX.row+3][posX.col-3] == "S"
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_4.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	wordSearch, row, positionsX, xmasCount := [][]string{}, 0, []position{}, 0
	for inputScanner.Scan() {
		wordSearch = append(wordSearch, []string{})
		for col, letter := range strings.Split(inputScanner.Text(), "") {
			wordSearch[row] = append(wordSearch[row], letter)
			if letter == "X" {
				positionsX = append(positionsX, position{row: row, col: col})
			}
		}
		row++
	}
	for _, posX := range positionsX {
		if xmasHorizontal(posX, wordSearch) {
			xmasCount++
		}
		if xmasHorizontalRev(posX, wordSearch) {
			xmasCount++
		}
		if xmasVertical(posX, wordSearch) {
			xmasCount++
		}
		if xmasVerticalRev(posX, wordSearch) {
			xmasCount++
		}
		if xmasDiagonal(posX, wordSearch) {
			xmasCount++
		}
		if xmasDiagonalRev(posX, wordSearch) {
			xmasCount++
		}
		if xmasAntidiag(posX, wordSearch) {
			xmasCount++
		}
		if xmasAntidiagRev(posX, wordSearch) {
			xmasCount++
		}
	}
	fmt.Println(xmasCount)
}
