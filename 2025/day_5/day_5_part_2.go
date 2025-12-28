package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type freshRange struct {
	start int
	end   int
}

func deduplicateRanges(freshRs []freshRange) []freshRange {
	freshRsDedup := []freshRange{freshRs[0]}
	for _, freshR := range freshRs[1:] {
		if freshR.start <= freshRsDedup[len(freshRsDedup)-1].end+1 {
			freshRsDedup[len(freshRsDedup)-1].end = max(freshR.end,
				freshRsDedup[len(freshRsDedup)-1].end)
		} else {
			freshRsDedup = append(freshRsDedup, freshR)
		}
	}
	return freshRsDedup
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_5.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	var freshRs []freshRange
	for inputScanner.Scan() {
		freshStrings := strings.Split(inputScanner.Text(), "-")
		if len(freshStrings) < 2 {
			break
		}
		if start, err := strconv.Atoi(freshStrings[0]); err != nil {
			panic(err)
		} else if end, err := strconv.Atoi(freshStrings[1]); err != nil {
			panic(err)
		} else {
			freshRs = append(freshRs, freshRange{start: start, end: end})
		}
	}
	slices.SortFunc(freshRs, func(a freshRange, b freshRange) int {
		return a.start - b.start
	})
	freshIngredients := 0
	for _, freshR := range deduplicateRanges(freshRs) {
		freshIngredients += freshR.end - freshR.start + 1
	}
	fmt.Println(freshIngredients)
}
