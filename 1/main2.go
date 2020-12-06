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
			for _, k := range lines {
				if i == j || i == k || j == k {
					continue
				}
				if i+j+k == 2020 {
					fmt.Printf("%d\n", i*j*k)
					return
				}
			}
		}
	}
}
