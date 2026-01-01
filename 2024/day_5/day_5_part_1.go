package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_5.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	middleNums, rules := 0, map[string][]string{}
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
			if middleNum, err := strconv.Atoi(update[len(update)/2]); err != nil {
				panic(err)
			} else {
				middleNums += middleNum
			}
		}
	}
	fmt.Println(middleNums)
}
