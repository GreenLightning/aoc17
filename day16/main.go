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

	permutation, transmutation := parseMoves(moves)

	programs, helper := makeSlice(), makeSlice()

	fmt.Println("--- Part One ---")
	applyOnce(programs, helper, permutation, transmutation)
	printPrograms(programs)

	fmt.Println("--- Part Two ---")
	applyCount(programs, helper, permutation, transmutation, 1, 1e9 - 1)
	printPrograms(programs)
}

func makeSlice() []int {
	slice := make([]int, 16)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	return slice
}

func parseMoves(moves []string) ([]int, []int) {
	permutation, transmutation, helper := makeSlice(), makeSlice(), makeSlice()
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
	return permutation, transmutation
}

func applyCount(programs, helper, permutation, transmutation []int, power, count int) int {
	if 2 * power <= count {
		doublePermutation := makeSlice()
		for i, p := range permutation {
			helper[i] = doublePermutation[p]
		}
		for i, p := range permutation {
			doublePermutation[i] = helper[p]
		}

		doubleTransmutation := makeSlice()
		for step := 0; step < 2; step++ {
			for i, p := range doubleTransmutation {
				doubleTransmutation[i] = transmutation[p]
			}
		}

		count = applyCount(programs, helper, doublePermutation, doubleTransmutation, 2 * power, count)
	}

	for ; power <= count; count -= power {
		applyOnce(programs, helper, permutation, transmutation)
	}

	return count
}

func applyOnce(programs, helper, permutation, transmutation []int) {
	for i, p := range permutation {
		helper[i] = programs[p]
	}
	for i, p := range helper {
		programs[i] = transmutation[p]
	}
}

func printPrograms(programs []int) {
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
