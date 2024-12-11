package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func expandDiskmap(diskmap []int, input string) []int {
	fileId, free := 0, false
	for _, digitString := range strings.Split(input, "") {
		digit, err := strconv.Atoi(digitString)
		errCheck(err)
		if free {
			for i := 0; i < digit; i++ {
				diskmap = append(diskmap, -1)
			}
		} else {
			for i := 0; i < digit; i++ {
				diskmap = append(diskmap, fileId)
			}
			fileId += 1
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
				left += 1
			}
			right -= 1
		} else {
			left += 1
		}
	}
	return diskmap
}

func main() {
	input, err := os.Open("day_9.txt")
	errCheck(err)
	defer input.Close()
	checksums := []int{}
	inputScanner := bufio.NewScanner(input)
	for inputScanner.Scan() {
		checksum := 0
		diskmap := expandDiskmap([]int{}, inputScanner.Text())
		diskmap = defragDiskmap(diskmap)
		for fileIndex, fileId := range diskmap {
			if fileId == -1 {
				break
			}
			checksum += fileIndex * fileId
		}
		checksums = append(checksums, checksum)
	}
	fmt.Println(checksums)
}
