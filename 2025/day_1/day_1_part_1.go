package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_1.txt"); err == nil {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	} else {
		panic(err)
	}
	dial, password := 50, 0
	for inputScanner.Scan() {
		rotation := inputScanner.Text()
		if rotationNum, err := strconv.Atoi(rotation[1:]); err != nil {
			panic(err)
		} else if rotation[0] == 'L' {
			dial -= rotationNum
		} else {
			dial += rotationNum
		}
		if dial%100 == 0 {
			password++
		}
	}
	fmt.Println(password)
}
