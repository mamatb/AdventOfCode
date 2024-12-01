package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := os.Open("day_1.txt")
	errCheck(err)
	defer input.Close()
	leftLocations, rightLocations := map[int]int{}, map[int]int{}
	inputScanner := bufio.NewScanner(input)
	inputScan := inputScanner.Scan()
	inputLine := inputScanner.Text()
	for inputScan {
		inputLineSplit := strings.Split(inputLine, "   ")
		leftLocation, err := strconv.Atoi(inputLineSplit[0])
		errCheck(err)
		leftLocations[leftLocation] += 1
		rightLocation, err := strconv.Atoi(inputLineSplit[1])
		errCheck(err)
		rightLocations[rightLocation] += 1
		inputScan = inputScanner.Scan()
		inputLine = inputScanner.Text()
	}
	similarityScore := 0
	for key, value := range leftLocations {
		similarityScore += key * value * rightLocations[key]
	}
	fmt.Println(similarityScore)
}
