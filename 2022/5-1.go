package twenttwo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
			if val == " " {
				crates[col] = append(crates[col], "")
			} else {
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
	// fmt.Printf("stack: %v\n", stack)
	idx := len(stack) - 1
	for idx >= 0 {
		val := stack[idx]
		if val != "" {
			return idx, val
		}
		idx--
	}

	return 0, ""
}

func moveCrates(fidx int, tidx int, amount int, crates *[][]string) {
	// fmt.Printf("Moving %d crates from %d to %d\n", amount, fidx+1, tidx+1)
	moved := 0

	for moved < amount {
		movedIdx, crate := getTopCrate((*crates)[fidx])
		toIdx, _ := getTopCrate((*crates)[tidx])
		if toIdx+1 == len((*crates)[tidx]) {
			(*crates)[tidx] = append((*crates)[tidx], crate)
		} else {
			(*crates)[tidx][toIdx+1] = crate
		}
		(*crates)[fidx][movedIdx] = ""
		moved++
	}
}

func performOperation(line string, crates *[][]string) {
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
}

func D5p1(items []string) {
	crates := parseCrates(items)

	// for _, line := range items[len(crates)+1 : len(crates)+10] {
	for _, line := range items[len(crates)+1:] {
		if line == "" {
			continue // <3 Aoz
		}

		performOperation(line, &crates)
	}

	fmt.Printf("%s\n", getTopCrates(crates))
}
