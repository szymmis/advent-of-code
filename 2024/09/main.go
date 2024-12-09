package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Printf("Part One: %v\n", PartOne(input))
	fmt.Printf("Part Two: %v\n", PartTwo(input))
}

type Filesystem []*Block

func ParseFilesystem(input string) Filesystem {
	var fs []*Block

	offset := 0
	for i, char := range input {
		size := int(char) - '0'
		if i%2 == 0 {
			fs = append(fs, &Block{id: i / 2, offset: offset, size: size})
		}
		offset += size
	}

	return Filesystem(fs)
}

func (fs Filesystem) Print() {
	offset := 0
	for _, f := range fs {
		if f.offset >= offset {
			print(strings.Repeat(".", f.offset-offset))
			offset = f.offset
		}
		print(strings.Repeat(strconv.Itoa(f.id), f.size))
		offset += f.size
	}
	println()
}

func (fs Filesystem) NextBlock() (int, Block) {
	offset := 0
	for index, block := range fs {
		if block.offset > offset {
			return index, Block{id: -1, offset: offset, size: block.offset - offset}
		}
		offset += block.size
	}
	return -1, Block{offset: -1, size: 0}
}

func (fs Filesystem) GetFreeBlocks() []Block {
	var blocks []Block
	offset := 0

	for _, block := range fs {
		if block.offset > offset {
			blocks = append(blocks, Block{id: -1, offset: offset, size: block.offset - offset})
		}
		offset = block.offset + block.size
	}
	return blocks
}

type Block struct {
	id     int
	offset int
	size   int
}

func (fs *Filesystem) Compress() {
	for {
		index, emptyBlock := fs.NextBlock()
		if emptyBlock.offset == -1 {
			return
		}

		last := (*fs)[len(*fs)-1]

		if emptyBlock.size >= last.size {
			last.offset = emptyBlock.offset
			*fs = Filesystem([]*Block(*fs)[0 : len(*fs)-1])
			*fs = Filesystem(slices.Insert([]*Block(*fs), index, &Block{id: last.id, offset: emptyBlock.offset, size: last.size}))
		} else {
			last.size -= emptyBlock.size
			*fs = Filesystem(slices.Insert([]*Block(*fs), index, &Block{id: last.id, offset: emptyBlock.offset, size: emptyBlock.size}))
		}
	}
}

func (fs *Filesystem) CompressWithoutFragmentation() {
	var backup []*Block = make([]*Block, len(*fs))
	copy(backup, []*Block(*fs))

	for i := len(backup) - 1; i > 0; i-- {
		block := backup[i]

		for _, free := range fs.GetFreeBlocks() {
			if free.size >= block.size && free.offset < block.offset {
				block.offset = free.offset
				slices.SortFunc([]*Block(*fs), func(a *Block, b *Block) int {
					return a.offset - b.offset
				})
				break
			}
		}
	}
}

func (fs Filesystem) Checksum() int {
	sum := 0
	for _, block := range fs {
		for n := range block.size {
			sum += block.id * (n + block.offset)
		}
	}
	return sum
}

func PartOne(input []string) int {
	fs := ParseFilesystem(input[0])
	fs.Compress()

	return fs.Checksum()
}

func PartTwo(input []string) int {
	fs := ParseFilesystem(input[0])
	fs.CompressWithoutFragmentation()

	return fs.Checksum()
}
