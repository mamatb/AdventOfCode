package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func isIncreasing(report []int) bool {
	var levelPrev int
	for index, level := range report {
		if index > 0 && (level <= levelPrev || level-levelPrev > 3) {
			return false
		}
		levelPrev = level
	}
	return true
}

func isDecreasing(report []int) bool {
	var levelPrev int
	for index, level := range report {
		if index > 0 && (level >= levelPrev || levelPrev-level > 3) {
			return false
		}
		levelPrev = level
	}
	return true
}

func main() {
	input, err := os.Open("day_2.txt")
	errCheck(err)
	defer input.Close()
	safeReports := 0
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		report := []int{}
		for _, levelString := range strings.Split(inputScanner.Text(), " ") {
			level, err := strconv.Atoi(levelString)
			errCheck(err)
			report = append(report, level)
		}
		if isIncreasing(report) || isDecreasing(report) {
			safeReports += 1
		}
	}
	fmt.Println(safeReports)
}
