package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type position struct {
	row int
	col int
}

type wideBox struct {
	left  position
	right position
}

func movedPos(pos position, rDelta int, cDelta int) position {
	return position{row: pos.row + rDelta, col: pos.col + cDelta}
}

func pushPositions(rDelta int, cDelta int, boxToMove wideBox) []position {
	positions := []position{}
	if rDelta == 0 { // horizontal
		if cDelta > 0 {
			positions = append(positions, boxToMove.right)
		} else {
			positions = append(positions, boxToMove.left)
		}
	} else { // vertical
		positions = append(positions, boxToMove.left)
		positions = append(positions, boxToMove.right)
	}
	return positions
}

func potMovedBoxes(rDelta int, cDelta int, posToMove position) []wideBox {
	boxesMoved, posBoxesMoved := []wideBox{}, []position{}
	if rDelta == 0 { // horizontal
		posBoxesMoved = append(posBoxesMoved, movedPos(posToMove, 0, cDelta))
		posBoxesMoved = append(posBoxesMoved, movedPos(posToMove, 0, cDelta*2))
		if cDelta < 0 {
			slices.Reverse(posBoxesMoved)
		}
	} else { // vertical
		posBoxesMoved = append(posBoxesMoved, movedPos(posToMove, rDelta, -1))
		posBoxesMoved = append(posBoxesMoved, movedPos(posToMove, rDelta, 0))
		posBoxesMoved = append(posBoxesMoved, movedPos(posToMove, rDelta, 1))
	}
	for posIdx := range len(posBoxesMoved) - 1 {
		boxesMoved = append(boxesMoved, wideBox{
			left:  posBoxesMoved[posIdx],
			right: posBoxesMoved[posIdx+1]})
	}
	return boxesMoved
}

func moveRobot(rDelta int, cDelta int, robot position, boxes map[wideBox]bool,
	walls map[position]bool) (position, map[wideBox]bool) {
	robotMoved := movedPos(robot, rDelta, cDelta)
	if walls[robotMoved] {
		return robot, boxes
	}
	boxToMove, boxesToMove := wideBox{}, []wideBox{}
	boxesMoved, boxesToDel := []wideBox{}, map[wideBox]bool{}
	for _, potMovedBox := range potMovedBoxes(rDelta, cDelta, robot) {
		if boxes[potMovedBox] {
			boxesToDel[potMovedBox] = true
			boxesToMove = append(boxesToMove, potMovedBox)
		}
	}
	for len(boxesToMove) > 0 {
		boxToMove, boxesToMove = boxesToMove[0], boxesToMove[1:]
		boxMoved := wideBox{
			left:  movedPos(boxToMove.left, rDelta, cDelta),
			right: movedPos(boxToMove.right, rDelta, cDelta)}
		if walls[boxMoved.left] || walls[boxMoved.right] {
			return robot, boxes
		}
		boxesMoved = append(boxesMoved, boxMoved)
		for _, pushPos := range pushPositions(rDelta, cDelta, boxToMove) {
			for _, potMovedBox := range potMovedBoxes(rDelta, cDelta, pushPos) {
				if boxes[potMovedBox] && !boxesToDel[potMovedBox] {
					boxesToDel[potMovedBox] = true
					boxesToMove = append(boxesToMove, potMovedBox)
				}
			}
		}
	}
	for boxToDel := range boxesToDel {
		delete(boxes, boxToDel)
	}
	for _, boxMoved := range boxesMoved {
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
	robot, boxes, walls := position{}, map[wideBox]bool{}, map[position]bool{}
	for inputScanner.Scan() {
		inputLine := strings.Split(inputScanner.Text(), "")
		if len(inputLine) == 0 {
			break
		}
		for col, inputSymbol := range inputLine {
			switch inputSymbol {
			case "#":
				walls[position{row: row, col: col * 2}] = true
				walls[position{row: row, col: col*2 + 1}] = true
			case "O":
				boxes[wideBox{
					left:  position{row: row, col: col * 2},
					right: position{row: row, col: col*2 + 1}}] = true
			case "@":
				robot.row, robot.col = row, col*2
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
		coordinates += box.left.row*100 + box.left.col
	}
	fmt.Println(coordinates)
}
