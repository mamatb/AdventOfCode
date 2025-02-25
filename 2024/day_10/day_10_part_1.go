package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row int
	col int
}

func posInMap(pos position, mapRows int, mapCols int) bool {
	return pos.row >= 0 && pos.row < mapRows && pos.col >= 0 && pos.col < mapCols
}

func getScore(trailhead position, topomap [][]int) int {
	pos, positions, positionsLen := position{}, []position{trailhead}, 1
	positionsReached := map[position]bool{trailhead: true}
	for height := range 9 {
		for range positionsLen {
			pos, positions = positions[0], positions[1:]
			posReachables := []position{pos, pos, pos, pos}
			posReachables[0].row-- // north
			posReachables[1].col++ // east
			posReachables[2].row++ // south
			posReachables[3].col-- // west
			for _, reachable := range posReachables {
				if posInMap(reachable, len(topomap), len(topomap[0])) &&
					topomap[reachable.row][reachable.col] == height+1 &&
					!positionsReached[reachable] {
					positions = append(positions, reachable)
					positionsReached[reachable] = true
				}
			}
		}
		positionsLen = len(positions)
	}
	return positionsLen
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_10.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	score, row, topomap, trailheads := 0, 0, [][]int{}, []position{}
	for inputScanner.Scan() {
		topomap = append(topomap, []int{})
		for col, heightString := range strings.Split(inputScanner.Text(), "") {
			if height, err := strconv.Atoi(heightString); err == nil {
				topomap[row] = append(topomap[row], height)
				if height == 0 {
					trailheads = append(trailheads, position{row: row, col: col})
				}
			} else {
				panic(err)
			}
		}
		row++
	}
	for _, trailhead := range trailheads {
		score += getScore(trailhead, topomap)
	}
	fmt.Println(score)
}
