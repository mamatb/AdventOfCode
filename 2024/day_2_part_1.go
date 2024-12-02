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

func isIncreasing(report []string) bool {
	var levelPrev int
	for index, levelString := range report {
		level, err := strconv.Atoi(levelString)
		errCheck(err)
		if index > 0 && (level <= levelPrev || level-levelPrev > 3) {
			return false
		}
		levelPrev = level
	}
	return true
}

func isDecreasing(report []string) bool {
	var levelPrev int
	for index, levelString := range report {
		level, err := strconv.Atoi(levelString)
		errCheck(err)
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
	inputScan := inputScanner.Scan()
	inputLine := inputScanner.Text()
	for inputScan {
		inputLineSplit := strings.Split(inputLine, " ")
		if isIncreasing(inputLineSplit) || isDecreasing(inputLineSplit) {
			safeReports += 1
		}
		inputScan = inputScanner.Scan()
		inputLine = inputScanner.Text()
	}
	fmt.Println(safeReports)
}
