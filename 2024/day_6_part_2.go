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

func loopInMap(mapRows int, mapCols int, guardPos position, guardDir int,
	obstacles map[position]bool) bool {
	visitedByDir := make([]map[position]bool, 4)
	for i := 0; i < 4; i++ {
		visitedByDir[i] = map[position]bool{}
	}
	for guardInMap(mapRows, mapCols, guardPos) {
		if visitedByDir[guardDir][guardPos] {
			return true
		}
		visitedByDir[guardDir][guardPos] = true
		guardPosBackup := guardPos
		guardPos = nextPosition(guardPos, guardDir)
		if obstacles[guardPos] {
			guardPos = guardPosBackup
			guardDir = nextDirection(guardDir)
		}
	}
	return false
}

func main() {
	input, err := os.Open("day_6.txt")
	errCheck(err)
	defer input.Close()
	mapRows, mapCols, loops, guardPos, guardDir := 0, 0, 0, position{}, north
	gaps, obstacles := []position{}, map[position]bool{}
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		for col, symbol := range strings.Split(inputScanner.Text(), "") {
			switch symbol {
			case "^":
				guardPos.row, guardPos.col = mapRows, col
			case "#":
				obstacles[position{row: mapRows, col: col}] = true
			default: // "."
				gaps = append(gaps, position{row: mapRows, col: col})
			}
			if mapRows == 0 {
				mapCols += 1
			}
		}
		mapRows += 1
	}
	for gapIndex, gap := range gaps {
		if gapIndex > 0 {
			delete(obstacles, gaps[gapIndex-1])
		}
		obstacles[gap] = true
		if loopInMap(mapRows, mapCols, guardPos, guardDir, obstacles) {
			loops += 1
		}
	}
	fmt.Println(loops)
}
