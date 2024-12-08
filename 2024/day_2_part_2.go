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

func isIncreasing(report []int, dampenerAvailable bool) bool {
	var levelPrev int
	for levelIndex, level := range report {
		if levelIndex > 0 && (level <= levelPrev || level-levelPrev > 3) {
			if dampenerAvailable {
				dampenedL := append([]int{}, report[:levelIndex-1]...)
				dampenedL = append(dampenedL, report[levelIndex:]...)
				dampenedR := append([]int{}, report[:levelIndex]...)
				dampenedR = append(dampenedR, report[levelIndex+1:]...)
				return isIncreasing(dampenedL, false) || isIncreasing(dampenedR, false)
			} else {
				return false
			}
		}
		levelPrev = level
	}
	return true
}

func isDecreasing(report []int, dampenerAvailable bool) bool {
	var levelPrev int
	for levelIndex, level := range report {
		if levelIndex > 0 && (level >= levelPrev || levelPrev-level > 3) {
			if dampenerAvailable {
				dampenedL := append([]int{}, report[:levelIndex-1]...)
				dampenedL = append(dampenedL, report[levelIndex:]...)
				dampenedR := append([]int{}, report[:levelIndex]...)
				dampenedR = append(dampenedR, report[levelIndex+1:]...)
				return isDecreasing(dampenedL, false) || isDecreasing(dampenedR, false)
			} else {
				return false
			}
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
		if isIncreasing(report, true) || isDecreasing(report, true) {
			safeReports += 1
		}
	}
	fmt.Println(safeReports)
}
