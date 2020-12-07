//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//
//	mapp := map[string][]string{}
//
//	for scanner.Scan() {
//		line := scanner.Text()
//
//		col, cols := scanLine(line)
//		//fmt.Printf("%s: %+v\n", col, cols)
//		mapp[col] = cols
//	}
//
//	n := contains(mapp, "shiny gold")
//	for _, c := range n {
//		fmt.Printf("'%s'\n", strings.Trim(c, " "))
//	}
//}
//
//func contains(mapp map[string][]string, s string) []string {
//	res := []string{}
//	for k, vals := range mapp {
//		for _, val := range vals {
//			if val == s {
//				res = append(res, k)
//				//				fmt.Printf(" = '%s'\n", k)
//				res = append(res, contains(mapp, strings.Trim(k, " "))...)
//			}
//		}
//	}
//	return res
//}
//
//func scanLine(line string) (string, []string) {
//	parts := strings.Split(line, "contain")
//
//	col := parts[0]
//	col = strings.Replace(col, " bags", "", 1)
//	col = strings.Replace(col, " bag", "", 1)
//	//	fmt.Printf("'%s'\n", col)
//
//	if parts[1] == " no other bags." {
//		return col, []string{}
//	}
//
//	parts = strings.Split(parts[1], ",")
//	cols := []string{}
//	for _, part := range parts {
//		c := part[3:]
//		c = strings.Replace(c, " bags.", "", 1)
//		c = strings.Replace(c, " bags", "", 1)
//		c = strings.Replace(c, " bag.", "", 1)
//		c = strings.Replace(c, " bag", "", 1)
//		//		fmt.Printf("- '%s'\n", c)
//		cols = append(cols, c)
//	}
//	return col, cols
//}
//