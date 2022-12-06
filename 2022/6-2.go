package twenttwo

import (
	"fmt"
	"strings"

	"github.com/isacikgoz/slices"
)

func D6p2(items []string) {
	input := items[0]
	numUniques := 14

	uniqueList := strings.Split(input[:numUniques], "")
	if isUnique(uniqueList) {
		fmt.Println(numUniques)
		return
	}

	for _, val := range input[numUniques:] {
		uniqueList = slices.Delete(uniqueList, 0)
		uniqueList = slices.Push(uniqueList, string(val))

		if isUnique(uniqueList) {
			// fmt.Println(idx + numUniques + 1)
			return
		}
	}
}

func D6p2_2(items []string) {
	input := items[0]
	numUniques := 14
	total := 0
	uniques := map[rune]int{}

	for idx, val := range input {
		if idx < numUniques {
			uniques[val]++
			if uniques[val] == 1 {
				total++
			} else if uniques[val] == 2 {
				total--
			}
			continue
		}

		falloffIdx := idx - numUniques
		falloff := rune(input[falloffIdx])
		incoming := rune(input[idx])

		uniques[falloff]--
		if uniques[falloff] == 1 {
			total++
		} else if uniques[falloff] == 0 {
			total--
		}

		uniques[incoming]++
		if uniques[incoming] == 1 {
			total++
		} else if uniques[incoming] == 2 {
			total--
		}

		if total == 14 {
			// fmt.Println(idx + 1)
			return
		}
	}
}
