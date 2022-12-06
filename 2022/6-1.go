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
	fmt.Printf("%v\n", last4)
	if isUnique(last4) == true {
		fmt.Println(4)
		return
	}

	for idx, val := range input[4:] {
		// fmt.Printf("idx: %d, val: %c\n", idx, val)
		last4 = slices.Delete(last4, 0)
		last4 = slices.Push(last4, string(val))

		if isUnique(last4) {
			fmt.Println(idx + 5)
			return
		}
	}
}
