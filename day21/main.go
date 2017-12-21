package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rules2 := make(map[string]string)
	rules3 := make(map[string]string)

	regex2 := regexp.MustCompile(`([.#]{2})/([.#]{2}) => ([.#]{3})/([.#]{3})/([.#]{3})`)
	regex3 := regexp.MustCompile(`([.#]{3})/([.#]{3})/([.#]{3}) => ([.#]{4})/([.#]{4})/([.#]{4})/([.#]{4})`)

	for scanner.Scan() {
		line := scanner.Text()
		if results := regex2.FindStringSubmatch(line); results != nil {
			input := []string{ results[1], results[2] }
			output := results[3] + results[4] + results[5]
			for _, permutation := range permute(input) {
				rules2[permutation] = output
			}
		} else if results := regex3.FindStringSubmatch(line); results != nil {
			input := []string{ results[1], results[2], results[3] }
			output := results[4] + results[5] + results[6] + results[7]
			for _, permutation := range permute(input) {
				rules3[permutation] = output
			}
		} else {
			panic(line)
		}
	}

	pattern := []string{ ".#.", "..#", "###" }

	fmt.Println("--- Part One ---")
	pattern = transformIterations(pattern, 5, rules2, rules3)
	fmt.Println(countPixels(pattern))

	fmt.Println("--- Part Two ---")
	pattern = transformIterations(pattern, 13, rules2, rules3)
	fmt.Println(countPixels(pattern))
}

func permute(input []string) []string {
	result := make([]string, 0)

	for i := 0; i < 4; i++ {
		{
			value := ""
			for _, str := range input {
				value += str
			}
			result = append(result, value)
		}

		{
			value := ""
			for i := len(input) - 1; i >= 0; i-- {
				value += input[i]
			}
			result = append(result, value)
		}

		{
			value := ""
			for _, str := range input {
				runes := []rune(str)
				for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
					runes[i], runes[j] = runes[j], runes[i]
				}
				value += string(runes)
			}
			result = append(result, value)
		}

		input = rotate(input)
	}
	return result
}

func rotate(input []string) []string {
	result := make([]string, len(input))
	for i := 0; i < len(result); i++ {
		value := ""
		for j := 0; j < len(input); j++ {
			value += input[j][len(input) - i - 1:len(input) - i]
		}
		result[i] = value
	}
	return result
}

func transformIterations(pattern []string, iterations int, rules2 map[string]string, rules3 map[string]string) []string {
	for iter := 0; iter < iterations; iter++ {
		if len(pattern) % 2 == 0 {
			pattern = transformSize(pattern, 2, rules2)
		} else {
			pattern = transformSize(pattern, 3, rules3)
		}
	}
	return pattern
}

func transformSize(pattern []string, size int, rules map[string]string) []string {
	result := make([]string, 0)
	for y := 0; y < len(pattern); y += size {
		current := make([]string, size + 1)
		for x := 0; x < len(pattern); x += size {
			input := ""
			for i := 0; i < size; i++ {
				input += pattern[y+i][x:x+size]
			}
			output := rules[input]
			for i := 0; i < size + 1; i++ {
				current[i] += output[(size + 1) * i : (size + 1) * (i + 1)]
			}
		}
		result = append(result, current...)
	}
	return result
}

func countPixels(pattern []string) int {
	count := 0
	for _, str := range pattern {
		for _, chr := range str {
			if chr == '#' {
				count++
			}
		}
	}
	return count
}
