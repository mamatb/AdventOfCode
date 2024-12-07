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

type direction struct {
	rowDelta int
	colDelta int
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func guardInMap(mapRows int, mapCols int, guardPos position) bool {
	return (guardPos.row >= 0 &&
		guardPos.row < mapRows &&
		guardPos.col >= 0 &&
		guardPos.col < mapCols)
}

func nextDirection(guardDir direction) direction {
	if guardDir.rowDelta == -1 && guardDir.colDelta == 0 { // up to right
		guardDir.rowDelta, guardDir.colDelta = 0, 1
	} else if guardDir.rowDelta == 0 && guardDir.colDelta == 1 { // right to down
		guardDir.rowDelta, guardDir.colDelta = 1, 0
	} else if guardDir.rowDelta == 1 && guardDir.colDelta == 0 { // down to left
		guardDir.rowDelta, guardDir.colDelta = 0, -1
	} else { // left to up
		guardDir.rowDelta, guardDir.colDelta = -1, 0
	}
	return guardDir
}

func main() {
	input, err := os.Open("day_6.txt")
	errCheck(err)
	defer input.Close()
	mapRows, mapCols := 0, 0
	obstacles, visited := map[position]bool{}, map[position]bool{}
	guardPos, guardDir := position{}, direction{rowDelta: -1, colDelta: 0}
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		for col, symbol := range strings.Split(inputScanner.Text(), "") {
			switch symbol {
			case "#":
				obstacles[position{row: mapRows, col: col}] = true
			case "^":
				guardPos.row, guardPos.col = mapRows, col
			}
			if mapRows == 0 {
				mapCols += 1
			}
		}
		mapRows += 1
	}
	for guardInMap(mapRows, mapCols, guardPos) {
		visited[guardPos] = true
		guardPos.row += guardDir.rowDelta
		guardPos.col += guardDir.colDelta
		if obstacles[guardPos] {
			guardPos.row -= guardDir.rowDelta
			guardPos.col -= guardDir.colDelta
			guardDir = nextDirection(guardDir)
		}
	}
	fmt.Println(len(visited))
}
