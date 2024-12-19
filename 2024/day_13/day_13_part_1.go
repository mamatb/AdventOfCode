package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isPushNum(num float64) bool {
	return num > 0 && num < 100 && float64(int(num)) == num
}

func solveEquations(equationA []float64, equationB []float64) (float64, float64) {
	x := (equationA[1]*equationB[2] - equationB[1]*equationA[2]) /
		(equationA[1]*equationB[0] - equationB[1]*equationA[0])
	y := (equationA[2] - equationA[0]*x) / equationA[1]
	return x, y
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_13.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	opExp, err := regexp.Compile(`[^:]+: X[\+=]([0-9]+), Y[\+=]([0-9]+)`)
	if err != nil {
		panic(err)
	}
	tokens := 0.0
	for inputScanner.Scan() {
		equationA, equationB := []float64{}, []float64{} // Ax + By = C
		for i := 0; i < 3; i++ {
			operators := opExp.FindStringSubmatch(inputScanner.Text())[1:]
			if op, err := strconv.ParseFloat(operators[0], 64); err == nil {
				equationA = append(equationA, op)
			} else {
				panic(err)
			}
			if op, err := strconv.ParseFloat(operators[1], 64); err == nil {
				equationB = append(equationB, op)
			} else {
				panic(err)
			}
			inputScanner.Scan()
		}
		pushesA, pushesB := solveEquations(equationA, equationB)
		if isPushNum(pushesA) && isPushNum(pushesB) {
			tokens += pushesA*3 + pushesB
		}
	}
	fmt.Println(int(tokens))
}
