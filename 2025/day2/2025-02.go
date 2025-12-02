package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getendpoints(x string) (int, int) {
	first, last, _ := strings.Cut(x, "-")
	firstint, _ := strconv.Atoi(first)
	lastint, _ := strconv.Atoi(last)
	return firstint, lastint
}

func testid(x int) bool {
	numerals := strconv.Itoa(x)
	length := len(numerals)
	if length%2 == 0 {
		midpoint := length / 2
		return numerals[0:midpoint] == numerals[midpoint:length]
	}
	return false
}

func testidpartdeux(x int) bool {
	numerals := strconv.Itoa(x)
	for i := 1; i <= len(numerals)/2; i++ {
		if len(numerals)%i == 0 {
			if strings.Repeat(numerals[0:i], len(numerals)/i) == numerals {
				return true
			}
		}
	}
	return false
}

func partone() {
	file, _ := os.ReadFile("2025-02.txt")
	input := string(file)
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")
	idsum := 0
	for _, idrange := range ranges {
		first, last := getendpoints(idrange)
		for i := first; i <= last; i++ {
			if testid(i) {
				idsum += i
			}
		}
	}
	fmt.Println(idsum)
}

func parttwo() {
	file, _ := os.ReadFile("2025-02.txt")
	input := string(file)
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")
	idsum := 0
	for _, idrange := range ranges {
		first, last := getendpoints(idrange)
		for i := first; i <= last; i++ {
			if testidpartdeux(i) {
				idsum += i
			}
		}
	}
	fmt.Println(idsum)
}

func main() {
	partone()
	parttwo()
}
