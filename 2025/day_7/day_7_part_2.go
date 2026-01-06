package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row int
	col int
}

func countTimelines(start position, rows int, splits map[position]bool) int {
	timelines, pending, pos := map[position]int{}, []position{start}, position{}
	for len(pending) > 0 {
		pos, pending = pending[len(pending)-1], pending[:len(pending)-1]
		if _, ok := timelines[pos]; ok {
			continue
		}
		if pos.row >= rows {
			timelines[pos] = 1
			continue
		}
		if splits[pos] {
			left := position{row: pos.row, col: pos.col - 1}
			right := position{row: pos.row, col: pos.col + 1}
			if tlLeft, ok := timelines[left]; !ok {
				pending = append(pending, pos)
				pending = append(pending, left)
			} else if tlRight, ok := timelines[right]; !ok {
				pending = append(pending, pos)
				pending = append(pending, right)
			} else {
				timelines[pos] = tlLeft + tlRight
			}
		} else {
			down := position{row: pos.row + 1, col: pos.col}
			if tlDown, ok := timelines[down]; !ok {
				pending = append(pending, pos)
				pending = append(pending, down)
			} else {
				timelines[pos] = tlDown
			}
		}
	}
	return timelines[start]
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_7.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	splits, rows, start := map[position]bool{}, 0, position{}
	for inputScanner.Scan() {
		for col, symbol := range inputScanner.Text() {
			if symbol == 'S' {
				start = position{row: rows, col: col}
			} else if symbol == '^' {
				splits[position{row: rows, col: col}] = true
			}
		}
		rows++
	}
	fmt.Println(countTimelines(start, rows, splits))
}
