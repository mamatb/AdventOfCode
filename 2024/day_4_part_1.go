package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func xmasHorizontal(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexCol < len(matrix[indexRow])-3 &&
		matrix[indexRow][indexCol+1] == "M" &&
		matrix[indexRow][indexCol+2] == "A" &&
		matrix[indexRow][indexCol+3] == "S")
}

func xmasHorizontalRev(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexCol > 2 &&
		matrix[indexRow][indexCol-1] == "M" &&
		matrix[indexRow][indexCol-2] == "A" &&
		matrix[indexRow][indexCol-3] == "S")
}

func xmasVertical(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexRow < len(matrix)-3 &&
		matrix[indexRow+1][indexCol] == "M" &&
		matrix[indexRow+2][indexCol] == "A" &&
		matrix[indexRow+3][indexCol] == "S")
}

func xmasVerticalRev(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexRow > 2 &&
		matrix[indexRow-1][indexCol] == "M" &&
		matrix[indexRow-2][indexCol] == "A" &&
		matrix[indexRow-3][indexCol] == "S")
}

func xmasDiagonal(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexCol < len(matrix[indexRow])-3 && indexRow < len(matrix)-3 &&
		matrix[indexRow+1][indexCol+1] == "M" &&
		matrix[indexRow+2][indexCol+2] == "A" &&
		matrix[indexRow+3][indexCol+3] == "S")
}

func xmasDiagonalRev(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexCol > 2 && indexRow > 2 &&
		matrix[indexRow-1][indexCol-1] == "M" &&
		matrix[indexRow-2][indexCol-2] == "A" &&
		matrix[indexRow-3][indexCol-3] == "S")
}

func xmasAntidiag(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexCol < len(matrix[indexRow])-3 && indexRow > 2 &&
		matrix[indexRow-1][indexCol+1] == "M" &&
		matrix[indexRow-2][indexCol+2] == "A" &&
		matrix[indexRow-3][indexCol+3] == "S")
}

func xmasAntidiagRev(matrix [][]string, indexRow int, indexCol int) bool {
	return (indexCol > 2 && indexRow < len(matrix)-3 &&
		matrix[indexRow+1][indexCol-1] == "M" &&
		matrix[indexRow+2][indexCol-2] == "A" &&
		matrix[indexRow+3][indexCol-3] == "S")
}

func main() {
	input, err := os.Open("day_4.txt")
	errCheck(err)
	defer input.Close()
	inputMatrix, xmasCount := [][]string{}, 0
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		inputMatrix = append(inputMatrix, strings.Split(inputScanner.Text(), ""))
	}
	for indexRow := range inputMatrix {
		for indexCol := range inputMatrix[indexRow] {
			if inputMatrix[indexRow][indexCol] == "X" {
				if xmasHorizontal(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasHorizontalRev(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasVertical(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasVerticalRev(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasDiagonal(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasDiagonalRev(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasAntidiag(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
				if xmasAntidiagRev(inputMatrix, indexRow, indexCol) {
					xmasCount += 1
				}
			}
		}
	}
	fmt.Println(xmasCount)
}
