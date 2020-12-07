package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mapp := map[rune]bool{}
	total := 0

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)
		if line == "" {
			fmt.Printf("%+v\n", mapp)
			total += len(mapp)
			mapp = map[rune]bool{}
			fmt.Printf("count: %d\n", count)
			fmt.Printf("\n")
		}
		for _, l := range line {
			mapp[l] = true
		}
	}

	fmt.Printf("=> %d\n", total+len(mapp))
}
