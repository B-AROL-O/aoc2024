package main

import (
	"fmt"
)

type Block struct {
	start  int
	length int
}

type IndexedBlock struct {
	index  int
	start  int
	length int
}

func printMem(files []Block, free []Block) {
	for i := range len(files) {
		for range files[i].length {
			fmt.Print(i)
		}
		if i < len(free) {
			for range free[i].length {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}

func printMemAsInts(files []Block, free []Block) []int {
	mem := make([]int, 0)
	if len(files) <= len(free) {
		fmt.Println("problem")
	}
	for i := range len(files) {
		for range files[i].length {
			mem = append(mem, i)
		}
		if i < len(free) {
			for range free[i].length {
				mem = append(mem, -1)
			}
		}
	}
	return mem
}

func printCompactedMem(files []Block, newFiles []IndexedBlock) {
	pos, i, j := 0, 0, 0
	for i < len(files) || j < len(newFiles) {
		if i < len(files) {
			printed := false
			for pos >= files[i].start && pos < files[i].start+files[i].length {
				fmt.Printf("%d", i)
				pos++
				printed = true
			}
			if printed {
				i++
			}
		}
		if j < len(newFiles) {
			printed := false
			for pos >= newFiles[j].start && pos < newFiles[j].start+newFiles[j].length {
				fmt.Printf("%d", newFiles[j].index)
				pos++
				printed = true
			}
			if printed {
				j++
			}
		}
	}
}

func part1WithIntervals(line string) {
	files := make([]Block, 0, len(line)/2)
	free := make([]Block, 0, len(line)/2)
	pos := 0
	for i := 0; i < len(line); i++ {
		length := int(line[i] - '0')
		if i%2 == 0 {
			files = append(files, Block{pos, length})
		} else {
			free = append(free, Block{pos, length})
		}
		pos += length
	}

	// printMem(files, free)

	// two pointers strategy:
	// left goes through free contiguous blocks ascending
	// right goes through file blocks descending
	l, r := 0, len(files)-1
	newFiles := make([]IndexedBlock, 0)
	for l < len(free) && r >= 0 {
		freeBlock := &free[l]
		fileBlock := &files[r]
		if freeBlock.start > fileBlock.start {
			break
		}
		if freeBlock.length < fileBlock.length {
			newFiles = append(newFiles, IndexedBlock{r, freeBlock.start, freeBlock.length})
			fileBlock.length -= freeBlock.length
			l++
		} else if freeBlock.length > fileBlock.length {
			newFiles = append(newFiles, IndexedBlock{r, freeBlock.start, fileBlock.length})
			freeBlock.start += fileBlock.length
			freeBlock.length -= fileBlock.length
			files = files[:r] // remove last file block (already moved to free spaces)
			r--
		} else {
			newFiles = append(newFiles, IndexedBlock{r, freeBlock.start, freeBlock.length})
			files = files[:r] // remove last file block (already moved to free spaces)
			r--
			l++ // no need to remove the free block, it will be ignored
		}
	}

	checksum := int64(0)
	for i := range len(files) {
		for j := range files[i].length {
			checksum += int64(i) * int64(files[i].start+j)
		}
	}

	for _, f := range newFiles {
		for j := range f.length {
			checksum += int64(f.index) * int64(f.start+j)
		}
	}

	fmt.Println(checksum)
}
