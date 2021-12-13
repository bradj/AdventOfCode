package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1-1.input")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	prev := -1
	curr := -1
	inc := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		curr, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf("%s is not an integer: %v", scanner.Text(), err)
			return
		}

		if prev == -1 {
			prev = curr
			continue
		}

		if curr > prev {
			inc++
		}

		prev = curr
	}

	fmt.Printf("%d increases found\n", inc)
}
