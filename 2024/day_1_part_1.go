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
	leftLocations, rightLocations, totalDistance := []int{}, []int{}, 0.0
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		locations := strings.Split(inputScanner.Text(), "   ")
		leftLocation, err := strconv.Atoi(locations[0])
		errCheck(err)
		leftLocations = append(leftLocations, leftLocation)
		rightLocation, err := strconv.Atoi(locations[1])
		errCheck(err)
		rightLocations = append(rightLocations, rightLocation)
	}
	sort.Ints(leftLocations)
	sort.Ints(rightLocations)
	for index := range leftLocations {
		totalDistance += math.Abs(float64(leftLocations[index] - rightLocations[index]))
	}
	fmt.Println(int(totalDistance))
}
