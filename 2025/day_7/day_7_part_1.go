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

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_7.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	var pending []position
	splitMap, rows := map[position]bool{}, 0
	for inputScanner.Scan() {
		for col, symbol := range inputScanner.Text() {
			if symbol == 'S' {
				pending = append(pending, position{row: rows, col: col})
			} else if symbol == '^' {
				splitMap[position{row: rows, col: col}] = true
			}
		}
		rows++
	}
	splitCount, posVisited, pos := 0, map[position]bool{}, position{}
	for len(pending) > 0 {
		pos, pending = pending[len(pending)-1], pending[:len(pending)-1]
		if pos.row >= rows || posVisited[pos] {
			continue
		}
		posVisited[pos] = true
		pos.row++
		if splitMap[pos] {
			splitCount++
			pending = append(pending, position{row: pos.row, col: pos.col - 1})
			pending = append(pending, position{row: pos.row, col: pos.col + 1})
		} else {
			pending = append(pending, pos)
		}
	}
	fmt.Println(splitCount)
}
