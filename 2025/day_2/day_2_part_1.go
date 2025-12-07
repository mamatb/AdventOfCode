package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_2.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	for inputScanner.Scan() {
		var err error
		var invalidSum, num, numMax int
		for _, r := range strings.Split(inputScanner.Text(), ",") {
			rSplit := strings.Split(r, "-")
			if num, err = strconv.Atoi(rSplit[0]); err != nil {
				panic(err)
			}
			if numMax, err = strconv.Atoi(rSplit[1]); err != nil {
				panic(err)
			}
			numLog10 := int(math.Log10(float64(num)))
			pow10Div := int(math.Pow10((numLog10 + 1) / 2))
			if numLog10%2 == 0 {
				num = int(math.Pow10(numLog10+1)) + pow10Div
			}
			for num <= numMax {
				numLog10 = int(math.Log10(float64(num)))
				pow10Div = int(math.Pow10((numLog10 + 1) / 2))
				left, right := num/pow10Div, num%pow10Div
				if left > right {
					num += left - right
				} else if left < right {
					num = (left+1)*pow10Div + left
				} else {
					invalidSum += num
					num = (left+1)*pow10Div + right + 1
				}
			}
		}
		fmt.Println(invalidSum)
	}
}
