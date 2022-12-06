package twenttwo

import (
	"fmt"

	"github.com/isacikgoz/slices"
)

func moveCratesP2(fromCol int, toCol int, amount int, crates [][]string) [][]string {
	fromLen := len(crates[fromCol])
	from := crates[fromCol][fromLen-amount:]

	crates[toCol] = append(crates[toCol], from...)

	for idx := fromLen; idx > fromLen-amount; idx-- {
		_, crates[fromCol] = slices.Pop(crates[fromCol])
	}

	return crates
}

func D5p2(items []string) {
	crates := parseCrates(items)

	for _, line := range items[len(crates)+1:] {
		if line == "" {
			continue // <3 Aoz
		}

		crates = performOperation(line, crates, moveCratesP2)
	}

	fmt.Printf("%s\n", getTopCrates(crates))
}
