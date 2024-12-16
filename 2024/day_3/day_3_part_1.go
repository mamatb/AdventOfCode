package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_3.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	mulRegexp, err := regexp.Compile(`mul\(([0-9]{1,3})\,([0-9]{1,3})\)`)
	if err != nil {
		panic(err)
	}
	result := 0
	for inputScanner.Scan() {
		for _, mul := range mulRegexp.FindAllStringSubmatch(inputScanner.Text(), -1) {
			mulNum := 1
			if mulNum1, err := strconv.Atoi(mul[1]); err == nil {
				mulNum *= mulNum1
			} else {
				panic(err)
			}
			if mulNum2, err := strconv.Atoi(mul[2]); err == nil {
				mulNum *= mulNum2
			} else {
				panic(err)
			}
			result += mulNum
		}
	}
	fmt.Println(result)
}
