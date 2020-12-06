package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ok := 0
	nFound := 0
	cidFound := false
	corrects := true
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if line == "" {
			if corrects && (nFound == 8 || (!cidFound && nFound == 7)) {
				ok++
			}
			nFound = 0
			cidFound = false
			corrects = true
			continue
		}

		for _, p := range parts {
			subparts := strings.Split(p, ":")
			//			fmt.Printf("%s\n", subparts[0])
			field := subparts[0]
			if field == "cid" {
				cidFound = true
			}
			val := subparts[1]
			switch field {
			case "byr":
				if !validByr(val) {
					corrects = false
					fmt.Printf("err byr %s\n", val)
				}
			case "iyr":
				if !validIyr(val) {
					corrects = false
					fmt.Printf("err iyr %s\n", val)
				}
			case "eyr":
				if !validEyr(val) {
					corrects = false
					fmt.Printf("err eyr %s\n", val)
				}
			case "hgt":
				if !validHgt(val) {
					corrects = false
					fmt.Printf("err hgt %s\n", val)
				}
			case "hcl":
				if !validHcl(val) {
					corrects = false
					fmt.Printf("err hcl %s\n", val)
				}
			case "ecl":
				if !validEcl(val) {
					corrects = false
					fmt.Printf("err ecl %s\n", val)
				}
			case "pid":
				if !validPid(val) {
					corrects = false
					fmt.Printf("err pid %s\n", val)
				}
			}
			nFound++
		}
	}
	if corrects && (nFound == 8 || (!cidFound && nFound == 7)) {
		ok++
	}
	fmt.Printf("%d\n", ok)
}

func validByr(s string) bool {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if v < 1920 || v > 2002 {
		return false
	}
	return true
}

func validIyr(s string) bool {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if v < 2010 || v > 2020 {
		return false
	}
	return true
}

func validEyr(s string) bool {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if v < 2020 || v > 2030 {
		return false
	}
	return true
}

func validHgt(s string) bool {
	if strings.HasSuffix(s, "cm") {
		v, err := strconv.Atoi(s[0:3])
		if err != nil {
			return false
		}
		if v < 150 || v > 193 {
			return false
		}
		return true

	}
	if strings.HasSuffix(s, "in") {
		v, err := strconv.Atoi(s[0:2])
		if err != nil {
			return false
		}
		if v < 59 || v > 76 {
			return false
		}
		return true

	}
	return false
}

func validHcl(s string) bool {
	m, err := regexp.MatchString(`^#[0-9a-f]{6}?$`, s)
	if err != nil {
		return false
	}
	return m
}

func validEcl(s string) bool {
	m, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, s)
	if err != nil {
		return false
	}
	return m
}

func validPid(s string) bool {
	m, err := regexp.MatchString(`^[0-9]{9}?$`, s)
	if err != nil {
		fmt.Printf("err regexp\n")
		return false
	}
	return m
}
