package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/matthewrobertson/aoc-2021/pkg/util"
)

const in = "../../input/day07a.txt"

func main() {
	f, err := os.Open(in)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var crabs []int

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			crabs = util.ToIntArray(strings.Split(line, ","))
		}
	}

	result := math.MaxInt
	max := util.Max(crabs)
	current := 0
	for i := 0; i <= max; i++ {
		for _, crab := range crabs {
			var dist int
			if i > crab {
				dist = i - crab
			} else {
				dist = crab - i
			}
			current += dist * (dist + 1) / 2
		}
		if current < result {
			result = current
		}
		current = 0
	}

	// fmt.Printf("The result is: %v", buffer)
	fmt.Printf("The result is: %v\n", result)
}
