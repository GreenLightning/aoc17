package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	stateClean    = 0
	stateWeakened = 1
	stateInfected = 2
	stateFlagged  = 3
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	input := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	{
		fmt.Println("--- Part One ---")

		grid := make(map[position]bool)

		y := -(len(input) / 2)
		for _, line := range input {
			x := -(len(line) / 2)
			for _, char := range line {
				if char == '#' {
					grid[position{ x, y }] = true
				}
				x++
			}
			y++
		}

		pos := position{ 0, 0 }
		dir := 0
		count := 0

		for i := 0; i < 10000; i++ {
			infected := grid[pos]
			if infected {
				dir = (dir + 1) % 4
			} else {
				dir = (dir + 3) % 4
			}
			if !infected {
				count++
			}
			grid[pos] = !infected
			switch dir {
				case 0: pos.y--
				case 1: pos.x++
				case 2: pos.y++
				case 3: pos.x--
			}
		}

		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		grid := make(map[position]int)

		y := -(len(input) / 2)
		for _, line := range input {
			x := -(len(line) / 2)
			for _, char := range line {
				if char == '#' {
					grid[position{ x, y }] = stateInfected
				}
				x++
			}
			y++
		}

		pos := position{ 0, 0 }
		dir := 0
		count := 0

		for i := 0; i < 10000000; i++ {
			state := grid[pos]
			switch state {
				case stateClean:    dir = (dir + 3) % 4
				case stateInfected: dir = (dir + 1) % 4
				case stateFlagged:  dir = (dir + 2) % 4
			}
			if state == stateWeakened {
				count++
			}
			grid[pos] = (state + 1) % 4
			switch dir {
				case 0: pos.y--
				case 1: pos.x++
				case 2: pos.y++
				case 3: pos.x--
			}
		}

		fmt.Println(count)
	}
}

type position struct {
	x, y int
}
