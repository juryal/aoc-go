package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Freshrange struct {
	start int
	end   int
}

type CoolerFreshrange struct {
	Freshrange
	count int
}

func (f Freshrange) isinrange(x int) bool {
	return x >= f.start && x <= f.end
}

func partone() int {
	file, _ := os.Open("2025-05.txt")
	scanner := bufio.NewScanner(file)
	var fresh []Freshrange
	freshcount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			endpoints := strings.Split(line, "-")
			start, _ := strconv.Atoi(endpoints[0])
			end, _ := strconv.Atoi(endpoints[1])
			fresh = append(fresh, Freshrange{start, end})
		} else {
			food, _ := strconv.Atoi(line)
			for _, freshrange := range fresh {
				if freshrange.isinrange(food) {
					freshcount++
					break
				}
			}
		}
	}
	return freshcount
}

func parttwo() int {
	file, _ := os.Open("2025-05.txt")
	scanner := bufio.NewScanner(file)
	var fresh []CoolerFreshrange
	freshcount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			endpoints := strings.Split(line, "-")
			start, _ := strconv.Atoi(endpoints[0])
			end, _ := strconv.Atoi(endpoints[1])
			fresh = append(fresh, CoolerFreshrange{Freshrange{start, end}, end - start + 1})
		}
	}
	slices.SortFunc(fresh, func(a, b CoolerFreshrange) int { return a.count - b.count })
	for i := 0; i < len(fresh)-1; i++ {
		for j := i + 1; j < len(fresh); j++ {
			//Start
			if fresh[j].isinrange(fresh[i].start) {
				fresh[i].start = fresh[j].end + 1
				fresh[i].count = max(0, fresh[i].end-fresh[i].start+1)
			}
			//End
			if fresh[j].isinrange(fresh[i].end) {
				fresh[i].end = fresh[j].start - 1
				fresh[i].count = max(0, fresh[i].end-fresh[i].start+1)
			}
			if fresh[i].count == 0 {
				break
			}

		}
	}
	for _, x := range fresh {
		freshcount += x.count
	}
	return freshcount
}

func main() {
	fmt.Println(partone())
	fmt.Println(parttwo())
}
