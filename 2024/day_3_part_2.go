package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := os.Open("day_3.txt")
	errCheck(err)
	defer input.Close()
	mulRegexp, err := regexp.Compile(`mul\(([0-9]{1,3})\,([0-9]{1,3})\)`)
	errCheck(err)
	result, dontEnding := 0, false
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		for index, do := range strings.Split(inputScanner.Text(), "do()") {
			if index == 0 && dontEnding {
				continue
			}
			doSliced := strings.Split(do, "don't()")
			dontEnding = len(doSliced) > 1
			for _, mul := range mulRegexp.FindAllStringSubmatch(doSliced[0], -1) {
				mulNum1, err := strconv.Atoi(mul[1])
				errCheck(err)
				mulNum2, err := strconv.Atoi(mul[2])
				errCheck(err)
				result += mulNum1 * mulNum2
			}
		}
	}
	fmt.Println(result)
}
