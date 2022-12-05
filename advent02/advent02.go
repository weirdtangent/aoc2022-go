package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var games []string

func main() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		s := scanner.Text()
		if len(s) != 0 {
			count++
			games = append(games, s)
		} else {
			break
		}
	}
	fmt.Println("Loaded ", count, " games")

	part1()
	part2()
}

func part1() {
	total := 0
	count := 0
	for _, plays := range games {
		choices := strings.Split(plays, " ")
		if len(choices) == 2 {
			count++
			total += playRPS(choices[0], choices[1])
		} else {
			fmt.Println("invalid game ", count, " with ", len(choices), " choices")
		}
	}
	fmt.Println("Total score: ", total, " after ", count, " games")
}

func part2() {
	total := 0
	count := 0
	for _, plays := range games {
		choices := strings.Split(plays, " ")
		if len(choices) == 2 {
			count++
			me := ""
			if choices[0] == "A" {
				if choices[1] == "X" {
					me = "Z"
				} else if choices[1] == "Y" {
					me = "X"
				} else {
					me = "Y"
				}
			}
			if choices[0] == "B" {
				if choices[1] == "X" {
					me = "X"
				} else if choices[1] == "Y" {
					me = "Y"
				} else {
					me = "Z"
				}
			}
			if choices[0] == "C" {
				if choices[1] == "X" {
					me = "Y"
				} else if choices[1] == "Y" {
					me = "Z"
				} else {
					me = "X"
				}
			}
			total += playRPS(choices[0], me)
		} else {
			fmt.Println("invalid game ", count, " with ", len(choices), " choices")
		}
	}
	fmt.Println("Total score: ", total, " after ", count, " games")
}

func playRPS(op, me string) int {
	score := 0
	if me == "X" {
		score += 1
	} else if me == "Y" {
		score += 2
	} else if me == "Z" {
		score += 3
	}
	if (op == "A" && me == "X") || (op == "B" && me == "Y") || (op == "C" && me == "Z") {
		score += 3
	} else if (op == "A" && me == "Y") || (op == "B" && me == "Z") || (op == "C" && me == "X") {
		score += 6
	}
	return score
}
