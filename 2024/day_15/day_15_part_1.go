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

func movedPos(pos position, rDelta int, cDelta int) position {
	return position{row: pos.row + rDelta, col: pos.col + cDelta}
}

func moveRobot(rDelta int, cDelta int, robot position, boxes map[position]bool,
	walls map[position]bool) (position, map[position]bool) {
	robotMoved := movedPos(robot, rDelta, cDelta)
	if walls[robotMoved] {
		return robot, boxes
	}
	if boxes[robotMoved] {
		boxMoved := movedPos(robotMoved, rDelta, cDelta)
		for boxes[boxMoved] {
			boxMoved = movedPos(boxMoved, rDelta, cDelta)
		}
		if walls[boxMoved] {
			return robot, boxes
		}
		delete(boxes, robotMoved)
		boxes[boxMoved] = true
	}
	return robotMoved, boxes
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_15.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	row, coordinates := 0, 0
	robot, boxes, walls := position{}, map[position]bool{}, map[position]bool{}
	for inputScanner.Scan() {
		inputLine := strings.Split(inputScanner.Text(), "")
		if len(inputLine) == 0 {
			break
		}
		for col, inputSymbol := range inputLine {
			switch inputSymbol {
			case "#":
				walls[position{row: row, col: col}] = true
			case "O":
				boxes[position{row: row, col: col}] = true
			case "@":
				robot.row, robot.col = row, col
			}
		}
		row++
	}
	for inputScanner.Scan() {
		for _, movement := range strings.Split(inputScanner.Text(), "") {
			rDelta, cDelta := 0, 0
			switch movement {
			case "^":
				rDelta--
			case ">":
				cDelta++
			case "v":
				rDelta++
			case "<":
				cDelta--
			}
			robot, boxes = moveRobot(rDelta, cDelta, robot, boxes, walls)
		}
	}
	for box := range boxes {
		coordinates += box.row*100 + box.col
	}
	fmt.Println(coordinates)
}
