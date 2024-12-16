package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(stones map[string]int) map[string]int {
	stonesBlink := map[string]int{}
	for stone, stoneCount := range stones {
		if stone == "0" {
			stonesBlink["1"] += stoneCount
		} else if len(stone)%2 == 0 {
			stoneLeft := strings.TrimLeft(stone[:len(stone)/2], "0")
			if len(stoneLeft) == 0 {
				stoneLeft = "0"
			}
			stoneRight := strings.TrimLeft(stone[len(stone)/2:], "0")
			if len(stoneRight) == 0 {
				stoneRight = "0"
			}
			stonesBlink[stoneLeft] += stoneCount
			stonesBlink[stoneRight] += stoneCount
		} else {
			if stoneInt, err := strconv.Atoi(stone); err == nil {
				stonesBlink[strconv.Itoa(stoneInt*2024)] += stoneCount
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
	blinks, stonesLen := 75, []int{}
	for inputScanner.Scan() {
		stonesLen = append(stonesLen, 0)
		stones := map[string]int{}
		for _, stone := range strings.Split(inputScanner.Text(), " ") {
			stones[stone] += 1
		}
		for i := 0; i < blinks; i++ {
			stones = blink(stones)
		}
		for _, stoneCount := range stones {
			stonesLen[len(stonesLen)-1] += stoneCount
		}
	}
	fmt.Println(stonesLen)
}
