package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func updateOrdered(update []string, rules map[string][]string) []string {
	pagesPrev := map[string]bool{}
	for pageIdx, page := range update {
		for _, pageRule := range rules[page] {
			if pagesPrev[pageRule] {
				update = append(update[:pageIdx], update[pageIdx+1:]...)
				pageIdx = slices.Index(update, pageRule)
				update = slices.Insert(update, pageIdx, page)
				return updateOrdered(update, rules)
			}
		}
		pagesPrev[page] = true
	}
	return update
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_5.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
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
		if !updateIsOrdered(update, rules) {
			update = updateOrdered(update, rules)
			if middleNum, err := strconv.Atoi(update[len(update)/2]); err == nil {
				middleNums += middleNum
			} else {
				panic(err)
			}
		}
	}
	fmt.Println(middleNums)
}
