package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type jBox struct {
	x int
	y int
	z int
}

type jBoxPair struct {
	boxA     jBox
	boxB     jBox
	distance float64
}

func getDistance(boxA jBox, boxB jBox) float64 {
	return math.Sqrt(math.Pow(float64(boxA.x-boxB.x), 2) +
		math.Pow(float64(boxA.y-boxB.y), 2) +
		math.Pow(float64(boxA.z-boxB.z), 2))
}

func getNClosestPairs(boxes []jBox, n int) []jBoxPair {
	var pairs []jBoxPair
	for indexA, boxA := range boxes {
		for _, boxB := range boxes[indexA+1:] {
			pairs = append(pairs, jBoxPair{
				boxA:     boxA,
				boxB:     boxB,
				distance: getDistance(boxA, boxB)},
			)
		}
	}
	slices.SortFunc(pairs, func(a jBoxPair, b jBoxPair) int {
		if a.distance <= b.distance {
			return -1
		} else if a.distance > b.distance {
			return 1
		}
		return 0
	})
	return pairs[:n]
}

func getCircuitRoot(circuitRoots map[jBox]jBox, box jBox) jBox {
	if root, ok := circuitRoots[box]; !ok {
		circuitRoots[box] = box
		return box
	} else {
		for root != circuitRoots[root] {
			root = circuitRoots[root]
		}
		return root
	}
}

func getNLargestCircuitSizes(pairs []jBoxPair, n int) []int {
	circuitRoots := map[jBox]jBox{}
	for _, pair := range pairs {
		rootA := getCircuitRoot(circuitRoots, pair.boxA)
		rootB := getCircuitRoot(circuitRoots, pair.boxB)
		circuitRoots[rootA] = rootB
	}
	circuitSizes := map[jBox]int{}
	for box := range circuitRoots {
		circuitSizes[getCircuitRoot(circuitRoots, box)] += 1
	}
	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	slices.Sort(sizes)
	return sizes[len(sizes)-n:]
}

func main() {
	var inputScanner *bufio.Scanner
	if inputFile, err := os.Open("day_8.txt"); err != nil {
		panic(err)
	} else {
		defer inputFile.Close()
		inputScanner = bufio.NewScanner(inputFile)
	}
	var err error
	var boxes []jBox
	for inputScanner.Scan() {
		box, coordinates := jBox{}, strings.Split(inputScanner.Text(), ",")
		if box.x, err = strconv.Atoi(coordinates[0]); err != nil {
			panic(err)
		}
		if box.y, err = strconv.Atoi(coordinates[1]); err != nil {
			panic(err)
		}
		if box.z, err = strconv.Atoi(coordinates[2]); err != nil {
			panic(err)
		}
		boxes = append(boxes, box)
	}
	closestPairs := getNClosestPairs(boxes, 1000)
	largestCircuitSizes := getNLargestCircuitSizes(closestPairs, 3)
	fmt.Println(largestCircuitSizes[0] * largestCircuitSizes[1] * largestCircuitSizes[2])
}
