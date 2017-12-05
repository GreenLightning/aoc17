package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	input := make([]int, 0)

	for scanner.Scan() {
		input = append(input, toInt(scanner.Text()))
	}

	fmt.Println("--- Part One ---")
	fmt.Println(simulate(input, func(offset int) int {
		return 1
	}))

	fmt.Println("--- Part Two ---")
	fmt.Println(simulate(input, func(offset int) int {
		if offset >= 3 {
			return -1
		} else {
			return 1
		}
	}))
}

func simulate(input []int, change func(int)int) int {
	offsets := make([]int, len(input))
	copy(offsets, input)
	steps, pointer := 0, 0
	for pointer >= 0 && pointer < len(offsets) {
		offset := offsets[pointer]
		offsets[pointer] += change(offset)
		pointer += offset
		steps++
	}
	return steps
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
