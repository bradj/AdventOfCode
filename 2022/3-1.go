package twenttwo

import (
	"fmt"

	"github.com/bradj/AdventOfCode/util"
)

func D3p1() {
	items := util.GetItems("2022/3.input")
	total := 0

	for _, item := range items {
		if item == "" {
			continue
		}

		set := util.RuneSet{}
		half := len(item) / 2

		for idx, char := range item {
			if idx < half {
				set.Add(char)
				continue
			}

			if !set.Exists(char) {
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

	fmt.Printf("total: %d\n", total)
}
