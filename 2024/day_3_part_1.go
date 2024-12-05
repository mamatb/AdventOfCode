package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	result := 0
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		for _, mul := range mulRegexp.FindAllStringSubmatch(inputScanner.Text(), -1) {
			mulNum1, err := strconv.Atoi(mul[1])
			errCheck(err)
			mulNum2, err := strconv.Atoi(mul[2])
			errCheck(err)
			result += mulNum1 * mulNum2
		}
	}
	fmt.Println(result)
}
