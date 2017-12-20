package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	particles := make([]particle, 0)

	particleRule := regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)

	for scanner.Scan() {
		line := scanner.Text()
		results := particleRule.FindStringSubmatch(line)
		pos := vector3{ toInt(results[1]), toInt(results[2]), toInt(results[3]) }
		vel := vector3{ toInt(results[4]), toInt(results[5]), toInt(results[6]) }
		acc := vector3{ toInt(results[7]), toInt(results[8]), toInt(results[9]) }
		particles = append(particles, particle{ pos, vel, acc })
	}

	fmt.Println("--- Part One ---")
	minIndex, minAcc := 0, particles[0].acceleration.length()
	for i := 1; i < len(particles); i++ {
		acc := particles[i].acceleration.length()
		if acc < minAcc {
			minIndex, minAcc = i, acc
		}
	}
	fmt.Println(minIndex)

	fmt.Println("--- Part Two ---")
	unchangedIterations := 0
	for unchangedIterations < 1000 {
		for i := 0; i < len(particles); i++ {
			particles[i].velocity.add(particles[i].acceleration)
			particles[i].position.add(particles[i].velocity)
		}
		positions := make(map[vector3][]int)
		for i, p := range particles {
			positions[p.position] = append(positions[p.position], i)
		}
		toRemove := make([]int, 0)
		for _, indices := range positions {
			if len(indices) >= 2 {
				toRemove = append(toRemove, indices...)
			}
		}
		sort.Ints(toRemove)
		for i := len(toRemove) - 1; i >= 0; i-- {
			current := toRemove[i]
			particles[current] = particles[len(particles)-1]
			particles = particles[:len(particles)-1]
		}
		if len(toRemove) > 0 {
			unchangedIterations = 0
		} else {
			unchangedIterations++
		}
	}
	fmt.Println(len(particles))
}

type vector3 struct {
	x, y, z int
}

func (v *vector3) add(o vector3) {
	v.x += o.x
	v.y += o.y
	v.z += o.z
}

func (v *vector3) length() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

type particle struct {
	position, velocity, acceleration vector3
}

func abs(v int) int {
	if v < 0 { return -v }
	return v
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
