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

func countSplits(start position, rows int, splits map[position]bool) int {
	count, pending, pos := 0, []position{start}, position{}
	posVisited := map[position]bool{}
	for len(pending) > 0 {
		pos, pending = pending[len(pending)-1], pending[:len(pending)-1]
		if pos.row >= rows || posVisited[pos] {
			continue
		}
		posVisited[pos] = true
		pos.row++
		if splits[pos] {
			count++
			pending = append(pending, position{row: pos.row, col: pos.col - 1})
			pending = append(pending, position{row: pos.row, col: pos.col + 1})
		} else {
			pending = append(pending, pos)
		}
	}
	return count
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
	fmt.Println(countSplits(start, rows, splits))
}
