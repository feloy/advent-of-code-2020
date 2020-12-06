package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	found1 := 0
	found3 := 0
	found5 := 0
	found7 := 0
	found2 := 0
	x1 := 0
	x3 := 0
	x5 := 0
	x7 := 0
	x2 := 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		s1 := line[x1]
		if s1 == '#' {
			found1++
		}
		s3 := line[x3]
		if s3 == '#' {
			found3++
		}
		s5 := line[x5]
		if s5 == '#' {
			found5++
		}
		s7 := line[x7]
		if s7 == '#' {
			found7++
		}
		s2 := line[x2]
		if s2 == '#' && i%2 == 0 {
			found2++
		}
		x1 = (x1 + 1) % len(line)
		x3 = (x3 + 3) % len(line)
		x5 = (x5 + 5) % len(line)
		x7 = (x7 + 7) % len(line)
		if i%2 == 0 {
			x2 = (x2 + 1) % len(line)
		}
		i++
	}
	fmt.Printf("%d\n", found1*found3*found5*found7*found2)
}
