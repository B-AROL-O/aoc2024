package main

import (
	"fmt"
	"os"
	"slices"
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
		(*a) = (*a) >> opVal
	case 1: // bxl
		(*b) = (*b) ^ operand
	case 2: // bst
		opVal := getComboOpVal(operand, *a, *b, *c)
		(*b) = opVal & 7
	case 3: // jnz
		if (*a) != 0 {
			return true
		}
	case 4: // bxc
		(*b) = (*b) ^ (*c)
	case 5: // out
		opVal := getComboOpVal(operand, *a, *b, *c)
		fmt.Printf("%d,", opVal&7)
	case 6: // bdv
		opVal := getComboOpVal(operand, *a, *b, *c)
		(*b) = (*a) >> opVal
	case 7: // cdv
		opVal := getComboOpVal(operand, *a, *b, *c)
		(*c) = (*a) >> opVal
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

func checkEquation(x byte, res byte) bool {
	// it would be:
	// ((x ^ 010) ^ ((a >> i) >> (x ^ 010)) ^ 011) & 111 = res
	// but (a >> i) can be replaced with (a >> i) & 111, that is x
	// because we just care about the last 3 bits of all this operations
	// therefore the compacted equation is:
	// ((x ^ 010) ^ (x >> (x ^ 010)) ^ 011) & 111 = res
	this := ((x ^ 0b010) ^ (x >> (x ^ 0b010)) ^ 0b011) & 0b111
	fmt.Println(this)
	return this == res
}

func part2(lines []string) {
	results := []byte{2, 4, 1, 2, 7, 5, 4, 7, 1, 3, 5, 5, 0, 3, 3, 0}
	a := 0
	for i := 0; i <= 45; i += 3 {
		res := results[i/3]
		// x = ( a >> i ) & 111
		// This extracts the 3 bits starting at position i in a's binary
		// representation.
		valid := false
		for x := 0; x < 8; x++ {
			if checkEquation(byte(x), res) {
				a |= (x << i)
				valid = true
				break
			}
		}
		if !valid {
			panic("no valid x found for i=" + strconv.Itoa(i))
		}
	}
	fmt.Println(a)
	fmt.Println(1 << 45)
}

func part2gpt() {
	// The given output sequence
	results := []byte{2, 4, 1, 2, 7, 5, 4, 7, 1, 3, 5, 5, 0, 3, 3, 0}
	a := 0

	// Reconstruct a in reverse
	shift := 0
	for i := len(results) - 1; i >= 0; i-- {
		b := int(results[i]) // b & 111 is the output
		b ^= 0b011           // Reverse final XOR
		c := a >> shift      // Recompute c as the shifted part of a
		b ^= c               // Reverse second XOR
		b ^= 0b010           // Reverse first XOR
		a |= int(b) << shift // Add the reconstructed bits to a
		shift += 3           // Shift by 3 bits for the next chunk
	}

	fmt.Printf("Reconstructed a: %d\n", a)
}

func outputSingleIteration(a int) byte {
	b := a & 0b111
	b ^= 0b010
	c := a >> b
	b ^= c
	b ^= 0b011
	return byte(b & 0b111)
}

func completeOutput(a int) []byte {
	output := make([]byte, 0, 16)
	for a > 0 {
		output = append(output, outputSingleIteration(a))
		a /= 8
	}
	return output
}

func part2b() {
	output := []byte{2, 4, 1, 2, 7, 5, 4, 7, 1, 3, 5, 5, 0, 3, 3, 0}
	tot := len(output)

	a := 1
	for i := 1; i <= tot; i++ {
		for !slices.Equal(completeOutput(a), output[tot-i:]) {
			a++
		}
		if i != tot {
			a *= 8
		}
	}

	fmt.Println(a)
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
	part2b()
	fmt.Printf("part2: %s\n", time.Since(start))
	// part2gpt()
}
