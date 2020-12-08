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

	term, _ := terminates(lines)
	fmt.Printf("%v\n", term)
	p := 0
	for {
		modified := make([]Line, len(lines))
		copy(modified, lines)
		switch modified[p].Operation {
		case "jmp":
			fmt.Printf("modify at %d\n", p)
			modified[p].Operation = "nop"
			term, acc := terminates(modified)
			if term {
				fmt.Printf("%d\n", acc)
				return
			}
		case "nop":
			fmt.Printf("modify at %d\n", p)
			modified[p].Operation = "jmp"
			term, acc := terminates(modified)
			if term {
				fmt.Printf("%d\n", acc)
				return
			}
		}
		p++
	}
}

func scanLine(s string) Line {
	var l Line
	fmt.Sscanf(s, "%s %d", &l.Operation, &l.Value)
	return l
}

func terminates(lines []Line) (bool, int) {
	//	fmt.Printf("%#v\n", lines)
	acc := 0
	p := 0

	used := map[int]bool{}

	for {
		if used[p] {
			return false, 0
		}
		if p == len(lines) {
			return true, acc
		}
		op := lines[p]
		//		fmt.Printf("%d\n", p)
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
