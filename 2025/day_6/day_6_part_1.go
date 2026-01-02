package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_6.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	inputScanner.Scan()
	total, problems := 0, len(strings.Fields(inputScanner.Text()))
	added := slices.Repeat([]int{0}, problems)
	multiplied := slices.Repeat([]int{1}, problems)
	for inputScanner.Text()[0] != '+' && inputScanner.Text()[0] != '*' {
		for numIdx, numString := range strings.Fields(inputScanner.Text()) {
			if num, err := strconv.Atoi(numString); err != nil {
				panic(err)
			} else {
				added[numIdx] += num
				multiplied[numIdx] *= num
			}
		}
		inputScanner.Scan()
	}
	for symbolIdx, symbol := range strings.Fields(inputScanner.Text()) {
		if symbol == "+" {
			total += added[symbolIdx]
		} else {
			total += multiplied[symbolIdx]
		}
	}
	fmt.Println(total)
}
