package twenttwo

import (
	"github.com/bradj/AdventOfCode/util"
)

func D3p1(items []string) {
	total := 0

	for _, item := range items {
		set := util.RuneSet{}
		half := len(item) / 2

		for idx, char := range item {
			if idx < half {
				set.Add(char)
				continue
			}

			if !set.Contains(char) {
				continue
			}

			code := int(char)
			if code < 91 {
				code -= 38 // uppercase conversion
			} else {
				code -= 96 // lowercase conversion
			}

			total += code
			break
		}
	}
}
