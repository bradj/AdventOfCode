package twenttwo

import (
	"fmt"
)

func convertRune(char rune) int {
	code := int(char)
	if code < 91 {
		code -= 38 // uppercase conversion
	} else {
		code -= 96 // lowercase conversion
	}

	return code
}

func common(group []string) (rune, error) {
	common := make(map[rune]map[int]struct{})

	for idx, line := range group {
		for _, char := range line {
			if common[char] == nil {
				common[char] = make(map[int]struct{})
			}

			common[char][idx] = struct{}{}

			if len(common[char]) == 3 {
				return char, nil
			}
		}
	}

	return ' ', fmt.Errorf("not found")
}

func D3p2(groups []string) {
	total := 0
	grpNum := 0

	for grpNum < len(groups)-1 {
		char, err := common(groups[grpNum : grpNum+3])
		if err != nil {
			panic(err)
		}

		grpNum += 3
		total += convertRune(char)
	}

	fmt.Printf("%d\n", total)
}
