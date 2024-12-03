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
	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	sum := 0
	for _, l := range lines {
		for _, s := range r.FindAllString(l, -1) {
			var a int
			var b int
			fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
			sum += a * b
		}
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	mul_regexp := `(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`
	// do_regexp := `do`
	// dont_regexp := `don't`
	r, _ := regexp.Compile(mul_regexp)
	sum := 0
	is_enabled := true
	for _, l := range lines {
		commands := r.FindAllString(l, -1)
		fmt.Println(commands)
		for _, cmd := range commands {
			if cmd == "do()" {
				is_enabled = true
			} else if cmd == "don't()" {
				is_enabled = false
			} else if is_enabled {
				var a int
				var b int
				fmt.Sscanf(cmd, "mul(%d,%d)", &a, &b)
				fmt.Println(a, b)
				sum += a * b
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	data, err := os.ReadFile("./input03.txt")
	check(err)

	lines := strings.Split(strings.Trim(string(data), "\r\n"), "\r\n")

	// part1(lines)
	part2(lines)
}
