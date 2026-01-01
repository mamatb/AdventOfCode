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

func posInMap(pos position, mapRows int, mapCols int) bool {
	return pos.row >= 0 && pos.row < mapRows && pos.col >= 0 && pos.col < mapCols
}

func nextPosDir(guardPos position, guardDir int, obstacles map[position]bool) (
	position, int) {
	guardPosBackup := guardPos
	switch guardDir {
	case north:
		guardPos.row--
	case east:
		guardPos.col++
	case south:
		guardPos.row++
	default: // west
		guardPos.col--
	}
	if obstacles[guardPos] {
		guardPos, guardDir = guardPosBackup, (guardDir+1)%4
	}
	return guardPos, guardDir
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_6.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	mapRows, mapCols, guardPos, guardDir := 0, 0, position{}, north
	obstacles, visited := map[position]bool{}, map[position]bool{}
	for inputScanner.Scan() {
		for col, symbol := range strings.Split(inputScanner.Text(), "") {
			switch symbol {
			case "^":
				guardPos.row, guardPos.col = mapRows, col
			case "#":
				obstacles[position{row: mapRows, col: col}] = true
			}
			if mapRows == 0 {
				mapCols++
			}
		}
		mapRows++
	}
	for posInMap(guardPos, mapRows, mapCols) {
		visited[guardPos] = true
		guardPos, guardDir = nextPosDir(guardPos, guardDir, obstacles)
	}
	fmt.Println(len(visited))
}
