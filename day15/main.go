package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
)

const factorA = 16807
const factorB = 48271

const mod = (uint(1) << 31) - 1

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	inputA := scanner.Text()

	scanner.Scan()
	inputB := scanner.Text()

	rule := regexp.MustCompile(`Generator (A|B) starts with (\d+)`)

	resultsA := rule.FindStringSubmatch(inputA)
	resultsB := rule.FindStringSubmatch(inputB)

	startA := toUint(resultsA[2])
	startB := toUint(resultsB[2])

	{
		fmt.Println("--- Part One ---")
		count := 0
		valueA, valueB := startA, startB
		for i := 0; i < 40e6; i++ {
			if valueA & 0xffff == valueB & 0xffff { count++ }
			valueA = (valueA * factorA) % mod
			valueB = (valueB * factorB) % mod
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		count := 0
		valueA, valueB := startA, startB
		for i := 0; i < 5e6; i++ {
			for valueA % 4 != 0 {
				valueA = (valueA * factorA) % mod
			}
			for valueB % 8 != 0 {
				valueB = (valueB * factorB) % mod
			}
			if valueA & 0xffff == valueB & 0xffff { count++ }
			valueA = (valueA * factorA) % mod
			valueB = (valueB * factorB) % mod
		}
		fmt.Println(count)
	}
}

func toUint(v string) uint {
	result, err := strconv.ParseUint(v, 10, 0)
	if err != nil { panic(err) }
	return uint(result)
}
