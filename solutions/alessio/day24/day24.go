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

type Operation struct {
	operand1, operand2 string
	operator           string
	resDest            string
	done               bool
}

func execOp(op Operation, vals map[string]byte) byte {
	switch op.operator {
	case "AND":
		return vals[op.operand1] & vals[op.operand2]
	case "XOR":
		return vals[op.operand1] ^ vals[op.operand2]
	case "OR":
		return vals[op.operand1] | vals[op.operand2]
	default:
		fmt.Println("error: invalid operator")
		return 0
	}
}

func part1(lines []string) {
	vals := map[string]byte{}
	i := 0
	for i < len(lines) && lines[i] != "" {
		res := strings.Split(lines[i], ": ")
		val, _ := strconv.Atoi(res[1])
		vals[res[0]] = byte(val)
		i++
	}
	i++

	ops := make([]Operation, 0, len(lines)-i)
	for i < len(lines) {
		op := Operation{done: false}
		fmt.Sscanf(lines[i], "%s %s %s -> %s", &op.operand1, &op.operator, &op.operand2, &op.resDest)
		ops = append(ops, op)
		i++
	}

	doneCnt := 0
	for doneCnt < len(ops) {
		for i, op := range ops {
			if !op.done {
				_, ready1 := vals[op.operand1]
				_, ready2 := vals[op.operand2]
				if ready1 && ready2 {
					res := execOp(op, vals)
					vals[op.resDest] = res
					ops[i] = Operation{done: true}
					doneCnt++
				}
			}
		}
	}

	result := 0

	for k := range vals {
		if k[0] == 'z' {
			pos := 0
			fmt.Sscanf(k, "z%d", &pos)
			if vals[k] == 1 {
				result |= (1 << pos)
			}
		}
	}
	fmt.Println(result)
}

func execAll(opsOrig []Operation, valsOrig map[string]byte) int {
	ops := make([]Operation, len(opsOrig))
	copy(ops, opsOrig)

	vals := make(map[string]byte, len(valsOrig))
	for k, v := range valsOrig {
		vals[k] = v
	}

	doneCnt := 0
	for doneCnt < len(ops) {
		for i, op := range ops {
			if !op.done {
				_, ready1 := vals[op.operand1]
				_, ready2 := vals[op.operand2]
				if ready1 && ready2 {
					res := execOp(op, vals)
					vals[op.resDest] = res
					ops[i] = Operation{done: true}
					doneCnt++
				}
			}
		}
	}

	result := 0

	for k := range vals {
		if k[0] == 'z' {
			pos := 0
			fmt.Sscanf(k, "z%d", &pos)
			if vals[k] == 1 {
				result |= (1 << pos)
			}
		}
	}

	return result
}

func getErrBits(ops []Operation, vals map[string]byte) []string {
	errBits := []string{}
	for i := 0; i < 45; i++ {
		err := false

		if i > 0 {
			vals[fmt.Sprintf("x%02d", i-1)] = 0
			vals[fmt.Sprintf("y%02d", i-1)] = 0
		}
		for j := range ops {
			ops[j].done = false
		}

		vals[fmt.Sprintf("x%02d", i)] = 1
		vals[fmt.Sprintf("y%02d", i)] = 0

		result := execAll(ops, vals)
		resStr := fmt.Sprintf("%b", result)
		if len(resStr) != i+1 {
			err = true
			fmt.Printf("error: result string for z%02d is %d bits long\n", i, len(resStr))
		}

		ibit := result & (1 << i)
		if ibit == 0 {
			err = true
			fmt.Printf("error: expected 1 for z%02d\n", i)
		}

		vals[fmt.Sprintf("x%02d", i)] = 0
		vals[fmt.Sprintf("y%02d", i)] = 1

		result = execAll(ops, vals)
		resStr = fmt.Sprintf("%b", result)
		if len(resStr) != i+1 {
			err = true
			fmt.Printf("error: result string for z%02d is %d bits long\n", i, len(resStr))
		}

		ibit = result & (1 << i)
		if ibit == 0 {
			err = true
			fmt.Printf("error: expected 1 for z%02d\n", i)
		}

		vals[fmt.Sprintf("x%02d", i)] = 1
		vals[fmt.Sprintf("y%02d", i)] = 1

		result = execAll(ops, vals)

		ibit = result & (1 << (i + 1))
		if ibit == 0 {
			err = true
			fmt.Printf("error: expected carry bit 1 for z%02d (%b)\n", i, result)
		}
		ibit = result & (1 << i)
		if ibit != 0 {
			err = true
			fmt.Printf("error: expected sum bit 0 for z%02d (%b)\n", i, result)
		}

		if err {
			errBits = append(errBits, fmt.Sprintf("x%02d", i))
		}

		fmt.Println()
	}
	vals["x44"] = 0
	vals["y44"] = 0

	return errBits
}

func part2(lines []string) {
	vals := map[string]byte{}
	i := 0
	for i < len(lines) && lines[i] != "" {
		res := strings.Split(lines[i], ": ")
		vals[res[0]] = 0
		i++
	}
	i++

	ops := make([]Operation, 0, len(lines)-i)
	for i < len(lines) {
		op := Operation{done: false}
		fmt.Sscanf(lines[i], "%s %s %s -> %s", &op.operand1, &op.operator, &op.operand2, &op.resDest)
		ops = append(ops, op)
		i++
	}

	getErrBits(ops, vals)
}

func main() {
	data, err := os.ReadFile("./input24_def.txt")
	check(err)
	dataStr := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.Trim(dataStr, "\n"), "\n")

	start := time.Now()
	part1(lines)
	fmt.Printf("part1: %s\n", time.Since(start))
	start = time.Now()
	fmt.Printf("part2: %s\n", time.Since(start))
	part2(lines)
}
