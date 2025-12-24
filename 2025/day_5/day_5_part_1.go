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

func isContained(ingredient int, freshRs []freshRange) bool {
	left, right := 0, len(freshRs)-1
	for left <= right {
		mid := (left + right) / 2
		if freshRs[mid].start <= ingredient && ingredient <= freshRs[mid].end {
			return true
		}
		if freshRs[mid].start > ingredient {
			right = mid - 1
		}
		if freshRs[mid].end < ingredient {
			left = mid + 1
		}
	}
	return false
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
	freshRs = deduplicateRanges(freshRs)
	freshIngredients := 0
	for inputScanner.Scan() {
		if ingredient, err := strconv.Atoi(inputScanner.Text()); err != nil {
			panic(err)
		} else if isContained(ingredient, freshRs) {
			freshIngredients++
		}
	}
	fmt.Println(freshIngredients)
}
