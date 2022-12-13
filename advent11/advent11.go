package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	items   []int
	op      string
	div     int
	iftrue  int
	iffalse int
	count   int
}

func main() {
	var m [8]monkey

	reMonkey, _ := regexp.Compile(`Monkey (\d+):`)
	reItems, _ := regexp.Compile(`Starting items: (.*)`)
	reOp, _ := regexp.Compile(`Operation: new = old (.+)`)
	reTest, _ := regexp.Compile(`Test: divisible by (\d+)`)
	reTrue, _ := regexp.Compile(`If true: throw to monkey (\d+)`)
	reFalse, _ := regexp.Compile(`If false: throw to monkey (\d+)`)

	var mnum int
	var items []int
	var op string
	var div int
	var iftrue int
	var iffalse int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if reMonkey.MatchString(s) {
			mnum, _ = strconv.Atoi(reMonkey.FindStringSubmatch(s)[1])
		} else if reItems.MatchString(s) {
			itemstr := reItems.FindStringSubmatch(s)[1]
			for _, i := range strings.Split(itemstr, ", ") {
				val, _ := strconv.Atoi(i)
				items = append(items, val)
			}
		} else if reOp.MatchString(s) {
			op = reOp.FindStringSubmatch(s)[1]
		} else if reTest.MatchString(s) {
			div, _ = strconv.Atoi(reTest.FindStringSubmatch(s)[1])
		} else if reTrue.MatchString(s) {
			iftrue, _ = strconv.Atoi(reTrue.FindStringSubmatch(s)[1])
		} else if reFalse.MatchString(s) {
			iffalse, _ = strconv.Atoi(reFalse.FindStringSubmatch(s)[1])
		} else {
			m[mnum] = monkey{
				items:   items,
				op:      op,
				div:     div,
				iftrue:  iftrue,
				iffalse: iffalse,
				count:   0,
			}
			items = []int{}
		}
	}

	for round := 0; round < 20; round++ {
		fmt.Println("Round ", round, " processing...")
		for mnum := 0; mnum < 8; mnum++ {
			if m[mnum].op != "" {
				ops := strings.Split(m[mnum].op, " ")
				for len(m[mnum].items) > 0 {
					worry := m[mnum].items[0]
					num := 0
					if ops[1] == "old" {
						num = worry
					} else {
						num, _ = strconv.Atoi(ops[1])
					}
					if ops[0] == "+" {
						worry += num
					} else if ops[0] == "*" {
						worry *= num
					}

					worry /= 3
					newM := 0
					if worry%m[mnum].div == 0 {
						newM = m[mnum].iftrue
					} else {
						newM = m[mnum].iffalse
					}

					m[newM].items = append(m[newM].items, worry)
					m[mnum].items = m[mnum].items[1:]
					m[mnum].count++
				}
			}
		}
	}

	fmt.Println("Processing complete...")

	var max1, max2 int
	for mnum := 0; mnum < 8; mnum++ {
		if m[mnum].count > max1 {
			if max1 > max2 {
				max2 = max1
			}
			max1 = m[mnum].count
		} else if m[mnum].count > max2 {
			max2 = m[mnum].count
		}
	}
	fmt.Println("monkey business:", max1*max2, " which is ", max1, " and ", max2)
}
