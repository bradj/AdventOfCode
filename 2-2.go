package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bcontent, err := ioutil.ReadFile("2-1.input")

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	items := strings.Split(content, "\n")

	x := 0
	y := 0
	aim := 0

	for _, v := range items {
		instruction := strings.Split(v, " ")

		if len(instruction) != 2 {
			continue
		}

		inc, err := strconv.Atoi(instruction[1])

		if err != nil {
			fmt.Printf("Error: %s - %v - %v", v, instruction, err)
			continue
		}

		switch instruction[0] {
		case "forward":
			x += inc

			if aim != 0 {
				y += inc * aim
			}

			break
		case "down":
			aim += inc
			break
		case "up":
			aim -= inc
			break
		}
	}

	fmt.Printf("final position: %d\n", x*y)
}
