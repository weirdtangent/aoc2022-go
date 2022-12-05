package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1count := 0
	part2count := 0
	limits := [2][2]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		for rangenum, rangestr := range strings.Split(s, ",") {
			for limitnum, limitstr := range strings.Split(rangestr, "-") {
				limits[rangenum][limitnum], _ = strconv.Atoi(limitstr)
			}
		}

		if (limits[0][0] <= limits[1][0] && limits[0][1] >= limits[1][1]) || (limits[1][0] <= limits[0][0] && limits[1][1] >= limits[0][1]) {
			part1count++
		}
		if (limits[0][0] <= limits[1][0] && limits[0][1] >= limits[1][0]) || (limits[1][0] <= limits[0][0] && limits[1][1] >= limits[0][0]) {
			part2count++
		}
	}

	fmt.Println(part1count, "ranges overlap entirely")
	fmt.Println(part2count, "ranges overlap at all")
}
