package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//    [D]
//[N] [C]
//[Z] [M] [P]
// 1   2   3
//
//move 1 from 2 to 1
//move 3 from 1 to 3
//move 2 from 2 to 1
//move 1 from 1 to 2

var crate = [10][]byte{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		chars := []byte(s)
		if chars[1] == '1' {
			break
		}
		cnum := 1
		for pos := 1; pos < len(s); pos += 4 {
			if chars[pos] != ' ' {
				crate[cnum] = append(crate[cnum][:], chars[pos])
			}
			cnum++
		}
	}

	scanner.Scan()

	for scanner.Scan() {
		s := scanner.Text()
		cmd := strings.Split(s, " ")
		amount, _ := strconv.Atoi(cmd[1])
		from, _ := strconv.Atoi(cmd[3])
		to, _ := strconv.Atoi(cmd[5])

		move(amount, from, to)
	}

	str := ""
	for cnum, _ := range crate {
		if len(crate[cnum]) > 0 {
			str += string(crate[cnum][0])
		}
	}
	fmt.Println(str)
}

func move(amount, from, to int) {
	moving := crate[from][0:amount]
	crate[from] = crate[from][amount:]
	for x := len(moving) - 1; x >= 0; x-- {
		crate[to] = append([]byte{moving[x]}, crate[to]...)
	}
}
