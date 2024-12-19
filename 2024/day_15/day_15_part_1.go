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

func moveRobot(rowDelta int, colDelta int, robot position,
	boxes map[position]bool) (position, map[position]bool) {
	robot.row, robot.col = robot.row+rowDelta, robot.col+colDelta
	if boxes[robot] {
		delete(boxes, robot)
		boxNew := position{row: robot.row + rowDelta, col: robot.col + colDelta}
		for boxes[boxNew] {
			boxNew.row, boxNew.col = boxNew.row+rowDelta, boxNew.col+colDelta
		}
		boxes[boxNew] = true
	}
	return robot, boxes
}

func canMoveRobot(rowDelta int, colDelta int, robot position,
	boxes map[position]bool, walls map[position]bool) bool {
	robot.row, robot.col = robot.row+rowDelta, robot.col+colDelta
	for !walls[robot] {
		if !boxes[robot] {
			return true
		}
		robot.row, robot.col = robot.row+rowDelta, robot.col+colDelta
	}
	return false
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_15.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
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
		row += 1
	}
	for inputScanner.Scan() {
		for _, movement := range strings.Split(inputScanner.Text(), "") {
			rowDelta, colDelta := 0, 0
			switch movement {
			case "^":
				rowDelta -= 1
			case ">":
				colDelta += 1
			case "v":
				rowDelta += 1
			case "<":
				colDelta -= 1
			}
			if canMoveRobot(rowDelta, colDelta, robot, boxes, walls) {
				robot, boxes = moveRobot(rowDelta, colDelta, robot, boxes)
			}
		}
	}
	for box := range boxes {
		coordinates += box.row*100 + box.col
	}
	fmt.Println(coordinates)
}
