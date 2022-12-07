package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var letters = []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println("Read", s)
		charcount := 0
		for _, b := range []byte(s) {
			charcount++
			letters = append(letters, string(b))
			if len(letters) == 4 {
				if countUnique(letters) == 4 {
					fmt.Println(charcount)
					letters = []string{}
					break
				}
				letters = letters[1:]
			}
		}
	}
}

func countUnique(s []string) int {
	unique := 0
	letters := make(map[string]bool)

	for _, b := range s {
		if !letters[b] {
			letters[b] = true
			unique++
		}
	}
	return unique
}
