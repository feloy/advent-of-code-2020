package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	good := 0

	for scanner.Scan() {
		line := scanner.Text()
		var min, max int
		var letter byte
		var password string
		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &letter, &password)
		if isValid(min, max, letter, password) {
			good++
		}
	}
	fmt.Printf("%d\n", good)

}

func isValid(min int, max int, letter byte, password string) bool {
	n := strings.Count(password, string(letter))
	return n >= min && n <= max
}
