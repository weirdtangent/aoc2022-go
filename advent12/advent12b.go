package main

import (
	"bufio"
	"fmt"
	"os"
)

type next struct {
	steps, x, y int
}
type point struct {
	x, y int
}

var queue []next
var seen [41][77]bool

func main() {
	height := [41][77]int{}
	try := []point{}

	scanner := bufio.NewScanner(os.Stdin)
	x, y := 0, 0
	endx, endy := 0, 0
	for scanner.Scan() {
		s := scanner.Text()
		for _, alt := range []byte(s) {
			if alt == 'S' {
				alt = 97
			} else if alt == 'E' {
				endx, endy = x, y
				alt = 97 + 26
			}
			height[y][x] = int(alt) - 97
			if alt == 'a' {
				try = append(try, point{x: x, y: y})
			}
			x++
		}
		x = 0
		y++
	}

	min := 0
	for _, t := range try {
		s := shortest(t.x, t.y, endx, endy, height)
		if s > 0 && (min == 0 || s < min) {
			min = s
		}
	}
	fmt.Println("only ", min, "steps")
}

func shortest(startx, starty, endx, endy int, height [41][77]int) int {
	seen := [41][77]bool{}
	queue := []next{}
	queue = append(queue, next{steps: 0, x: startx, y: starty})

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q.x == endx && q.y == endy {
			break
		}
		if seen[q.y][q.x] {
			continue
		}
		seen[q.y][q.x] = true

		if q.x-1 >= 0 && q.x-1 < 77 && q.y >= 0 && q.y < 41 && height[q.y][q.x-1]-height[q.y][q.x] < 2 {
			queue = append(queue, next{steps: q.steps + 1, x: q.x - 1, y: q.y})
		}
		if q.x+1 >= 0 && q.x+1 < 77 && q.y >= 0 && q.y < 41 && height[q.y][q.x+1]-height[q.y][q.x] < 2 {
			queue = append(queue, next{steps: q.steps + 1, x: q.x + 1, y: q.y})
		}
		if q.x >= 0 && q.x < 77 && q.y-1 >= 0 && q.y-1 < 41 && height[q.y-1][q.x]-height[q.y][q.x] < 2 {
			queue = append(queue, next{steps: q.steps + 1, x: q.x, y: q.y - 1})
		}
		if q.x >= 0 && q.x < 77 && q.y+1 >= 0 && q.y+1 < 41 && height[q.y+1][q.x]-height[q.y][q.x] < 2 {
			queue = append(queue, next{steps: q.steps + 1, x: q.x, y: q.y + 1})
		}
	}

	min := 0
	for _, q := range queue {
		if min == 0 || q.steps < min {
			min = q.steps
		}
	}
	return min
}
