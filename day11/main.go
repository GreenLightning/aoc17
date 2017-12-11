package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	x, y, z := 0, 0, 0
	max := 0

	for _, direction := range input {
		switch direction {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			z++
			y--
		case "sw":
			z++
			x--
		case "nw":
			y++
			x--
		default:
			panic(direction)
		}
		current := dist(x, y, z)
		if current > max { max = current }
	}

	fmt.Println("--- Part One ---")
	fmt.Println(dist(x, y, z))

	fmt.Println("--- Part Two ---")
	fmt.Println(max)
}

func dist(x, y, z int) int {
	return (abs(x) + abs(y) + abs(z)) / 2
}

func abs(v int) int {
	if v < 0 { return -v }
	return v
}
