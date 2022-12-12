package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xreg struct {
	before, during, after int
}

func main() {
	cycle := [1000]xreg{}
	cyclecnt := 0

	cycle[0] = xreg{before: 1, during: 1, after: 1}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cyclecnt++
		cycle[cyclecnt] = xreg{before: cycle[cyclecnt-1].after, during: cycle[cyclecnt-1].after, after: cycle[cyclecnt-1].after}
		s := scanner.Text()
		parse := strings.Split(s, " ")

		switch parse[0] {
		case "noop":
			pixel(cyclecnt, cycle[cyclecnt].during)

		case "addx":
			xchange, _ := strconv.Atoi(parse[1])
			pixel(cyclecnt, cycle[cyclecnt].during)

			cyclecnt++
			cycle[cyclecnt] = cycle[cyclecnt-1]
			cycle[cyclecnt].after += xchange
			pixel(cyclecnt, cycle[cyclecnt].during)
		}
	}

	part1 := (cycle[20].during * 20) + (cycle[60].during * 60) + (cycle[100].during * 100) + (cycle[140].during * 140) + (cycle[180].during * 180) + (cycle[220].during * 220)
	fmt.Println("\npart 1 sum: ", part1)
}

func pixel(cnt, val int) {
	for cnt > 40 {
		cnt -= 40
	}
	if cnt >= val && cnt <= val+2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if cnt%40 == 0 {
		fmt.Println()
	}
}
