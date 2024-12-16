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

func posInMap(pos position, mapRows int, mapCols int) bool {
	return pos.row >= 0 && pos.row < mapRows && pos.col >= 0 && pos.col < mapCols
}

func getNeighbours(pos position) []position {
	return []position{
		{row: pos.row - 1, col: pos.col}, // north
		{row: pos.row, col: pos.col + 1}, // east
		{row: pos.row + 1, col: pos.col}, // south
		{row: pos.row, col: pos.col - 1}} // west
}

func getRegion(garden [][]string, inRegion map[position]bool,
	pos position) map[position]bool {
	label, rows, cols := garden[pos.row][pos.col], len(garden), len(garden[0])
	region, posPending := map[position]bool{}, []position{pos}
	for len(posPending) > 0 {
		pos, posPending = posPending[0], posPending[1:]
		if !inRegion[pos] && garden[pos.row][pos.col] == label {
			region[pos], inRegion[pos] = true, true
			for _, posN := range getNeighbours(pos) {
				if posInMap(posN, rows, cols) {
					posPending = append(posPending, posN)
				}
			}
		}
	}
	return region
}

func getPrice(region map[position]bool) int {
	area, perimeter := len(region), 0
	for pos := range region {
		perimeter += 4
		for _, posN := range getNeighbours(pos) {
			if region[posN] {
				perimeter -= 1
			}
		}
	}
	return area * perimeter
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_12.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	garden, price := [][]string{}, 0
	regions, inRegion := []map[position]bool{}, map[position]bool{}
	for inputScanner.Scan() {
		line := []string{}
		for _, label := range strings.Split(inputScanner.Text(), "") {
			line = append(line, label)
		}
		garden = append(garden, line)
	}
	for row, line := range garden {
		for col := range line {
			pos := position{row: row, col: col}
			if !inRegion[pos] {
				regions = append(regions, getRegion(garden, inRegion, pos))
			}
		}
	}
	for _, region := range regions {
		price += getPrice(region)
	}
	fmt.Println(price)
}
