package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type position struct {
	row int
	col int
}

type bathRobot struct {
	rowVel int
	colVel int
}

func posInRobotRow(pos position, size int, robotsByPos map[position][]bathRobot,
	cols int) bool {
	if cols-pos.col < size {
		return false
	}
	for col := pos.col; col < pos.col+size; col++ {
		if len(robotsByPos[position{row: pos.row, col: col}]) == 0 {
			return false
		}
	}
	return true
}

func existsRobotRow(size int, robotsByPos map[position][]bathRobot, cols int) bool {
	for pos := range robotsByPos {
		if posInRobotRow(pos, size, robotsByPos, cols) {
			return true
		}
	}
	return false
}

func moveRobots(robotsByPos map[position][]bathRobot, rows int,
	cols int) map[position][]bathRobot {
	robotsByPosMov := map[position][]bathRobot{}
	for pos, robots := range robotsByPos {
		for _, robot := range robots {
			posMov := position{
				row: (pos.row + robot.rowVel) % rows,
				col: (pos.col + robot.colVel) % cols}
			if posMov.row < 0 {
				posMov.row += rows
			}
			if posMov.col < 0 {
				posMov.col += cols
			}
			robotsByPosMov[posMov] = append(robotsByPosMov[posMov], robot)
		}
	}
	return robotsByPosMov
}

func plotRobots(robotsByPos map[position][]bathRobot, rows int, cols int) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if len(robotsByPos[position{row: row, col: col}]) > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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
	robotsByPos, rows, cols, seconds := map[position][]bathRobot{}, 103, 101, 0
	for inputScanner.Scan() {
		pos, robot := position{}, bathRobot{}
		for idx, numStr := range robotExp.FindStringSubmatch(inputScanner.Text())[1:] {
			if num, err := strconv.Atoi(numStr); err == nil {
				switch idx {
				case 0:
					pos.col = num
				case 1:
					pos.row = num
				case 2:
					robot.colVel = num
				case 3:
					robot.rowVel = num
				}
			} else {
				panic(err)
			}
		}
		robotsByPos[pos] = append(robotsByPos[pos], robot)
	}
	for !existsRobotRow(16, robotsByPos, cols) {
		robotsByPos = moveRobots(robotsByPos, rows, cols)
		seconds += 1
	}
	plotRobots(robotsByPos, rows, cols)
	fmt.Println(seconds)
}
