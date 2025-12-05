package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"

	"os"
)

func findmaxjoltage(batteries string) int {
	maxjoltage := '0'
	maxjoltageones := '0'
	var maxindex int
	for index, battery := range batteries[0 : len(batteries)-1] {
		if battery > maxjoltage {
			maxjoltage = battery
			maxindex = index
		}
	}
	for _, battery := range batteries[maxindex+1:] {
		if battery > maxjoltageones {
			maxjoltageones = battery
		}
	}
	tens := int(maxjoltage - '0')
	ones := int(maxjoltageones - '0')
	return tens*10 + ones
}

func findmegaultrajoltage(batteries string) int {
	var megaultrajoltage []rune
	marker := 0
	for i := 1; i <= 12; i++ {
		insidemarker := 0
		joltage := '0'
		for index, battery := range batteries[marker:(len(batteries) - (12 - i))] {
			if battery > joltage {
				joltage = battery
				insidemarker = index + 1
			}
		}
		megaultrajoltage = append(megaultrajoltage, joltage)
		marker += insidemarker
	}
	megaultrajoltageint := 0
	slices.Reverse(megaultrajoltage)
	for index, battery := range megaultrajoltage {
		megaultrajoltageint += int(battery-'0') * int(math.Pow(10, float64(index)))
	}
	return megaultrajoltageint
}

func partone() {
	file, _ := os.Open("2025-03.txt")
	scanner := bufio.NewScanner(file)
	joltage := 0
	for scanner.Scan() {
		joltage += findmaxjoltage(scanner.Text())
	}
	fmt.Println(joltage)
}

func parttwo() {
	file, _ := os.Open("2025-03.txt")
	scanner := bufio.NewScanner(file)
	joltage := 0
	for scanner.Scan() {
		joltage += findmegaultrajoltage(scanner.Text())
	}
	fmt.Println(joltage)
}

func main() {
	partone()
	parttwo()
}
