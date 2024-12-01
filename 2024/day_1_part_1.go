package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	leftLocations, rightLocations := []int{}, []int{}
	inputScanner := bufio.NewScanner(input)
	inputScan := inputScanner.Scan()
	inputLine := inputScanner.Text()
	for inputScan {
		inputLineSplit := strings.Split(inputLine, "   ")
		leftLocation, err := strconv.Atoi(inputLineSplit[0])
		errCheck(err)
		leftLocations = append(leftLocations, leftLocation)
		rightLocation, err := strconv.Atoi(inputLineSplit[1])
		errCheck(err)
		rightLocations = append(rightLocations, rightLocation)
		inputScan = inputScanner.Scan()
		inputLine = inputScanner.Text()
	}
	sort.Ints(leftLocations)
	sort.Ints(rightLocations)
	totalDistance := 0.0
	for i := range leftLocations {
		totalDistance += math.Abs(float64(leftLocations[i] - rightLocations[i]))
	}
	fmt.Println(int(totalDistance))
}
