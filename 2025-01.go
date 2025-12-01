package main

import (
	"bufio"
	"os"
	"strconv"
)

func FlooredModInt(a, n int) int {
	res := a % n
	if (res < 0 && n > 0) || (res > 0 && n < 0) {
		return res + n
	}
	return res
}

func Absint(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func partone() {
	file, _ := os.Open("2025-01.txt")
	scanner := bufio.NewScanner(file)
	selectednumber := 50
	zerocount := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := 1
		if line[0:1] == "L" {
			direction = -1
		}
		clicks, _ := strconv.Atoi(line[1:])
		selectednumber = FlooredModInt(selectednumber+(direction*clicks), 100)
		if selectednumber == 0 {
			zerocount++
		}
	}
	print("Part one is: ", zerocount, "\r\n")
}

func parttwo() {
	file, _ := os.Open("2025-01.txt")
	scanner := bufio.NewScanner(file)
	selectednumber := 50
	zerocount := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := 1
		if line[0:1] == "L" {
			direction = -1
		}
		clicks, _ := strconv.Atoi(line[1:])
		selectednumber = selectednumber + (direction * clicks)
		for selectednumber > 100 {
			selectednumber = selectednumber - 100
			zerocount += 1
		}
		if selectednumber < 0 && selectednumber == clicks*direction {
			selectednumber += 100
		}
		for selectednumber < 0 {
			selectednumber = selectednumber + 100
			zerocount += 1
		}
		if selectednumber == 100 {
			selectednumber = selectednumber - 100
		}
		if selectednumber == 0 {
			zerocount += 1
		}
	}
	print("Part two is: ", zerocount)

}

func main() {
	partone()
	parttwo()
}
