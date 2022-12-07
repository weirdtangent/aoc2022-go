package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dirsize := make(map[string]int)
	path := []string{}
	inls := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.Split(s, " ")
		if inls {
			if words[0] == "$" {
				inls = false
			} else {
				if words[0] != "dir" {
					filesize, _ := strconv.Atoi(words[0])
					dirname := ""
					for _, dir := range path {
						dirname += dir + "/"
						dirsize[dirname] += filesize
					}

				}
			}
		}

		if words[1] == "cd" {
			if words[2] != ".." {
				if words[2] == "/" {
					path = []string{}
					words[2] = "<root>"
				}
				path = append(path, words[2])
			} else {
				path = path[0 : len(path)-1]
			}
		} else if s == "$ ls" {
			inls = true
		}
	}

	fmt.Println("Finding dirs > 10000")
	total := 0
	for dir, _ := range dirsize {
		if dir != "<root>/" && dirsize[dir] <= 100000 {
			fmt.Println(dir, dirsize[dir])
			total += dirsize[dir]
		}
	}
	fmt.Println(total)
}
