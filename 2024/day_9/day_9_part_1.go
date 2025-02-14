package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func expandDiskmap(diskmap []int, input string) []int {
	fileId, free := 0, false
	for _, digitString := range strings.Split(input, "") {
		if digit, err := strconv.Atoi(digitString); err == nil {
			if free {
				for i := 0; i < digit; i++ {
					diskmap = append(diskmap, -1)
				}
			} else {
				for i := 0; i < digit; i++ {
					diskmap = append(diskmap, fileId)
				}
				fileId++
			}
		} else {
			panic(err)
		}
		free = !free
	}
	return diskmap
}

func defragDiskmap(diskmap []int) []int {
	left, right := 0, len(diskmap)-1
	for left < right {
		if diskmap[left] == -1 {
			if diskmap[right] != -1 {
				diskmap[left], diskmap[right] = diskmap[right], diskmap[left]
				left++
			}
			right--
		} else {
			left++
		}
	}
	return diskmap
}

func main() {
	var inputScanner *bufio.Scanner
	if input, err := os.Open("day_9.txt"); err == nil {
		defer input.Close()
		inputScanner = bufio.NewScanner(input)
	} else {
		panic(err)
	}
	checksums := []int{}
	for inputScanner.Scan() {
		checksum := 0
		diskmap := expandDiskmap([]int{}, inputScanner.Text())
		diskmap = defragDiskmap(diskmap)
		for fileIdx, fileId := range diskmap {
			if fileId == -1 {
				break
			}
			checksum += fileIdx * fileId
		}
		checksums = append(checksums, checksum)
	}
	fmt.Println(checksums)
}
