package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type content struct {
	color string
	count int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mapp := map[string][]content{}

	for scanner.Scan() {
		line := scanner.Text()

		col, cols := scanLine(line)
		//		fmt.Printf("%s: %+v\n", col, cols)
		mapp[col] = cols
	}
	fmt.Printf("%d\n", numbags(mapp, "shiny gold"))
}

func numbags(mapp map[string][]content, s string) int {
	res := 0
	for _, val := range mapp[s] {
		//	fmt.Printf("%d\n", val.count)
		res += val.count + val.count*numbags(mapp, val.color)
	}
	return res
}

func scanLine(line string) (string, []content) {
	parts := strings.Split(line, "contain")

	col := parts[0]
	col = strings.Replace(col, " bags", "", 1)
	col = strings.Replace(col, " bag", "", 1)
	//	fmt.Printf("'%s'\n", col)

	if parts[1] == " no other bags." {
		return col, []content{}
	}

	parts = strings.Split(parts[1], ",")
	cols := []content{}
	for _, part := range parts {
		//		fmt.Printf("%s\n", part[1:2])
		count, _ := strconv.Atoi(part[1:2])
		c := part[3:]
		c = strings.Replace(c, " bags.", "", 1)
		c = strings.Replace(c, " bags", "", 1)
		c = strings.Replace(c, " bag.", "", 1)
		c = strings.Replace(c, " bag", "", 1)
		//		fmt.Printf("- %d: '%s'\n", count, c)
		cols = append(cols, content{
			color: strings.Trim(c, " "),
			count: count,
		})
	}
	return strings.Trim(col, " "), cols
}
