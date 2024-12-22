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

func mix(a int, b int) int {
	return a ^ b
}

func prune(x int) int {
	return x & (16777216 - 1)
}

func nextSecret(secret int) int {
	mult64 := secret << 6
	secret = mix(mult64, secret)
	secret = prune(secret)

	div32 := secret >> 5
	secret = mix(div32, secret)
	secret = prune(secret)

	mult2048 := secret << 11
	secret = mix(mult2048, secret)
	secret = prune(secret)

	return secret
}

func calcNextSecrets(start int, iterations int) int {
	secret := start
	for range iterations {
		secret = nextSecret(secret)
	}
	return secret
}

func part1(lines []string) {
	sum := 0
	for _, l := range lines {
		secret, _ := strconv.Atoi(l)
		sum += calcNextSecrets(secret, 2000)
		// fmt.Printf("%d: %d\n", secret, finalSecret)
	}
	fmt.Println(sum)
}

type Seq struct {
	a, b, c, d int
}

func part2(lines []string) {
	prices := [][]int{}
	diffs := [][]int{}
	iterations := 2000
	for _, l := range lines {
		v := make([]int, 0, iterations)
		d := make([]int, 0, iterations)

		secret, _ := strconv.Atoi(l)
		v = append(v, secret%10)
		d = append(d, 0) // just to avoid having an offset for this array
		// fmt.Printf("%d: %d\n", secret, secret%10)
		prev := secret % 10
		for range iterations - 1 {
			secret = nextSecret(secret)
			v = append(v, secret%10)
			d = append(d, secret%10-prev%10)
			// fmt.Printf("%d: %d (%d)\n", secret, secret%10, secret%10-prev%10)
			prev = secret
		}
		prices = append(prices, v)
		diffs = append(diffs, d)
	}

	// only sequences that sum up to >= 0 make sense, so we only consider them
	soldAmount := map[Seq]int{}
	for buyer := range len(lines) {
		buyerSeqs := map[Seq]bool{}
		for i := 1; i+3 < iterations; i++ {
			s := Seq{
				a: diffs[buyer][i],
				b: diffs[buyer][i+1],
				c: diffs[buyer][i+2],
				d: diffs[buyer][i+3],
			}

			if s.a+s.b+s.c+s.d < 0 {
				continue
			}
			if buyerSeqs[s] {
				continue
			}

			buyerSeqs[s] = true
			soldAmount[s] += prices[buyer][i+3]
		}
	}

	max := 0
	for s := range soldAmount {
		if soldAmount[s] > max {
			max = soldAmount[s]
		}
	}
	fmt.Printf("%d\n", max)
}

func main() {
	data, err := os.ReadFile("./input22_def.txt")
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
