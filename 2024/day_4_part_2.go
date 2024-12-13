package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func masDiagonal(matrix [][]string, rowAIdx int, colAIdx int) bool {
	return rowAIdx > 0 && colAIdx > 0 &&
		matrix[rowAIdx-1][colAIdx-1] == "M" &&
		rowAIdx < len(matrix)-1 && colAIdx < len(matrix[rowAIdx])-1 &&
		matrix[rowAIdx+1][colAIdx+1] == "S"
}

func masDiagonalRev(matrix [][]string, rowAIdx int, colAIdx int) bool {
	return rowAIdx > 0 && colAIdx > 0 &&
		matrix[rowAIdx-1][colAIdx-1] == "S" &&
		rowAIdx < len(matrix)-1 && colAIdx < len(matrix[rowAIdx])-1 &&
		matrix[rowAIdx+1][colAIdx+1] == "M"
}

func masAntidiag(matrix [][]string, rowAIdx int, colAIdx int) bool {
	return rowAIdx < len(matrix)-1 && colAIdx > 0 &&
		matrix[rowAIdx+1][colAIdx-1] == "M" &&
		rowAIdx > 0 && colAIdx < len(matrix[rowAIdx])-1 &&
		matrix[rowAIdx-1][colAIdx+1] == "S"
}

func masAntidiagRev(matrix [][]string, rowAIdx int, colAIdx int) bool {
	return rowAIdx < len(matrix)-1 && colAIdx > 0 &&
		matrix[rowAIdx+1][colAIdx-1] == "S" &&
		rowAIdx > 0 && colAIdx < len(matrix[rowAIdx])-1 &&
		matrix[rowAIdx-1][colAIdx+1] == "M"
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
			if inputMatrix[rowIdx][colIdx] == "A" {
				if (masDiagonal(inputMatrix, rowIdx, colIdx) ||
					masDiagonalRev(inputMatrix, rowIdx, colIdx)) &&
					(masAntidiag(inputMatrix, rowIdx, colIdx) ||
						masAntidiagRev(inputMatrix, rowIdx, colIdx)) {
					xmasCount += 1
				}
			}
		}
	}
	fmt.Println(xmasCount)
}
