package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncreasing(report []int) bool {
	var levelPrev int
	for levelIdx, level := range report {
		if levelIdx > 0 && (level <= levelPrev || level-levelPrev > 3) {
			return false
		}
		levelPrev = level
	}
	return true
}

func isDecreasing(report []int) bool {
	var levelPrev int
	for levelIdx, level := range report {
		if levelIdx > 0 && (level >= levelPrev || levelPrev-level > 3) {
			return false
		}
		levelPrev = level
	}
	return true
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_2.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	safeReports := 0
	for inputScanner.Scan() {
		report := []int{}
		for _, levelString := range strings.Split(inputScanner.Text(), " ") {
			if level, err := strconv.Atoi(levelString); err != nil {
				panic(err)
			} else {
				report = append(report, level)
			}
		}
		if isIncreasing(report) || isDecreasing(report) {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}
