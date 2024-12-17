package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getComboOpVal(operand int, a int, b int, c int) int {
	if operand >= 7 || operand < 0 {
		panic("invalid combo operand value: " + string(operand))
	}
	switch operand {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return operand
	}
}

func exec(opcode int, operand int, a *int, b *int, c *int) bool {
	switch opcode {
	case 0: // adv
		opVal := getComboOpVal(operand, *a, *b, *c)
		res := (*a) / (1 << opVal)
		*a = res
	case 1: // bxl
		(*b) = (*b) ^ operand
	case 2: // bst
		opVal := getComboOpVal(operand, *a, *b, *c)
		(*b) = opVal % 8
	case 3: // jnz
		if (*a) != 0 {
			return true
		}
	case 4: // bxc
		(*b) = (*b) ^ (*c)
	case 5: // out
		opVal := getComboOpVal(operand, *a, *b, *c)
		fmt.Printf("%d,", opVal%8)
	case 6: // bdv
		opVal := getComboOpVal(operand, *a, *b, *c)
		res := (*a) / (1 << opVal)
		*b = res
	case 7: // cdv
		opVal := getComboOpVal(operand, *a, *b, *c)
		res := (*a) / (1 << opVal)
		*c = res
	default:
		panic("invalid opcode: " + string(opcode))
	}

	return false
}

func solve(lines []string) {
	var a, b, c int
	fmt.Sscanf(lines[0], "Register A: %d", &a)
	fmt.Sscanf(lines[1], "Register B: %d", &b)
	fmt.Sscanf(lines[2], "Register C: %d", &c)

	cmdstr := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	cmds := make([]int, 0, len(cmdstr))
	for _, s := range cmdstr {
		n, _ := strconv.Atoi(s)
		cmds = append(cmds, n)
	}

	ptr := 0
	for ptr <= len(cmds)-1 {
		jump := exec(cmds[ptr], cmds[ptr+1], &a, &b, &c)
		if jump {
			ptr = cmds[ptr+1]
		} else {
			ptr += 2
		}
	}
}

func part1(lines []string) {
	solve(lines)

}

func part2(lines []string) {
	solve(lines)
}

func main() {
	data, err := os.ReadFile("./input17_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	start := time.Now()
	part1(lines)
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	part2(lines)
	fmt.Printf("part2: %s\n", time.Since(start))
}
