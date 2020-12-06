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
//	ok := 0
//	nFound := 0
//	cidFound := false
//	for scanner.Scan() {
//		line := scanner.Text()
//		parts := strings.Split(line, " ")
//
//		if line == "" {
//			if nFound == 8 || (!cidFound && nFound == 7) {
//				ok++
//			}
//			nFound = 0
//			cidFound = false
//			continue
//		}
//
//		for _, p := range parts {
//			subparts := strings.Split(p, ":")
//			fmt.Printf("%s\n", subparts[0])
//			field := subparts[0]
//			if field == "cid" {
//				cidFound = true
//			}
//			nFound++
//		}
//	}
//	if nFound == 8 || (!cidFound && nFound == 7) {
//		ok++
//	}
//	fmt.Printf("%d\n", ok)
//}
