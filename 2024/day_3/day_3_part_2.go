package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_3.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	mulExp, err := regexp.Compile(`mul\(([0-9]{1,3})\,([0-9]{1,3})\)`)
	if err != nil {
		panic(err)
	}
	result, dontEnding := 0, false
	for inputScanner.Scan() {
		for doIdx, do := range strings.Split(inputScanner.Text(), "do()") {
			if doIdx == 0 && dontEnding {
				continue
			}
			doSliced := strings.Split(do, "don't()")
			dontEnding = len(doSliced) > 1
			for _, mul := range mulExp.FindAllStringSubmatch(doSliced[0], -1) {
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
	}
	fmt.Println(result)
}
