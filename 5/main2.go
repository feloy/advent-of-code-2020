package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	found := map[int64]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		bin := strings.Replace(line, "B", "1", 10)
		bin = strings.Replace(bin, "F", "0", 10)
		bin = strings.Replace(bin, "R", "1", 10)
		bin = strings.Replace(bin, "L", "0", 10)
		i, _ := strconv.ParseInt(bin, 2, 64)
		found[i] = true
	}
	for i := int64(40); i < 801; i++ {
		if !found[i] {
			fmt.Printf("%d\n", i)
		}
	}
}
