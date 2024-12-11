package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type diskFile struct {
	fileId int
	index  int
	size   int
}

type diskSpace struct {
	index int
	size  int
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func expandDiskmap(diskmap []int, diskFiles []diskFile, diskSpaces []diskSpace,
	input string) ([]int, []diskFile, []diskSpace) {
	fileId, free := 0, false
	for _, digitString := range strings.Split(input, "") {
		digit, err := strconv.Atoi(digitString)
		errCheck(err)
		if free {
			diskSpaces = append(diskSpaces,
				diskSpace{index: len(diskmap), size: digit})
			for i := 0; i < digit; i++ {
				diskmap = append(diskmap, -1)
			}
		} else {
			diskFiles = append(diskFiles,
				diskFile{fileId: fileId, index: len(diskmap), size: digit})
			for i := 0; i < digit; i++ {
				diskmap = append(diskmap, fileId)
			}
			fileId += 1
		}
		free = !free
	}
	slices.Reverse(diskFiles)
	return diskmap, diskFiles, diskSpaces
}

func defragDiskmap(diskmap []int, diskFiles []diskFile, diskSpaces []diskSpace) []int {
	for _, file := range diskFiles {
		for spaceIndex, space := range diskSpaces {
			if space.index > file.index {
				break
			}
			if file.size == space.size {
				for i := 0; i < file.size; i++ {
					diskmap[space.index+i] = diskmap[file.index+i]
					diskmap[file.index+i] = -1
				}
				diskSpaces = append(diskSpaces[:spaceIndex], diskSpaces[spaceIndex+1:]...)
				break
			}
			if file.size < space.size {
				for i := 0; i < file.size; i++ {
					diskmap[space.index+i] = diskmap[file.index+i]
					diskmap[file.index+i] = -1
				}
				diskSpaces[spaceIndex] = diskSpace{index: space.index + file.size,
					size: space.size - file.size}
				break
			}
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
		diskmap, diskFiles, diskSpaces := expandDiskmap(
			[]int{}, []diskFile{}, []diskSpace{}, inputScanner.Text())
		diskmap = defragDiskmap(diskmap, diskFiles, diskSpaces)
		for fileIndex, fileId := range diskmap {
			if fileId != -1 {
				checksum += fileIndex * fileId
			}
		}
		checksums = append(checksums, checksum)
	}
	fmt.Println(checksums)
}
