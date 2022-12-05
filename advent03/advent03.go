package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var c1 map[int]int
var c2 map[int]int
var b map[int]int
var total int
var btotal int

func main() {
	total := 0
	btotal := 0
	b := make(map[int]int)
	lines := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines++
		c1 := make(map[int]int)
		c2 := make(map[int]int)
		str := scanner.Text()
		letters := strings.Split(str, "")
		count := len(str) / 2
		for pos, ltr := range letters {
			chr := []rune(ltr)
			ascii := int(chr[0])

			if pos < count {
				c1[ascii]++
			} else {
				c2[ascii]++
			}
		}

		// part 1
		for ascii, _ := range c1 {
			if c1[ascii] > 0 && c2[ascii] > 0 {
				if ascii > 96 {
					total += ascii - 96
				} else {
					total += ascii - 38
				}
			}
		}

		// part 2
		for ascii, _ := range c1 {
			if c1[ascii] > 0 {
				b[ascii]++
			}
		}
		for ascii, _ := range c2 {
			if c2[ascii] > 0 {
				b[ascii]++
			}
		}
		if lines%3 == 0 {
			for ascii, _ := range b {
				if b[ascii] == 3 {
					fmt.Println("For badge", string(ascii), "we found count", b[ascii])
					if ascii > 96 {
						btotal += ascii - 96
					} else {
						btotal += ascii - 38
					}
				}
				b[ascii] = 0
			}

		}

	}

	fmt.Println("Total ascii for duplicate items is:", total)
	fmt.Println("Badge total is:", btotal)

}
