package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	row    int
	column int
}

func getneighbors(grid []string, coordinate Coordinate) []Coordinate {
	var neighbors []Coordinate
	var activecoordinate Coordinate
	//UL
	activecoordinate.row = coordinate.row - 1
	activecoordinate.column = coordinate.column - 1
	if activecoordinate.row >= 0 && activecoordinate.column >= 0 {
		neighbors = append(neighbors, activecoordinate)
	}

	//UC
	activecoordinate.row = coordinate.row - 1
	activecoordinate.column = coordinate.column
	if activecoordinate.row >= 0 {
		neighbors = append(neighbors, activecoordinate)
	}

	//UR
	activecoordinate.row = coordinate.row - 1
	activecoordinate.column = coordinate.column + 1
	if activecoordinate.row >= 0 && activecoordinate.column < len(grid[activecoordinate.row]) {
		neighbors = append(neighbors, activecoordinate)
	}

	//CL
	activecoordinate.row = coordinate.row
	activecoordinate.column = coordinate.column - 1
	if activecoordinate.column >= 0 {
		neighbors = append(neighbors, activecoordinate)
	}

	//CR
	activecoordinate.row = coordinate.row
	activecoordinate.column = coordinate.column + 1
	if activecoordinate.column < len(grid[activecoordinate.row]) {
		neighbors = append(neighbors, activecoordinate)
	}

	//LL
	activecoordinate.row = coordinate.row + 1
	activecoordinate.column = coordinate.column - 1
	if activecoordinate.row < len(grid) && activecoordinate.column >= 0 {
		neighbors = append(neighbors, activecoordinate)
	}

	//LC
	activecoordinate.row = coordinate.row + 1
	activecoordinate.column = coordinate.column
	if activecoordinate.row < len(grid) {
		neighbors = append(neighbors, activecoordinate)
	}

	//LR
	activecoordinate.row = coordinate.row + 1
	activecoordinate.column = coordinate.column + 1
	if activecoordinate.row < len(grid) && activecoordinate.column < len(grid[activecoordinate.row]) {
		neighbors = append(neighbors, activecoordinate)
	}
	return neighbors
}

func checkforroll(grid []string, coordinate Coordinate) bool {
	if grid[coordinate.row][coordinate.column] == '@' {
		return true
	} else {
		return false
	}
}

func checkaccessible(grid []string, coordinate Coordinate) bool {
	neighbors := getneighbors(grid, coordinate)
	neighborrolls := 0
	for _, neighbor := range neighbors {
		if checkforroll(grid, neighbor) {
			neighborrolls++
		}
	}
	if neighborrolls < 4 {
		return true
	} else {
		return false
	}
}

func removerolls(grid []string, coordinate Coordinate) []string {
	affectedrow := []rune(grid[coordinate.row])
	affectedrow[coordinate.column] = '.'
	grid[coordinate.row] = string(affectedrow)
	return grid
}

func partone() int {
	file, _ := os.Open("2025-04.txt")
	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	fmt.Println(grid)
	accessiblerolls := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if checkforroll(grid, Coordinate{row, column}) && checkaccessible(grid, Coordinate{row, column}) {
				accessiblerolls++
			}
		}
	}
	return accessiblerolls
}

func parttwo() int {
	file, _ := os.Open("2025-04.txt")
	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	fmt.Println(grid)
	removedrolls := 0
	anyremoved := false
	for next := true; next; next = anyremoved {
		anyremoved = false
		for row := 0; row < len(grid); row++ {
			for column := 0; column < len(grid[row]); column++ {
				activecoordinate := Coordinate{row, column}
				if checkforroll(grid, activecoordinate) && checkaccessible(grid, activecoordinate) {
					grid = removerolls(grid, activecoordinate)
					removedrolls++
					anyremoved = true
				}
			}
		}
	}
	return removedrolls
}

func main() {
	fmt.Println(partone())
	fmt.Println(parttwo())
}
