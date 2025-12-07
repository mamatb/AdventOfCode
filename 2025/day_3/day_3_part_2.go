package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func maxJoltage(batteries string, battNum int) int {
	joltage, battIdxPrev := 0, -1
	for battNum > 0 {
		battIdx := -1
		for idx := battIdxPrev + 1; idx <= len(batteries)-battNum; idx++ {
			if battIdx == -1 || batteries[idx] > batteries[battIdx] {
				battIdx = idx
			}
		}
		joltage += int(batteries[battIdx]-'0') * int(math.Pow10(battNum-1))
		battIdxPrev = battIdx
		battNum--
	}
	return joltage
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_3.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	joltage := 0
	for inputScanner.Scan() {
		joltage += maxJoltage(inputScanner.Text(), 12)
	}
	fmt.Println(joltage)
}
