package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	good := 0

	for scanner.Scan() {
		line := scanner.Text()
		var pos1, pos2 int
		var letter byte
		var password string
		fmt.Sscanf(line, "%d-%d %c: %s", &pos1, &pos2, &letter, &password)
		if isValid(pos1, pos2, letter, password) {
			good++
		}
	}
	fmt.Printf("%d\n", good)

}

func isValid(pos1 int, pos2 int, letter byte, password string) bool {
	atPos1 := password[pos1-1:pos1] == string(letter)
	atPos2 := password[pos2-1:pos2] == string(letter)

	if atPos1 && atPos2 {
		return false
	}
	return atPos1 || atPos2
}
