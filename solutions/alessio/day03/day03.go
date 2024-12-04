package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(lines []string) {
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	check(err)
	sum := 0
	for _, l := range lines {
		for _, s := range r.FindAllString(l, -1) {
			var a int
			var b int
			_, err := fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
			check(err)
			sum += a * b
		}
	}
	_, err = fmt.Println(sum)
	check(err)
}

func part2(lines []string) {
	mul_regexp := `(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`
	r, err := regexp.Compile(mul_regexp)
	check(err)
	sum := 0
	isEnabled := true
	for _, l := range lines {
		commands := r.FindAllString(l, -1)
		_, err := fmt.Println(commands)
		check(err)
		for _, cmd := range commands {
			switch cmd {
			case "do()":
				isEnabled = true
			case "don't()":
				isEnabled = false
			default:
				if isEnabled {
					var a int
					var b int
					_, err := fmt.Sscanf(cmd, "mul(%d,%d)", &a, &b)
					check(err)
					sum += a * b
				}
			}
		}
	}
	_, err = fmt.Println(sum)
	check(err)
}

func main() {
	data, err := os.ReadFile("./input03.txt")
	check(err)

	lines := strings.Split(strings.Trim(string(data), "\r\n"), "\r\n")

	part1(lines)
	part2(lines)
}
