package twenttwo

import (
	"fmt"
	"strings"

	"github.com/isacikgoz/slices"
)

func isUnique(input []string) bool {
	for idx, val := range input {
		for ii := 0; ii < len(input); ii++ {
			if ii == idx {
				continue
			}

			if val == input[ii] {
				return false
			}
		}
	}

	return true
}

func D6p1(items []string) {
	input := items[0]

	last4 := strings.Split(input[:4], "")
	if isUnique(last4) {
		fmt.Println(4)
		return
	}

	for idx, val := range input[4:] {
		last4 = slices.Delete(last4, 0)
		last4 = slices.Push(last4, string(val))

		if isUnique(last4) {
			fmt.Println(idx + 5)
			return
		}
	}
}

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
