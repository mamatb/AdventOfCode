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
	fmt.Println(scores[end])
}
