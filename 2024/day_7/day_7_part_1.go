package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func equationExists(result int, operators []int) bool {
	resultsTemp, resultsTempLen := make([]int, 1), 1
	resultsTemp[0], operators = operators[0], operators[1:]
	var operator, resultTemp int
	for len(operators) > 0 && resultsTempLen > 0 {
		operator, operators = operators[0], operators[1:]
		for range resultsTempLen {
			resultTemp, resultsTemp = resultsTemp[0], resultsTemp[1:]
			resultAdd, resultMul := resultTemp, resultTemp
			resultAdd += operator
			if resultAdd <= result {
				resultsTemp = append(resultsTemp, resultAdd)
			}
			resultMul *= operator
			if resultMul <= result {
				resultsTemp = append(resultsTemp, resultMul)
			}
		}
		resultsTempLen = len(resultsTemp)
	}
	for _, resultTemp := range resultsTemp {
		if resultTemp == result {
			return true
		}
	}
	return false
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_7.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	resultCalibration := 0
	for inputScanner.Scan() {
		equation, operators := strings.Split(inputScanner.Text(), ": "), []int{}
		for _, numStr := range strings.Split(equation[1], " ") {
			if num, err := strconv.Atoi(numStr); err == nil {
				operators = append(operators, num)
			} else {
				panic(err)
			}
		}
		if result, err := strconv.Atoi(equation[0]); err == nil {
			if equationExists(result, operators) {
				resultCalibration += result
			}
		} else {
			panic(err)
		}

	}
	fmt.Println(resultCalibration)
}
