package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	found := 0
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := line[x]
		if s == '#' {
			found++
		}
		x += 3
		x = x % len(line)
	}
	fmt.Printf("%d\n", found)
}
