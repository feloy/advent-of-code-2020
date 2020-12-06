package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, s)
	}

	for _, i := range lines {
		for _, j := range lines {
			if i == j {
				continue
			}
			if i+j == 2020 {
				fmt.Printf("%d\n", i*j)
				return
			}
		}
	}
}
