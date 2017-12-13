package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	depthToSizes := make(map[int]int)
	totalDepth := -1

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ": ")
		depth, size := toInt(values[0]), toInt(values[1])
		depthToSizes[depth] = size
		if depth > totalDepth { totalDepth = depth }
	}

	sizes := make([]int, totalDepth + 1)
	for depth, size := range depthToSizes {
		sizes[depth] = size
	}

	periods := make([]int, totalDepth + 1)
	for depth, size := range depthToSizes {
		periods[depth] = 2 * (size - 1)
	}

	fmt.Println("--- Part One ---")

	severity := 0

	for depth, period := range periods {
		if period != 0 && depth % period == 0 {
			severity += depth * sizes[depth]
		}
	}

	fmt.Println(severity)

	fmt.Println("--- Part Two ---")

	delay := 0

	loop: for {
		for depth, period := range periods {
			if period != 0 && (delay + depth) % period == 0 {
				delay++
				continue loop
			}
		}
		break
	}

	fmt.Println(delay)
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
