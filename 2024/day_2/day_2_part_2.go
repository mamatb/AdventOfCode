package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncreasing(report []int, dampenerAvailable bool) bool {
	var levelPrev int
	for levelIdx, level := range report {
		if levelIdx > 0 && (level <= levelPrev || level-levelPrev > 3) {
			if dampenerAvailable {
				dampenedL := append([]int{}, report[:levelIdx-1]...)
				dampenedL = append(dampenedL, report[levelIdx:]...)
				dampenedR := append([]int{}, report[:levelIdx]...)
				dampenedR = append(dampenedR, report[levelIdx+1:]...)
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
	for levelIdx, level := range report {
		if levelIdx > 0 && (level >= levelPrev || levelPrev-level > 3) {
			if dampenerAvailable {
				dampenedL := append([]int{}, report[:levelIdx-1]...)
				dampenedL = append(dampenedL, report[levelIdx:]...)
				dampenedR := append([]int{}, report[:levelIdx]...)
				dampenedR = append(dampenedR, report[levelIdx+1:]...)
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
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_2.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	safeReports := 0
	for inputScanner.Scan() {
		report := []int{}
		for _, levelString := range strings.Split(inputScanner.Text(), " ") {
			if level, err := strconv.Atoi(levelString); err == nil {
				report = append(report, level)
			} else {
				panic(err)
			}
		}
		if isIncreasing(report, true) || isDecreasing(report, true) {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}
