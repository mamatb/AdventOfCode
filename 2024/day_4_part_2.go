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

func masDiagonal(matrix [][]string, indexRowA int, indexColA int) bool {
	return (indexRowA > 0 && indexColA > 0 &&
		matrix[indexRowA-1][indexColA-1] == "M" &&
		indexRowA < len(matrix)-1 && indexColA < len(matrix[indexRowA])-1 &&
		matrix[indexRowA+1][indexColA+1] == "S")
}

func masDiagonalRev(matrix [][]string, indexRowA int, indexColA int) bool {
	return (indexRowA > 0 && indexColA > 0 &&
		matrix[indexRowA-1][indexColA-1] == "S" &&
		indexRowA < len(matrix)-1 && indexColA < len(matrix[indexRowA])-1 &&
		matrix[indexRowA+1][indexColA+1] == "M")
}

func masAntidiag(matrix [][]string, indexRowA int, indexColA int) bool {
	return (indexRowA < len(matrix)-1 && indexColA > 0 &&
		matrix[indexRowA+1][indexColA-1] == "M" &&
		indexRowA > 0 && indexColA < len(matrix[indexRowA])-1 &&
		matrix[indexRowA-1][indexColA+1] == "S")
}

func masAntidiagRev(matrix [][]string, indexRowA int, indexColA int) bool {
	return (indexRowA < len(matrix)-1 && indexColA > 0 &&
		matrix[indexRowA+1][indexColA-1] == "S" &&
		indexRowA > 0 && indexColA < len(matrix[indexRowA])-1 &&
		matrix[indexRowA-1][indexColA+1] == "M")
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
			if inputMatrix[indexRow][indexCol] == "A" {
				if (masDiagonal(inputMatrix, indexRow, indexCol) ||
					masDiagonalRev(inputMatrix, indexRow, indexCol)) &&
					(masAntidiag(inputMatrix, indexRow, indexCol) ||
						masAntidiagRev(inputMatrix, indexRow, indexCol)) {
					xmasCount += 1
				}
			}
		}
	}
	fmt.Println(xmasCount)
}
