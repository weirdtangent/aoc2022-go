package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var hx, hy = 0, 0
var tx, ty = 0, 0
var tailsx = [10]int{}
var tailsy = [10]int{}
var trace map[string]bool

func main() {
	trace = make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		params := strings.Split(s, " ")
		dir := params[0]
		amt, _ := strconv.Atoi(params[1])
		for amt > 0 {
			for tailcnt := 1; tailcnt < 10; tailcnt++ {
				hx, hy = tailsx[tailcnt-1], tailsy[tailcnt-1]
				tx, ty = tailsx[tailcnt], tailsy[tailcnt]
				moveHead(dir, tailcnt)
				tailsx[tailcnt-1], tailsy[tailcnt-1] = hx, hy
				tailsx[tailcnt], tailsy[tailcnt] = tx, ty
			}
			amt--
		}

	}

	count := 0
	for range trace {
		count++
	}
	fmt.Println("TAILS visited", count, "unique positions")
}

func moveHead(dir string, tailcnt int) {
	if tailcnt == 1 {
		switch dir {
		case "U":
			hy--
		case "D":
			hy++
		case "L":
			hx--
		case "R":
			hx++
		}
	}

	if hx != tx && hy != ty && (hx-tx < -1 || hx-tx > 1 || hy-ty < -1 || hy-ty > 1) {
		if hy > ty {
			ty++
		} else {
			ty--
		}
		if hx > tx {
			tx++
		} else {
			tx--
		}
	} else if hx-tx < -1 || hx-tx > 1 || hy-ty < -1 || hy-ty > 1 {
		// if on same row
		if hx == tx {
			if hy > ty {
				ty++
			} else {
				ty--
			}
		}
		// if on same col
		if hy == ty {
			if hx > tx {
				tx++
			} else {
				tx--
			}
		}
	}

	if tailcnt == 9 {
		pos := strconv.Itoa(tx) + "," + strconv.Itoa(ty)
		trace[pos] = true
	}
}
