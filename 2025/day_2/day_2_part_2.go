package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkPattern(num int, size int) bool {
	numLog10, pow10Div := int(math.Log10(float64(num))), int(math.Pow10(size))
	if (numLog10+1)%size != 0 {
		return false
	}
	pending, chunk := num/pow10Div, num%pow10Div
	for pending > 0 {
		if pending%pow10Div != chunk {
			return false
		}
		pending /= pow10Div
	}
	return pending == 0
}

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
			for num <= numMax {
				for size := 1; int(math.Pow10(size)) < num; size++ {
					if checkPattern(num, size) {
						invalidSum += num
						break
					}
				}
				num++
			}
		}
		fmt.Println(invalidSum)
	}
}
