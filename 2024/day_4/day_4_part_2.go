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

func masDiagonal(posA position, matrix [][]string) bool {
	return posA.row > 0 && posA.col > 0 &&
		matrix[posA.row-1][posA.col-1] == "M" &&
		posA.row < len(matrix)-1 && posA.col < len(matrix[posA.row])-1 &&
		matrix[posA.row+1][posA.col+1] == "S"
}

func masDiagonalRev(posA position, matrix [][]string) bool {
	return posA.row > 0 && posA.col > 0 &&
		matrix[posA.row-1][posA.col-1] == "S" &&
		posA.row < len(matrix)-1 && posA.col < len(matrix[posA.row])-1 &&
		matrix[posA.row+1][posA.col+1] == "M"
}

func masAntidiag(posA position, matrix [][]string) bool {
	return posA.row < len(matrix)-1 && posA.col > 0 &&
		matrix[posA.row+1][posA.col-1] == "M" &&
		posA.row > 0 && posA.col < len(matrix[posA.row])-1 &&
		matrix[posA.row-1][posA.col+1] == "S"
}

func masAntidiagRev(posA position, matrix [][]string) bool {
	return posA.row < len(matrix)-1 && posA.col > 0 &&
		matrix[posA.row+1][posA.col-1] == "S" &&
		posA.row > 0 && posA.col < len(matrix[posA.row])-1 &&
		matrix[posA.row-1][posA.col+1] == "M"
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_4.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	wordSearch, row, positionsA, xmasCount := [][]string{}, 0, []position{}, 0
	for inputScanner.Scan() {
		wordSearch = append(wordSearch, []string{})
		for col, letter := range strings.Split(inputScanner.Text(), "") {
			wordSearch[row] = append(wordSearch[row], letter)
			if letter == "A" {
				positionsA = append(positionsA, position{row: row, col: col})
			}
		}
		row++
	}
	for _, posA := range positionsA {
		if (masDiagonal(posA, wordSearch) || masDiagonalRev(posA, wordSearch)) &&
			(masAntidiag(posA, wordSearch) || masAntidiagRev(posA, wordSearch)) {
			xmasCount++
		}
	}
	fmt.Println(xmasCount)
}
