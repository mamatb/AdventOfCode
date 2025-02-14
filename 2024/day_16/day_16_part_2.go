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

type direction struct {
	rDelta int
	cDelta int
}

type situation struct {
	pos   position
	dir   direction
	score int
}

func getNextPos(pos position, dir direction) position {
	return position{row: pos.row + dir.rDelta, col: pos.col + dir.cDelta}
}

func getNextDirs(dir direction) []direction {
	if dir.rDelta == 0 {
		return []direction{dir, {rDelta: 1, cDelta: 0}, {rDelta: -1, cDelta: 0}}
	} else {
		return []direction{dir, {rDelta: 0, cDelta: 1}, {rDelta: 0, cDelta: -1}}
	}
}

func getNeighbours(pos position, walls map[position]bool) []position {
	neighbours, neighboursPot := []position{}, []position{
		{row: pos.row - 1, col: pos.col}, // north
		{row: pos.row, col: pos.col + 1}, // east
		{row: pos.row + 1, col: pos.col}, // south
		{row: pos.row, col: pos.col - 1}} // west
	for _, neighbour := range neighboursPot {
		if !walls[neighbour] {
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func getScores(start position, walls map[position]bool) map[position]int {
	scores := map[position]int{start: 0}
	sit, sitPending := situation{}, []situation{{
		pos:   start,
		dir:   direction{rDelta: 0, cDelta: 1},
		score: 0}}
	for len(sitPending) > 0 {
		sit, sitPending = sitPending[0], sitPending[1:]
		for _, dirNext := range getNextDirs(sit.dir) {
			posNext := getNextPos(sit.pos, dirNext)
			if !walls[posNext] {
				scoreNext := sit.score + 1
				if dirNext != sit.dir {
					scoreNext += 1000
				}
				if score, ok := scores[posNext]; !ok || score > scoreNext {
					scores[posNext] = scoreNext
					sitPending = append(sitPending, situation{
						pos:   posNext,
						dir:   dirNext,
						score: scoreNext})
				}
			}
		}
	}
	return scores
}

func getBestPaths(start position, end position, walls map[position]bool,
	scores map[position]int) map[position]bool {
	bestPaths := map[position]bool{start: true, end: true}
	pos, posPending := position{}, []position{end}
	for len(posPending) > 0 {
		pos, posPending = posPending[0], posPending[1:]
		for _, posNext := range getNeighbours(pos, walls) {
			if bestPaths[posNext] {
				continue
			}
			scoreDelta := scores[pos] - scores[posNext]
			if scoreDelta == 1 || scoreDelta == 1001 {
				bestPaths[posNext] = true
				posPending = append(posPending, posNext)
			} else if scoreDelta > 0 {
				posNextNext := getNextPos(posNext, direction{
					rDelta: posNext.row - pos.row,
					cDelta: posNext.col - pos.col})
				scoreDelta = scores[pos] - scores[posNextNext]
				if scoreDelta == 2 || scoreDelta == 1002 {
					bestPaths[posNext], bestPaths[posNextNext] = true, true
					posPending = append(posPending, posNextNext)
				}
			} else {
				posPrev := getNextPos(pos, direction{
					rDelta: pos.row - posNext.row,
					cDelta: pos.col - posNext.col})
				scoreDelta = scores[posPrev] - scores[posNext]
				if (scoreDelta == 2 || scoreDelta == 1002) && bestPaths[posPrev] {
					bestPaths[posNext] = true
					posPending = append(posPending, posNext)
				}
			}
		}
	}
	return bestPaths
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_16.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	row, start, end, walls := 0, position{}, position{}, map[position]bool{}
	for inputScanner.Scan() {
		for col, tile := range strings.Split(inputScanner.Text(), "") {
			switch tile {
			case "#":
				walls[position{row: row, col: col}] = true
			case "S":
				start.row, start.col = row, col
			case "E":
				end.row, end.col = row, col
			}
		}
		row++
	}
	scores := getScores(start, walls)
	bestPaths := getBestPaths(start, end, walls, scores)
	fmt.Println(len(bestPaths))
}
