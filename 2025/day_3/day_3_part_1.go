package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxJoltage(batteries string) int {
	idxFirst, idxSecond := -1, -1
	for idx := 0; idx < len(batteries)-1; idx++ {
		if idxFirst == -1 || batteries[idx] > batteries[idxFirst] {
			idxFirst = idx
		}
	}
	for idx := idxFirst + 1; idx < len(batteries); idx++ {
		if idxSecond == -1 || batteries[idx] > batteries[idxSecond] {
			idxSecond = idx
		}
	}
	return int(batteries[idxFirst]-'0')*10 + int(batteries[idxSecond]-'0')
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
		joltage += maxJoltage(inputScanner.Text())
	}
	fmt.Println(joltage)
}
