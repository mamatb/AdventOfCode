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
	leftLocations, rightLocations, similarity := map[int]int{}, map[int]int{}, 0
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		locations := strings.Split(inputScanner.Text(), "   ")
		leftLocation, err := strconv.Atoi(locations[0])
		errCheck(err)
		leftLocations[leftLocation] += 1
		rightLocation, err := strconv.Atoi(locations[1])
		errCheck(err)
		rightLocations[rightLocation] += 1
	}
	for key, value := range leftLocations {
		similarity += key * value * rightLocations[key]
	}
	fmt.Println(similarity)
}
