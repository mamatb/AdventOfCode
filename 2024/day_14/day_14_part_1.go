package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type position struct {
	row int
	col int
}

func addRobot(robotsByQuad []int, pos position, rows int, cols int) []int {
	midRow, midCol := rows/2, cols/2
	switch cmp.Compare(pos.row, midRow) {
	case -1:
		switch cmp.Compare(pos.col, midCol) {
		case -1:
			robotsByQuad[0] += 1
		case 1:
			robotsByQuad[1] += 1
		}
	case 1:
		switch cmp.Compare(pos.col, midCol) {
		case -1:
			robotsByQuad[2] += 1
		case 1:
			robotsByQuad[3] += 1
		}
	}
	return robotsByQuad
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_14.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	robotExp, err := regexp.Compile(`p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)
	if err != nil {
		panic(err)
	}
	safetyFactor, seconds, rows, cols := 1, 100, 103, 101
	robotsByQuad := []int{0, 0, 0, 0} // topleft, topright, downleft, downright
	for inputScanner.Scan() {
		pos := position{}
		for idx, numStr := range robotExp.FindStringSubmatch(inputScanner.Text())[1:] {
			if num, err := strconv.Atoi(numStr); err == nil {
				switch idx {
				case 0:
					pos.col = num
				case 1:
					pos.row = num
				case 2:
					pos.col = (pos.col + num*seconds) % cols
					if pos.col < 0 {
						pos.col += cols
					}
				case 3:
					pos.row = (pos.row + num*seconds) % rows
					if pos.row < 0 {
						pos.row += rows
					}
				}
			} else {
				panic(err)
			}
		}
		robotsByQuad = addRobot(robotsByQuad, pos, rows, cols)
	}
	for _, robots := range robotsByQuad {
		safetyFactor *= robots
	}
	fmt.Println(safetyFactor)
}
