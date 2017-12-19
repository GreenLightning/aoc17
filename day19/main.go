package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	path, steps := "", 0

	x, y := 0, 0

	for input[y][x] != '|' {
		x++
	}

	dx, dy := 0, 1

	for input[y][x] != ' ' {
		if input[y][x] >= 'A' && input[y][x] <= 'Z' {
			path += string([]byte{ input[y][x] })
		} else if input[y][x] == '+' {
			if dx != 0 {
				if input[y+1][x] != ' ' {
					dx, dy = 0, 1
				} else if input[y-1][x] != ' ' {
					dx, dy = 0, -1
				} else {
					panic(fmt.Sprintf("(%d, %d)", x, y))
				}
			} else {
				if input[y][x+1] != ' ' {
					dx, dy = 1, 0
				} else if input[y][x-1] != ' ' {
					dx, dy = -1, 0
				} else {
					panic(fmt.Sprintf("(%d, %d)", x, y))
				}
			}
		}
		x += dx
		y += dy
		steps++
	}

	fmt.Println("--- Part One ---")
	fmt.Println(path)

	fmt.Println("--- Part Two ---")
	fmt.Println(steps)
}
