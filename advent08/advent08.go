package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mapsize = 99
var treemap = [99][99]int{}

func main() {
	mapx := 0
	mapy := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		mapx = 0
		for _, height := range []byte(s) {
			if string(height) != "" {
				treemap[mapy][mapx], _ = strconv.Atoi(string(height))
				mapx++
			}
		}
		mapy++
	}

	part1()
	part2()
}

func part1() {
	visible := 0
	for y := 0; y < mapsize; y++ {
		for x := 0; x < mapsize; x++ {
			if isVisible(y, x, -1, 0) || isVisible(y, x, 1, 0) || isVisible(y, x, 0, -1) || isVisible(y, x, 0, 1) {
				visible++
			} else {
			}
		}
	}
	fmt.Println(visible, "trees are visible from at least 1 direction")
}

func part2() {
	maxScore := 0
	for y := 0; y < mapsize; y++ {
		visibleScore := 0
		for x := 0; x < mapsize; x++ {
			visibleScore = countVisible(y, x, -1, 0) * countVisible(y, x, 1, 0) * countVisible(y, x, 0, -1) * countVisible(y, x, 0, 1)
			if visibleScore > maxScore {
				maxScore = visibleScore
			}
		}
	}
	fmt.Println("max visible tree score is", maxScore)
}

func isVisible(y, x, dy, dx int) bool {
	myHeight := treemap[y][x]
	y += dy
	x += dx
	for y >= 0 && y < mapsize && x >= 0 && x < mapsize {
		if treemap[y][x] >= myHeight {
			return false
		}
		y += dy
		x += dx
	}
	return true
}

func countVisible(y, x, dy, dx int) int {
	count := 0
	myHeight := treemap[y][x]
	y += dy
	x += dx
	for y >= 0 && y < mapsize && x >= 0 && x < mapsize {
		if treemap[y][x] >= myHeight {
			count++
			return count
		} else {
			count++
			y += dy
			x += dx
		}
	}
	return count
}
