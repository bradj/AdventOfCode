package twenttwo

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/isacikgoz/slices"
)

func parseCrates(items []string) [][]string {
	cols := len(items[0])/4 + 1
	crates := make([][]string, cols)

	for _, line := range items {
		if _, err := strconv.Atoi(string(line[1])); err == nil {
			break
		}

		col := 0

		for col*4 < len(line)-1 {
			if crates[col] == nil {
				crates[col] = []string{}
			}

			start := col * 4
			end := start + 4
			if end > len(line) {
				end--
			}

			val := string(line[start:end][1])
			if val != " " {
				crates[col] = append(crates[col], val)
			}

			col++
		}
	}

	// reverse crates
	for idx := range crates {
		for i, j := 0, len(crates[idx])-1; i < j; i, j = i+1, j-1 {
			crates[idx][i], crates[idx][j] = crates[idx][j], crates[idx][i]
		}
	}

	return crates
}

func getTopCrates(crates [][]string) []string {
	topCrates := []string{}

	for _, col := range crates {
		_, topCrate := getTopCrate(col)
		topCrates = append(topCrates, topCrate)
	}

	return topCrates
}

func getTopCrate(stack []string) (int, string) {
	idx := len(stack) - 1
	return idx, stack[idx]
}

func moveCrates(fidx int, tidx int, amount int, crates [][]string) [][]string {
	moved := 0

	for moved < amount {
		_, crate := getTopCrate(crates[fidx])
		crates[tidx] = slices.Push(crates[tidx], crate)
		_, crates[fidx] = slices.Pop(crates[fidx])
		moved++
	}

	return crates
}

func performOperation(line string, crates [][]string) [][]string {
	parts := strings.Split(line, " ")
	total, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	from, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Fatal(err)
	}
	from--
	to, err := strconv.Atoi(parts[5])
	if err != nil {
		log.Fatal(err)
	}
	to--

	moveCrates(from, to, total, crates)

	return crates
}

func D5p1(items []string) {
	crates := parseCrates(items)

	for _, line := range items[len(crates)+1:] {
		if line == "" {
			continue // <3 Aoz
		}

		crates = performOperation(line, crates)
	}

	fmt.Printf("%s\n", getTopCrates(crates))
}
