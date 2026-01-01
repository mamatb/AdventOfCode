package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_1.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	leftLocations, rightLocations, similarity := map[int]int{}, map[int]int{}, 0
	for inputScanner.Scan() {
		locations := strings.Split(inputScanner.Text(), "   ")
		if leftLocation, err := strconv.Atoi(locations[0]); err != nil {
			panic(err)
		} else {
			leftLocations[leftLocation]++
		}
		if rightLocation, err := strconv.Atoi(locations[1]); err != nil {
			panic(err)
		} else {
			rightLocations[rightLocation]++
		}
	}
	for key, value := range leftLocations {
		similarity += key * value * rightLocations[key]
	}
	fmt.Println(similarity)
}
