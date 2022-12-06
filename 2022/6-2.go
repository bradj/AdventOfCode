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

	for idx, val := range input[numUniques:] {
		uniqueList = slices.Delete(uniqueList, 0)
		uniqueList = slices.Push(uniqueList, string(val))

		if isUnique(uniqueList) {
			fmt.Println(idx + numUniques + 1)
			return
		}
	}
}
