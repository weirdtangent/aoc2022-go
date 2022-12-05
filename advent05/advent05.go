package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

		for x := 0; x < amount; x++ {
			move(from, to)
		}
	}

	str := ""
	for cnum, _ := range crate {
		if len(crate[cnum]) > 0 {
			str += string(crate[cnum][0])
		}
	}
	fmt.Println(str)
}

func move(from, to int) {
	moving := crate[from][0]
	crate[from] = crate[from][1:]
	crate[to] = append([]byte{moving}, crate[to]...)
}
