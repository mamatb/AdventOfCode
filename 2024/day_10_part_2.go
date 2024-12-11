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

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func posInMap(pos position, mapRows int, mapCols int) bool {
	return pos.row >= 0 && pos.row < mapRows && pos.col >= 0 && pos.col < mapCols
}

func getRating(trailhead position, topomap [][]int) int {
	pos, positions, positionsLen := position{}, []position{trailhead}, 1
	for height := 0; height < 9; height++ {
		for i := 0; i < positionsLen; i++ {
			pos, positions = positions[0], positions[1:]
			posReachables := []position{pos, pos, pos, pos}
			posReachables[0].row -= 1 // north
			posReachables[1].col += 1 // east
			posReachables[2].row += 1 // south
			posReachables[3].col -= 1 // west
			for _, reachable := range posReachables {
				if posInMap(reachable, len(topomap), len(topomap[0])) &&
					topomap[reachable.row][reachable.col] == height+1 {
					positions = append(positions, reachable)
				}
			}
		}
		positionsLen = len(positions)
	}
	return positionsLen
}

func main() {
	input, err := os.Open("day_10.txt")
	errCheck(err)
	defer input.Close()
	rating, row, topomap, trailheads := 0, 0, [][]int{}, []position{}
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		topomap = append(topomap, []int{})
		for col, heightString := range strings.Split(inputScanner.Text(), "") {
			height, err := strconv.Atoi(heightString)
			errCheck(err)
			topomap[row] = append(topomap[row], height)
			if height == 0 {
				trailheads = append(trailheads, position{row: row, col: col})
			}
		}
		row += 1
	}
	for _, trailhead := range trailheads {
		rating += getRating(trailhead, topomap)
	}
	fmt.Println(rating)
}
