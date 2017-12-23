package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := make([]instruction, 0)

	for scanner.Scan() {
		line := scanner.Text()
		i := parseInstruction(line)
		instructions = append(instructions, i)
	}

	{
		fmt.Println("--- Part One ---")
		m := machine{}
		m.run(instructions)
		fmt.Println(m.count)
	}

	{
		fmt.Println("--- Part Two ---")

		//
		// expected input format:
		// 0: set b <int>
		// ...
		// 4: mul b <int>
		// 5: sub b <int>
		// ...
		// 7: sub c <int>
		// ...
		// 30: sub b <int>
		// ...
		//

		if instructions[ 0].kind != instrKindSet || !instructions[ 0].x.isReg || instructions[ 0].x.value != 1 || instructions[ 0].y.isReg ||
		   instructions[ 4].kind != instrKindMul || !instructions[ 4].x.isReg || instructions[ 4].x.value != 1 || instructions[ 4].y.isReg ||
		   instructions[ 5].kind != instrKindSub || !instructions[ 5].x.isReg || instructions[ 5].x.value != 1 || instructions[ 5].y.isReg ||
		   instructions[ 7].kind != instrKindSub || !instructions[ 7].x.isReg || instructions[ 7].x.value != 2 || instructions[ 7].y.isReg ||
		   instructions[30].kind != instrKindSub || !instructions[30].x.isReg || instructions[30].x.value != 1 || instructions[30].y.isReg {
			panic("input not in expected format")
		}

		b := instructions[0].y.value * instructions[4].y.value - instructions[5].y.value
		c := b - instructions[7].y.value
		delta := -instructions[30].y.value

		h := 0

		loop: for ; b <= c; b += delta {
			for d := 2; d != b; d++ {
				if b % d == 0 {
					h++
					continue loop
				}
			}
		}

		fmt.Println(h)
	}
}

const (
	instrKindSet = iota
	instrKindSub
	instrKindMul
	instrKindJnz
)

type argument struct {
	isReg bool
	value int
}

type instruction struct {
	kind int
	x argument
	y argument
}

func split(line string) (string, string) {
	index := strings.Index(line, " ")
	if index == -1 {
		return line, ""
	} else {
		return line[:index], line[index+1:]
	}
}

func parseArg(line string) (argument, string) {
	value, line := split(line)
	if len(value) == 1 && value[0] >= 'a' && value[0] <= 'z' {
		return argument{ true, int(value[0] - 'a') }, line
	} else {
		return argument{ false, toInt(value) }, line
	}
}

func parseRegArg(line string) (argument, string) {
	arg, line := parseArg(line)
	if !arg.isReg {
		panic("expected register but found literal value")
	}
	return arg, line
}

func parseInstruction(line string) instruction {
	if strings.HasPrefix(line, "set ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindSet, x, y }
	} else if strings.HasPrefix(line, "sub ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindSub, x, y }
	} else if strings.HasPrefix(line, "mul ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindMul, x, y }
	} else if strings.HasPrefix(line, "jnz ") {
		line = line[4:]
		x, line := parseArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindJnz, x, y }
	} else {
		panic(fmt.Sprint("cannot parse instruction: ", line))
	}
}

type machine struct {
	registers [26]int
	count int
}

func (m *machine) run(instructions []instruction) {
	ip := 0
	for ip >= 0 && ip < len(instructions) {
		i := instructions[ip]
		ip += m.execute(i)
	}
}

func (m *machine) execute(i instruction) int {
	switch i.kind {
		case instrKindSet:
			reg := m.getReg(i.x)
			*reg = m.getVal(i.y)
			return 1
		case instrKindSub:
			reg := m.getReg(i.x)
			*reg -= m.getVal(i.y)
			return 1
		case instrKindMul:
			m.count++
			reg := m.getReg(i.x)
			*reg *= m.getVal(i.y)
			return 1
		case instrKindJnz:
			result := m.getVal(i.x)
			if result != 0 {
				return m.getVal(i.y)
			} else {
				return 1
			}
		default: panic(fmt.Sprintf("unknown instruction kind %d", i.kind))
	}
}

func (m *machine) getReg(arg argument) *int {
	if arg.isReg {
		return &m.registers[arg.value]
	} else {
		return nil
	}
}

func (m *machine) getVal(arg argument) int {
	if arg.isReg {
		return *m.getReg(arg)
	} else {
		return arg.value
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
