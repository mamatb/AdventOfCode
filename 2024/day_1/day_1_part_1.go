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
	if input, err := os.Open("day_1.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	leftLocs, rightLocs, totalDistance := []int{}, []int{}, 0.0
	for inputScanner.Scan() {
		locations := strings.Split(inputScanner.Text(), "   ")
		if leftLocation, err := strconv.Atoi(locations[0]); err == nil {
			leftLocs = append(leftLocs, leftLocation)
		} else {
			panic(err)
		}
		if rightLocation, err := strconv.Atoi(locations[1]); err == nil {
			rightLocs = append(rightLocs, rightLocation)
		} else {
			panic(err)
		}
	}
	sort.Ints(leftLocs)
	sort.Ints(rightLocs)
	for index := range leftLocs {
		totalDistance += math.Abs(float64(leftLocs[index] - rightLocs[index]))
	}
	fmt.Println(int(totalDistance))
}
