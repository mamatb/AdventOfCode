package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(stones []string) []string {
	stonesBlink := []string{}
	for _, stone := range stones {
		if stone == "0" {
			stonesBlink = append(stonesBlink, "1")
		} else if len(stone)%2 == 0 {
			stoneLeft := strings.TrimLeft(stone[:len(stone)/2], "0")
			if len(stoneLeft) == 0 {
				stoneLeft = "0"
			}
			stoneRight := strings.TrimLeft(stone[len(stone)/2:], "0")
			if len(stoneRight) == 0 {
				stoneRight = "0"
			}
			stonesBlink = append(stonesBlink, stoneLeft, stoneRight)
		} else {
			if stoneInt, err := strconv.Atoi(stone); err == nil {
				stonesBlink = append(stonesBlink, strconv.Itoa(stoneInt*2024))
			} else {
				panic(err)
			}
		}
	}
	return stonesBlink
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_11.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	blinks, stonesLen := 25, []int{}
	for inputScanner.Scan() {
		stones := strings.Split(inputScanner.Text(), " ")
		for range blinks {
			stones = blink(stones)
		}
		stonesLen = append(stonesLen, len(stones))
	}
	fmt.Println(stonesLen)
}
