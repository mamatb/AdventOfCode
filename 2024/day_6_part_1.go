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

func posInMap(pos position, mapRows int, mapCols int) bool {
	return pos.row >= 0 && pos.row < mapRows && pos.col >= 0 && pos.col < mapCols
}

func nextPosDir(guardPos position, guardDir int, obstacles map[position]bool) (
	position, int) {
	guardPosBackup := guardPos
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
	if obstacles[guardPos] {
		guardPos, guardDir = guardPosBackup, (guardDir+1)%4
	}
	return guardPos, guardDir
}

func main() {
	input, err := os.Open("day_6.txt")
	errCheck(err)
	defer input.Close()
	mapRows, mapCols, guardPos, guardDir := 0, 0, position{}, north
	obstacles, visited := map[position]bool{}, map[position]bool{}
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
	for posInMap(guardPos, mapRows, mapCols) {
		visited[guardPos] = true
		guardPos, guardDir = nextPosDir(guardPos, guardDir, obstacles)
	}
	fmt.Println(len(visited))
}
