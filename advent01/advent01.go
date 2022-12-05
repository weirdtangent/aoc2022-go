package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var total map[int]int

func main() {
	total = make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	elf := 1
	for scanner.Scan() {
		s := scanner.Text()
		num, err := strconv.Atoi(s)
		if err == nil {
			total[elf] += num
		} else {
			elf++
		}
	}

	part1()
	part2()
}

func part1() {
	max_cals := 0
	max_elf := 0
	for elf, cals := range total {
		if cals > max_cals {
			max_elf = elf
			max_cals = cals
		}
	}
	fmt.Println("Elf", max_elf, "is carrying the most cals:", max_cals)
}

func part2() {
	best3 := 0
	count := 0
	max_cals := 0
	max_elf := 0
	for count < 3 {
		for elf, cals := range total {
			if cals > max_cals {
				max_elf = elf
				max_cals = cals
			}
		}
		best3 += max_cals
		total[max_elf] = 0
		max_elf = 0
		max_cals = 0
		count++
	}
	fmt.Println("Top 3 Elfs are carrying", best3)
}
