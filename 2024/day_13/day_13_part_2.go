package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isPushNum(num float64) bool {
	return num > 0 && float64(int(num)) == num
}

func solveEquations(equationA []float64, equationB []float64) (float64, float64) {
	x := (equationA[1]*equationB[2] - equationB[1]*equationA[2]) /
		(equationA[1]*equationB[0] - equationB[1]*equationA[0])
	y := (equationA[2] - equationA[0]*x) / equationA[1]
	return x, y
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_13.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	opExp, err := regexp.Compile(`[^:]+: X[\+=]([0-9]+), Y[\+=]([0-9]+)`)
	if err != nil {
		panic(err)
	}
	tokens, conversionError := 0.0, 10_000_000_000_000.0
	for inputScanner.Scan() {
		equationA, equationB := []float64{}, []float64{} // Ax + By = C
		for range 3 {
			operators := opExp.FindStringSubmatch(inputScanner.Text())[1:]
			if op, err := strconv.ParseFloat(operators[0], 64); err != nil {
				panic(err)
			} else {
				equationA = append(equationA, op)
			}
			if op, err := strconv.ParseFloat(operators[1], 64); err != nil {
				panic(err)
			} else {
				equationB = append(equationB, op)
			}
			inputScanner.Scan()
		}
		equationA[2] += conversionError
		equationB[2] += conversionError
		pushesA, pushesB := solveEquations(equationA, equationB)
		if isPushNum(pushesA) && isPushNum(pushesB) {
			tokens += pushesA*3 + pushesB
		}
	}
	fmt.Println(int(tokens))
}
