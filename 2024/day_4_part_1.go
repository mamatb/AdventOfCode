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

func xmasHorizontal(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexColX < len(matrix[indexRowX])-3 &&
		matrix[indexRowX][indexColX+1] == "M" &&
		matrix[indexRowX][indexColX+2] == "A" &&
		matrix[indexRowX][indexColX+3] == "S")
}

func xmasHorizontalRev(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexColX > 2 &&
		matrix[indexRowX][indexColX-1] == "M" &&
		matrix[indexRowX][indexColX-2] == "A" &&
		matrix[indexRowX][indexColX-3] == "S")
}

func xmasVertical(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexRowX < len(matrix)-3 &&
		matrix[indexRowX+1][indexColX] == "M" &&
		matrix[indexRowX+2][indexColX] == "A" &&
		matrix[indexRowX+3][indexColX] == "S")
}

func xmasVerticalRev(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexRowX > 2 &&
		matrix[indexRowX-1][indexColX] == "M" &&
		matrix[indexRowX-2][indexColX] == "A" &&
		matrix[indexRowX-3][indexColX] == "S")
}

func xmasDiagonal(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexColX < len(matrix[indexRowX])-3 && indexRowX < len(matrix)-3 &&
		matrix[indexRowX+1][indexColX+1] == "M" &&
		matrix[indexRowX+2][indexColX+2] == "A" &&
		matrix[indexRowX+3][indexColX+3] == "S")
}

func xmasDiagonalRev(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexColX > 2 && indexRowX > 2 &&
		matrix[indexRowX-1][indexColX-1] == "M" &&
		matrix[indexRowX-2][indexColX-2] == "A" &&
		matrix[indexRowX-3][indexColX-3] == "S")
}

func xmasAntidiag(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexColX < len(matrix[indexRowX])-3 && indexRowX > 2 &&
		matrix[indexRowX-1][indexColX+1] == "M" &&
		matrix[indexRowX-2][indexColX+2] == "A" &&
		matrix[indexRowX-3][indexColX+3] == "S")
}

func xmasAntidiagRev(matrix [][]string, indexRowX int, indexColX int) bool {
	return (indexColX > 2 && indexRowX < len(matrix)-3 &&
		matrix[indexRowX+1][indexColX-1] == "M" &&
		matrix[indexRowX+2][indexColX-2] == "A" &&
		matrix[indexRowX+3][indexColX-3] == "S")
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
