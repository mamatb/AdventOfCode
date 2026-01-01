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

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_1.txt"); err != nil {
		panic(err)
	} else {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	}
	leftLocs, rightLocs, totalDistance := []int{}, []int{}, 0.0
	for inputScanner.Scan() {
		locations := strings.Split(inputScanner.Text(), "   ")
		if leftLocation, err := strconv.Atoi(locations[0]); err != nil {
			panic(err)
		} else {
			leftLocs = append(leftLocs, leftLocation)
		}
		if rightLocation, err := strconv.Atoi(locations[1]); err != nil {
			panic(err)
		} else {
			rightLocs = append(rightLocs, rightLocation)
		}
	}
	sort.Ints(leftLocs)
	sort.Ints(rightLocs)
	for index := range leftLocs {
		totalDistance += math.Abs(float64(leftLocs[index] - rightLocs[index]))
	}
	fmt.Println(int(totalDistance))
}
