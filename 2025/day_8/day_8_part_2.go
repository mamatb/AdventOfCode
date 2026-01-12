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

func getFarthestPairs(boxes []jBox) []jBoxPair {
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
		if a.distance > b.distance {
			return -1
		} else if a.distance < b.distance {
			return 1
		}
		return 0
	})
	return pairs
}

func getCircuitRoot(circuitRoots map[jBox]jBox, box jBox) jBox {
	root := circuitRoots[box]
	for root != circuitRoots[root] {
		root = circuitRoots[root]
	}
	return root
}

func getLastPairToNCircuits(boxes []jBox, farthestPairs []jBoxPair, n int) jBoxPair {
	circuitRoots, circuitRootsCount := map[jBox]jBox{}, len(boxes)
	for _, box := range boxes {
		circuitRoots[box] = box
	}
	var pair jBoxPair
	for len(farthestPairs) > 0 && circuitRootsCount > n {
		pair = farthestPairs[len(farthestPairs)-1]
		farthestPairs = farthestPairs[:len(farthestPairs)-1]
		rootA := getCircuitRoot(circuitRoots, pair.boxA)
		rootB := getCircuitRoot(circuitRoots, pair.boxB)
		if rootA != rootB {
			circuitRoots[rootA] = rootB
			circuitRootsCount--
		}
	}
	if circuitRootsCount > n {
		return jBoxPair{}
	}
	return pair
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
	farthestPairs := getFarthestPairs(boxes)
	lastPair := getLastPairToNCircuits(boxes, farthestPairs, 1)
	fmt.Println(lastPair.boxA.x * lastPair.boxB.x)
}
