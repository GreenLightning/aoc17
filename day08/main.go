package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instrRule := regexp.MustCompile(`(\w+) (dec|inc) (-?\d+) if (\w+) (==|!=|<=|>=|<|>) (-?\d+)`)

	regs := make(map[string]int)

	totalMax := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := instrRule.FindStringSubmatch(line)

		actionReg, action, actionValue := matches[1], matches[2], toInt(matches[3])
		testReg, test, testValue := matches[4], matches[5], toInt(matches[6])

		testResult := false
		switch test {
		case "==":
			testResult = (regs[testReg] == testValue)
		case "!=":
			testResult = (regs[testReg] != testValue)
		case "<=":
			testResult = (regs[testReg] <= testValue)
		case ">=":
			testResult = (regs[testReg] >= testValue)
		case "<":
			testResult = (regs[testReg] < testValue)
		case ">":
			testResult = (regs[testReg] > testValue)
		}

		if testResult {
			if action == "dec" {
				actionValue = -actionValue
			}
			actionResult := regs[actionReg] + actionValue
			regs[actionReg] = actionResult
			if actionResult > totalMax { totalMax = actionResult }
		}
	}

	finalMax := 0
	for _, value := range regs {
		if value > finalMax {
			finalMax = value
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(finalMax)

	fmt.Println("--- Part Two ---")
	fmt.Println(totalMax)
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
