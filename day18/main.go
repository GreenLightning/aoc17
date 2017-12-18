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
		m := soundMachine{}
		m.run(instructions)
		if m.hasRecovered {
			fmt.Println(m.recovered)
		}
	}

	{
		fmt.Println("--- Part Two ---")
		m := makeDuetMachine()
		m.run(instructions)
		fmt.Println(m.second.sendCounter)
	}
}

const (
	instrKindSnd = iota
	instrKindSet
	instrKindAdd
	instrKindMul
	instrKindMod
	instrKindRcv
	instrKindJgz
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
	if strings.HasPrefix(line, "snd ") {
		line = line[4:]
		x, _ := parseArg(line)
		return instruction{ instrKindSnd, x, argument{} }
	} else if strings.HasPrefix(line, "set ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindSet, x, y }
	} else if strings.HasPrefix(line, "add ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindAdd, x, y }
	} else if strings.HasPrefix(line, "mul ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindMul, x, y }
	} else if strings.HasPrefix(line, "mod ") {
		line = line[4:]
		x, line := parseRegArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindMod, x, y }
	} else if strings.HasPrefix(line, "rcv ") {
		line = line[4:]
		x, _ := parseRegArg(line)
		return instruction{ instrKindRcv, x, argument{} }
	} else if strings.HasPrefix(line, "jgz ") {
		line = line[4:]
		x, line := parseArg(line)
		y, _ := parseArg(line)
		return instruction{ instrKindJgz, x, y }
	} else {
		panic(fmt.Sprint("cannot parse instruction: ", line))
	}
}

type soundMachine struct {
	registers [26]int
	sent int
	recovered int
	hasRecovered bool
}

func (m *soundMachine) run(instructions []instruction) {
	ip := 0
	for !m.hasRecovered && ip >= 0 && ip < len(instructions) {
		i := instructions[ip]
		ip += m.execute(i)
	}
}

func (m *soundMachine) execute(i instruction) int {
	switch i.kind {
		case instrKindSnd:
			m.sent = m.getVal(i.x)
			return 1
		case instrKindSet:
			reg := m.getReg(i.x)
			*reg = m.getVal(i.y)
			return 1
		case instrKindAdd:
			reg := m.getReg(i.x)
			*reg += m.getVal(i.y)
			return 1
		case instrKindMul:
			reg := m.getReg(i.x)
			*reg *= m.getVal(i.y)
			return 1
		case instrKindMod:
			reg := m.getReg(i.x)
			*reg %= m.getVal(i.y)
			return 1
		case instrKindRcv:
			val := m.getVal(i.x)
			if val != 0 && !m.hasRecovered {
				m.recovered, m.hasRecovered = m.sent, true
			}
			return 1
		case instrKindJgz:
			result := m.getVal(i.x)
			if result > 0 {
				return m.getVal(i.y)
			} else {
				return 1
			}
		default: panic(fmt.Sprintf("unknown instruction kind4 %d", i.kind))
	}
}

func (m *soundMachine) getReg(arg argument) *int {
	if arg.isReg {
		return &m.registers[arg.value]
	} else {
		return nil
	}
}

func (m *soundMachine) getVal(arg argument) int {
	if arg.isReg {
		return *m.getReg(arg)
	} else {
		return arg.value
	}
}

type duetMachine struct {
	first, second parallelMachine
	firstToSecond []int
	secondToFirst []int
}

type parallelMachine struct {
	registers [26]int
	ip int
	blocked bool
	in, out *[]int
	sendCounter int
}

func makeDuetMachine() duetMachine {
	m := duetMachine{}
	m.first.in   = &m.secondToFirst
	m.first.out  = &m.firstToSecond
	m.second.in  = &m.firstToSecond
	m.second.out = &m.secondToFirst
	progReg := int('p' - 'a')
	m.first.registers[progReg] = 0
	m.second.registers[progReg] = 1
	return m
}

func (m *duetMachine) run(instructions []instruction) {
	for (!m.first.blocked || !m.second.blocked || len(m.firstToSecond) > 0 || len(m.secondToFirst) > 0) &&
	    ((m.first.ip >= 0 && m.first.ip < len(instructions)) || m.second.ip >= 0 && m.second.ip < len(instructions)) {
		if m.first.ip >= 0 && m.first.ip < len(instructions) {
			i := instructions[m.first.ip]
			offset, blocked := m.first.execute(i)
			m.first.ip += offset
			m.first.blocked = blocked
		}
		if m.second.ip >= 0 && m.second.ip < len(instructions) {
			i := instructions[m.second.ip]
			offset, blocked := m.second.execute(i)
			m.second.ip += offset
			m.second.blocked = blocked
		}
	}
}

func (m *parallelMachine) execute(i instruction) (int, bool) {
	switch i.kind {
		case instrKindSnd:
			value := m.getVal(i.x)
			*m.out = append(*m.out, value)
			m.sendCounter++
			return 1, false
		case instrKindSet:
			reg := m.getReg(i.x)
			*reg = m.getVal(i.y)
			return 1, false
		case instrKindAdd:
			reg := m.getReg(i.x)
			*reg += m.getVal(i.y)
			return 1, false
		case instrKindMul:
			reg := m.getReg(i.x)
			*reg *= m.getVal(i.y)
			return 1, false
		case instrKindMod:
			reg := m.getReg(i.x)
			*reg %= m.getVal(i.y)
			return 1, false
		case instrKindRcv:
			if len(*m.in) == 0 {
				return 0, true
			} else {
				value := (*m.in)[0]
				*m.in = (*m.in)[1:]
				reg := m.getReg(i.x)
				*reg = value
				return 1, false
			}
		case instrKindJgz:
			result := m.getVal(i.x)
			if result > 0 {
				return m.getVal(i.y), false
			} else {
				return 1, false
			}
		default: panic(fmt.Sprintf("unknown instruction kind4 %d", i.kind))
	}
}

func (m *parallelMachine) getReg(arg argument) *int {
	if arg.isReg {
		return &m.registers[arg.value]
	} else {
		return nil
	}
}

func (m *parallelMachine) getVal(arg argument) int {
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
