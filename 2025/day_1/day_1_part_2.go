package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func modulo(a int, b int) int {
	return (a%b + b) % b
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_1.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	dial, password := 50, 0
	for inputScanner.Scan() {
		rotation := inputScanner.Text()
		if rotationNum, err := strconv.Atoi(rotation[1:]); err != nil {
			panic(err)
		} else if rotation[0] == 'L' {
			dial -= rotationNum
			if dial <= 0 && dial != -rotationNum {
				password++
			}
			password += dial / -100
		} else {
			dial += rotationNum
			password += dial / 100
		}
		dial = modulo(dial, 100)
	}
	fmt.Println(password)
}
