package twenttwo

import (
	"fmt"
	"strings"

	"github.com/bradj/AdventOfCode/util"
)

var outcomes map[string]map[string]string

func D2p2() {
	items := util.GetItems("2022/2.input")

	outcomes = make(map[string]map[string]string)

	outcomes[OpRock] = make(map[string]string)
	outcomes[OpRock]["X"] = Scissors // lose
	outcomes[OpRock]["Y"] = Rock     // draw
	outcomes[OpRock]["Z"] = Paper    // win

	outcomes[OpPaper] = make(map[string]string)
	outcomes[OpPaper]["X"] = Rock
	outcomes[OpPaper]["Y"] = Paper
	outcomes[OpPaper]["Z"] = Scissors

	outcomes[OpScissors] = make(map[string]string)
	outcomes[OpScissors]["X"] = Paper
	outcomes[OpScissors]["Y"] = Scissors
	outcomes[OpScissors]["Z"] = Rock

	total := 0

	for _, item := range items {
		if item == "" {
			continue
		}

		round := strings.Split(item, " ")
		myChoice := outcomes[round[0]][round[1]]
		result := choices[round[0]][myChoice] + choiceScore[myChoice]
		total += result
	}

	fmt.Printf("total: %d\n", total)
}
