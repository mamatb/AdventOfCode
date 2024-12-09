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

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func posInMap(mapRows int, mapCols int, pos position) bool {
	return pos.row >= 0 && pos.row < mapRows && pos.col >= 0 && pos.col < mapCols
}

func addAntinodes(antinodes map[position]bool, antennas []position, mapRows int,
	mapCols int) map[position]bool {
	antenna, antennasPrev := position{}, []position{}
	for len(antennas) > 0 {
		antenna, antennas = antennas[0], antennas[1:]
		for _, antennaPrev := range antennasPrev {
			antinodeOne := position{row: 2*antenna.row - antennaPrev.row,
				col: 2*antenna.col - antennaPrev.col}
			if posInMap(mapRows, mapCols, antinodeOne) {
				antinodes[antinodeOne] = true
			}
			antinodeTwo := position{row: 2*antennaPrev.row - antenna.row,
				col: 2*antennaPrev.col - antenna.col}
			if posInMap(mapRows, mapCols, antinodeTwo) {
				antinodes[antinodeTwo] = true
			}
		}
		antennasPrev = append(antennasPrev, antenna)
	}
	return antinodes
}

func main() {
	input, err := os.Open("day_8.txt")
	errCheck(err)
	defer input.Close()
	antennasByFreq, antinodes := map[string][]position{}, map[position]bool{}
	mapRows, mapCols, inputScanner := 0, 0, bufio.NewScanner(input)
	for inputScanner.Scan() {
		for col, freq := range strings.Split(inputScanner.Text(), "") {
			if freq != "." {
				freqPos := position{row: mapRows, col: col}
				antennasByFreq[freq] = append(antennasByFreq[freq], freqPos)
			}
			if mapRows == 0 {
				mapCols += 1
			}
		}
		mapRows += 1
	}
	for _, antennas := range antennasByFreq {
		antinodes = addAntinodes(antinodes, antennas, mapRows, mapCols)
	}
	fmt.Println(len(antinodes))
}
