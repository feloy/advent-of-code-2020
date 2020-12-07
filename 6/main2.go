package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mapp := map[rune]int{}
	total := 0

	groupsize := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)
		if line == "" {
			fmt.Printf("group size: %d\n", groupsize)
			fmt.Printf("%+v\n", mapp)
			count := 0
			for _, m := range mapp {
				if m == groupsize {
					count++
				}
			}
			fmt.Printf("%d\n\n", count)
			total += count
			groupsize = 0
			mapp = map[rune]int{}
			continue
		}
		for _, l := range line {
			mapp[l]++
		}
		groupsize++
	}

	count := 0
	for _, m := range mapp {
		if m == groupsize {
			count++
		}
	}

	fmt.Printf("=> %d\n", total+count)
}
