//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	max := int64(0)
//	for scanner.Scan() {
//		line := scanner.Text()
//		bin := strings.Replace(line, "B", "1", 10)
//		bin = strings.Replace(bin, "F", "0", 10)
//		bin = strings.Replace(bin, "R", "1", 10)
//		bin = strings.Replace(bin, "L", "0", 10)
//		i, _ := strconv.ParseInt(bin, 2, 64)
//		if i > max {
//			max = i
//		}
//	}
//	fmt.Printf("%d\n", max)
//}
//