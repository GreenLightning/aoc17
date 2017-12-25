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

	rules := make(map[string]rule)

	scanner.Scan()
	initialState := regexp.MustCompile(`Begin in state ([A-Z]).`).FindStringSubmatch(scanner.Text())[1]

	scanner.Scan()
	totalSteps := toInt(regexp.MustCompile(`Perform a diagnostic checksum after (\d+) steps.`).FindStringSubmatch(scanner.Text())[1])

	statePattern    := regexp.MustCompile(`In state ([A-Z]):`)
	valuePattern    := regexp.MustCompile(`If the current value is (\d):`)
	writePattern    := regexp.MustCompile(`Write the value (\d).`)
	movePattern     := regexp.MustCompile(`Move one slot to the (right|left).`)
	continuePattern := regexp.MustCompile(`Continue with state ([A-Z]).`)

	scanner.Scan() // empty line

	for scanner.Scan() {
		s := statePattern.FindStringSubmatch(scanner.Text())[1]

		parts := [2]part{}

		for i := 0; i < 2; i++ {
			scanner.Scan()
			p := toInt(valuePattern.FindStringSubmatch(scanner.Text())[1])
			scanner.Scan()
			value := toInt(writePattern.FindStringSubmatch(scanner.Text())[1])
			scanner.Scan()
			offset := 0
			switch movePattern.FindStringSubmatch(scanner.Text())[1] {
				case "left": offset = -1
				case "right": offset = 1
			}
			scanner.Scan()
			state := continuePattern.FindStringSubmatch(scanner.Text())[1]
			parts[p] = part{ value, offset, state }
		}

		rules[s] = rule{ parts }

		scanner.Scan() // empty line
	}

	tape := make(map[int]int)
	cursor := 0
	state := initialState

	for steps := 0; steps < totalSteps; steps++ {
		part := rules[state].parts[tape[cursor]]
		tape[cursor] = part.value
		cursor += part.offset
		state = part.state
	}

	diagnosticChecksum := 0
	for _, v := range tape {
		if v == 1 {
			diagnosticChecksum++
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(diagnosticChecksum)
}

type rule struct {
	parts [2]part
}

type part struct {
	value, offset int
	state string
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
