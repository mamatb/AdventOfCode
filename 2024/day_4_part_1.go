package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func xmasHorizontal(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return colXIdx < len(matrix[rowXIdx])-3 &&
		matrix[rowXIdx][colXIdx+1] == "M" &&
		matrix[rowXIdx][colXIdx+2] == "A" &&
		matrix[rowXIdx][colXIdx+3] == "S"
}

func xmasHorizontalRev(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return colXIdx > 2 &&
		matrix[rowXIdx][colXIdx-1] == "M" &&
		matrix[rowXIdx][colXIdx-2] == "A" &&
		matrix[rowXIdx][colXIdx-3] == "S"
}

func xmasVertical(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return rowXIdx < len(matrix)-3 &&
		matrix[rowXIdx+1][colXIdx] == "M" &&
		matrix[rowXIdx+2][colXIdx] == "A" &&
		matrix[rowXIdx+3][colXIdx] == "S"
}

func xmasVerticalRev(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return rowXIdx > 2 &&
		matrix[rowXIdx-1][colXIdx] == "M" &&
		matrix[rowXIdx-2][colXIdx] == "A" &&
		matrix[rowXIdx-3][colXIdx] == "S"
}

func xmasDiagonal(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return colXIdx < len(matrix[rowXIdx])-3 && rowXIdx < len(matrix)-3 &&
		matrix[rowXIdx+1][colXIdx+1] == "M" &&
		matrix[rowXIdx+2][colXIdx+2] == "A" &&
		matrix[rowXIdx+3][colXIdx+3] == "S"
}

func xmasDiagonalRev(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return colXIdx > 2 && rowXIdx > 2 &&
		matrix[rowXIdx-1][colXIdx-1] == "M" &&
		matrix[rowXIdx-2][colXIdx-2] == "A" &&
		matrix[rowXIdx-3][colXIdx-3] == "S"
}

func xmasAntidiag(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return colXIdx < len(matrix[rowXIdx])-3 && rowXIdx > 2 &&
		matrix[rowXIdx-1][colXIdx+1] == "M" &&
		matrix[rowXIdx-2][colXIdx+2] == "A" &&
		matrix[rowXIdx-3][colXIdx+3] == "S"
}

func xmasAntidiagRev(matrix [][]string, rowXIdx int, colXIdx int) bool {
	return colXIdx > 2 && rowXIdx < len(matrix)-3 &&
		matrix[rowXIdx+1][colXIdx-1] == "M" &&
		matrix[rowXIdx+2][colXIdx-2] == "A" &&
		matrix[rowXIdx+3][colXIdx-3] == "S"
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_4.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	inputMatrix, xmasCount := [][]string{}, 0
	for inputScanner.Scan() {
		inputMatrix = append(inputMatrix, strings.Split(inputScanner.Text(), ""))
	}
	for rowIdx := range inputMatrix {
		for colIdx := range inputMatrix[rowIdx] {
			if inputMatrix[rowIdx][colIdx] == "X" {
				if xmasHorizontal(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasHorizontalRev(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasVertical(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasVerticalRev(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasDiagonal(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasDiagonalRev(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasAntidiag(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
				if xmasAntidiagRev(inputMatrix, rowIdx, colIdx) {
					xmasCount += 1
				}
			}
		}
	}
	fmt.Println(xmasCount)
}
