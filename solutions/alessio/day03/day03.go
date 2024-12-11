package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(lines []string) {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0
	for _, l := range lines {
		matches := r.FindAllStringSubmatch(l, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
		}
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	mulRegexp := `(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`
	r := regexp.MustCompile(mulRegexp)
	sum := 0
	isEnabled := true
	for _, l := range lines {
		commands := r.FindAllString(l, -1)
		for _, cmd := range commands {
			switch cmd {
			case "do()":
				isEnabled = true
			case "don't()":
				isEnabled = false
			default:
				if isEnabled {
					var a, b int
					fmt.Sscanf(cmd, "mul(%d,%d)", &a, &b)
					sum += a * b
				}
			}
		}
	}
	_, err := fmt.Println(sum)
	check(err)
}

func main() {
	data, err := os.ReadFile("./input03.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	part1(lines)
	part2(lines)
}
