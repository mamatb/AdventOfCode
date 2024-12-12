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

func updateIsOrdered(update []string, rules map[string][]string) bool {
	pagesPrev := map[string]bool{}
	for _, page := range update {
		for _, pageRule := range rules[page] {
			if pagesPrev[pageRule] {
				return false
			}
		}
		pagesPrev[page] = true
	}
	return true
}

func main() {
	input, err := os.Open("day_5.txt")
	errCheck(err)
	defer input.Close()
	middleNums, rules := 0, map[string][]string{}
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()
	rule := strings.Split(inputScanner.Text(), "|")
	for len(rule) > 1 {
		rules[rule[0]] = append(rules[rule[0]], rule[1])
		inputScanner.Scan()
		rule = strings.Split(inputScanner.Text(), "|")
	}
	for inputScanner.Scan() {
		update := strings.Split(inputScanner.Text(), ",")
		if updateIsOrdered(update, rules) {
			middleNum, err := strconv.Atoi(update[len(update)/2])
			errCheck(err)
			middleNums += middleNum
		}
	}
	fmt.Println(middleNums)
}
