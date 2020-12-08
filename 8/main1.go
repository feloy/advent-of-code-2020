package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
	Operation string
	Value     int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []Line{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, scanLine(line))
	}
	//	fmt.Printf("%#v\n", lines)
	acc := 0
	p := 0

	used := map[int]bool{}

	for {
		if used[p] {
			fmt.Printf("%d\n", acc)
			break
		}
		op := lines[p]
		used[p] = true
		switch op.Operation {
		case "acc":
			acc += op.Value
			p++
		case "jmp":
			p += op.Value
		case "nop":
			p++
		}
	}
}

func scanLine(s string) Line {
	var l Line
	fmt.Sscanf(s, "%s %d", &l.Operation, &l.Value)
	return l
}
