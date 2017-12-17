package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const table = "abcdefghijklmnop"

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := scanner.Text()

	moves := strings.Split(input, ",")

	helper := make([]int, 16)

	permutation := make([]int, 16)
	for i := 0; i < len(permutation); i++ {
		permutation[i] = i
	}

	transmutation := make([]int, 16)
	for i := 0; i < len(transmutation); i++ {
		transmutation[i] = i
	}

	for _, move := range moves {
		switch move[0] {
		case 's':
			amount := toInt(move[1:])
			for i := 0; i < amount; i++ {
				helper[i] = permutation[len(permutation) - amount + i]
			}
			for i := 0; i < len(permutation) - amount; i++ {
				helper[amount + i] = permutation[i]
			}
			permutation, helper = helper, permutation
		case 'x':
			split := strings.Split(move[1:], "/")
			a := toInt(split[0])
			b := toInt(split[1])
			permutation[a], permutation[b] = permutation[b], permutation[a]
		case 'p':
			split := strings.Split(move[1:], "/")
			a := strings.Index(table, split[0])
			b := strings.Index(table, split[1])
			ai, bi := -1, -1
			for i, t := range transmutation {
				if t == a { ai = i }
				if t == b { bi = i }
			}
			transmutation[ai], transmutation[bi] = transmutation[bi], transmutation[ai]
		}
	}

	programs := make([]int, 16)
	for i := 0; i < len(programs); i++ {
		programs[i] = i
	}

	for i, p := range permutation {
		helper[i] = programs[p]
	}

	for i, p := range helper {
		programs[i] = transmutation[p]
	}

	fmt.Println("--- Part One ---")
	for _, p := range programs {
		fmt.Print(string(table[p]))
	}
	fmt.Println()

	for i := 0; i < 1e9 - 1; i++ {
		for i, p := range permutation {
			helper[i] = programs[p]
		}

		for i, p := range helper {
			programs[i] = transmutation[p]
		}
	}

	fmt.Println("--- Part Two ---")
	for _, p := range programs {
		fmt.Print(string(table[p]))
	}
	fmt.Println()
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
