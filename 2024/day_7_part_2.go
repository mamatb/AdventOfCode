package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func equationExists(result int, operators []int) bool {
	resultsTemp, resultsTempLen := make([]int, 1), 1
	resultsTemp[0], operators = operators[0], operators[1:]
	var operator, resultTemp int
	for len(operators) > 0 && resultsTempLen > 0 {
		operator, operators = operators[0], operators[1:]
		for i := 0; i < resultsTempLen; i++ {
			resultTemp, resultsTemp = resultsTemp[0], resultsTemp[1:]
			resultAdd, resultMul, resultCon := resultTemp, resultTemp, resultTemp
			resultAdd += operator
			if resultAdd <= result {
				resultsTemp = append(resultsTemp, resultAdd)
			}
			resultMul *= operator
			if resultMul <= result {
				resultsTemp = append(resultsTemp, resultMul)
			}
			resultCon *= int(math.Pow10(int(math.Log10(float64(operator))) + 1))
			resultCon += operator
			if resultCon <= result {
				resultsTemp = append(resultsTemp, resultCon)
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
	input, err := os.Open("day_7.txt")
	errCheck(err)
	defer input.Close()
	resultCalibration := 0
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		equation := strings.Split(inputScanner.Text(), ": ")
		result, err := strconv.Atoi(equation[0])
		errCheck(err)
		operators := []int{}
		for _, numString := range strings.Split(equation[1], " ") {
			num, err := strconv.Atoi(numString)
			errCheck(err)
			operators = append(operators, num)
		}
		if equationExists(result, operators) {
			resultCalibration += result
		}
	}
	fmt.Println(resultCalibration)
}
