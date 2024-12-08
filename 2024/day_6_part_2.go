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

func posInMap(mapRows int, mapCols int, pos position) bool {
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

func loopInMap(mapRows int, mapCols int, guardPos position, guardDir int,
	obstacles map[position]bool) bool {
	visitedByDir := []map[position]bool{}
	for i := 0; i < 4; i++ {
		visitedByDir = append(visitedByDir, map[position]bool{})
	}
	for posInMap(mapRows, mapCols, guardPos) {
		if visitedByDir[guardDir][guardPos] {
			return true
		}
		visitedByDir[guardDir][guardPos] = true
		guardPos, guardDir = nextPosDir(guardPos, guardDir, obstacles)
	}
	return false
}

func main() {
	input, err := os.Open("day_6.txt")
	errCheck(err)
	defer input.Close()
	mapRows, mapCols, loops, guardPos, guardDir := 0, 0, 0, position{}, north
	obstacles := map[position]bool{}
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
	visited, guardPosPrev := map[position]bool{guardPos: true}, guardPos
	guardPos, guardDir = nextPosDir(guardPos, guardDir, obstacles)
	for posInMap(mapRows, mapCols, guardPos) {
		if !visited[guardPos] {
			obstacles[guardPos] = true
			if loopInMap(mapRows, mapCols, guardPosPrev, guardDir, obstacles) {
				loops += 1
			}
			obstacles[guardPos], visited[guardPos] = false, true
		}
		guardPosPrev = guardPos
		guardPos, guardDir = nextPosDir(guardPos, guardDir, obstacles)
	}
	fmt.Println(loops)
}
