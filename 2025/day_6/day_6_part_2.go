package main

import (
	"bufio"
	"fmt"
	"os"
)

func verticalNumber(lines []string, idx int) int {
	num := 0
	for _, line := range lines {
		if line[idx] != ' ' {
			num = num*10 + int(line[idx]-'0')
		}
	}
	return num
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_6.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	var lines []string
	for inputScanner.Scan() {
		if inputScanner.Text()[0] == '+' || inputScanner.Text()[0] == '*' {
			break
		}
		lines = append(lines, inputScanner.Text())
	}
	var total, problem, symbol int
	for charIdx, char := range inputScanner.Text() {
		switch char {
		case '+':
			total += problem
			problem, symbol = verticalNumber(lines, charIdx), int(char)
		case '*':
			total += problem
			problem, symbol = verticalNumber(lines, charIdx), int(char)
		case ' ':
			if symbol == '+' {
				problem += verticalNumber(lines, charIdx)
			} else {
				problem *= max(1, verticalNumber(lines, charIdx))
			}
		}
	}
	total += problem
	fmt.Println(total)
}
