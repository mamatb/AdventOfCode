package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

type position struct {
	row int
	col int
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

func nextPosition(guardPos position, guardDir int) position {
	switch guardDir {
	case north:
		guardPos.row -= 1
	case east:
		guardPos.col += 1
	case south:
		guardPos.row += 1
	default: // west
		guardPos.col -= 1
	}
	return guardPos
}

func nextDirection(guardDir int) int {
	return (guardDir + 1) % 4
}

func main() {
	input, err := os.Open("day_6.txt")
	errCheck(err)
	defer input.Close()
	mapRows, mapCols := 0, 0
	obstacles, visited := map[position]bool{}, map[position]bool{}
	guardPos, guardDir := position{}, north
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		for col, symbol := range strings.Split(inputScanner.Text(), "") {
			switch symbol {
			case "^":
				guardPos.row, guardPos.col = mapRows, col
			case "#":
				obstacles[position{row: mapRows, col: col}] = true
			}
			if mapRows == 0 {
				mapCols += 1
			}
		}
		mapRows += 1
	}
	for guardInMap(mapRows, mapCols, guardPos) {
		visited[guardPos] = true
		guardPosBackup := guardPos
		guardPos = nextPosition(guardPos, guardDir)
		if obstacles[guardPos] {
			guardPos = guardPosBackup
			guardDir = nextDirection(guardDir)
		}
	}
	fmt.Println(len(visited))
}
