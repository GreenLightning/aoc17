package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("--- Part One ---")
	fmt.Println(countValidPassphrases(lines, identity))
	fmt.Println("--- Part Two ---")
	fmt.Println(countValidPassphrases(lines, sorted))
}

func countValidPassphrases(lines []string, mapper func(string)string) int {
	count := 0
	for _, line := range lines {
		appearances := make(map[string]int)
		for _, word := range strings.Split(line, " ") {
			appearances[mapper(word)]++
		}
		valid := true
		for _, count := range appearances {
			if count > 1 {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func identity(a string) string {
	return a
}

func sorted(a string) string {
	letters := strings.Split(a, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}
