package main

import (
	"bufio"
	"fmt"
	"os"
)

var leastMoves int

func main() {
	height := [5][8]int{}

	scanner := bufio.NewScanner(os.Stdin)
	x, y := 0, 0
	startx, starty := 0, 0
	endx, endy := 0, 0
	for scanner.Scan() {
		s := scanner.Text()
		for _, alt := range []byte(s) {
			if alt == 'S' {
				startx, starty = x, y
				alt = 97
			} else if alt == 'E' {
				endx, endy = x, y
				alt = 97 + 26
			}
			height[y][x] = int(alt) - 97
			x++
		}
		x = 0
		y++
	}

	getCloser(0, startx, starty, endx, endy, height)
	fmt.Println("Solved in least-moves of ", leastMoves)

}

func getCloser(moves, startx, starty, endx, endy int, height [5][8]int) bool {
	alt := height[starty][startx]

	//fmt.Println("height here at ", starty, ",", startx, " is ", height[starty][startx])
	if startx == endx && starty == endy {
		if leastMoves == 0 || moves < leastMoves {
			leastMoves = moves
			fmt.Println("SOLVED! in ", moves, " moves this path. Least is ", leastMoves)
		}
		return true
	}

	if starty > 0 && (height[starty-1][startx]-alt) <= 1 {
		//fmt.Println("move: ", moves+1, " - alt here at ", starty, ",", startx, " is ", alt, " and we're going UP")
		height[starty][startx] = 100
		if !getCloser(moves+1, startx, starty-1, endx, endy, height) {
			height[starty-1][startx] = 199
		}
	}

	if starty < 4 && (height[starty+1][startx]-alt) <= 1 {
		//fmt.Println("move: ", moves+1, " - alt here at ", starty, ",", startx, " is ", alt, " and we're going DOWN")
		height[starty][startx] = 101
		if !getCloser(moves+1, startx, starty+1, endx, endy, height) {
			height[starty+1][startx] = 199
		}
	}

	if startx > 0 && (height[starty][startx-1]-alt) <= 1 {
		//fmt.Println("move: ", moves+1, " - alt here at ", starty, ",", startx, " is ", alt, " and we're going LEFT")
		height[starty][startx] = 102
		if !getCloser(moves+1, startx-1, starty, endx, endy, height) {
			height[starty][startx-1] = 199
		}
	}

	if startx < 7 && (height[starty][startx+1]-alt) <= 1 {
		//fmt.Println("move: ", moves+1, " - alt here at ", starty, ",", startx, " is ", alt, " and we're going RIGHT")
		height[starty][startx] = 103
		if !getCloser(moves+1, startx+1, starty, endx, endy, height) {
			height[starty][startx+1] = 199
		}
	}

	//fmt.Println("last move left us stuck, backtrack")
	return false
}
