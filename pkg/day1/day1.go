package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Part2() {
	f, err := os.Open("./input/day01b.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	prev := math.MinInt32
	count := -1
	i := -1
	var summer []int
	for scanner.Scan() {
		i = (i + 1) % 3

		if line := scanner.Text(); line != "" {
			x, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("%v", err)
			}
			if len(summer) < 3 {
				summer = append(summer, x)
				fmt.Printf("%v \n", summer)
			}
			if len(summer) < 3 {
				continue
			}
			summer[i] = x
			fmt.Printf("%v \n", summer)
			curr := sum(summer)
			if curr > prev {
				count++
				// fmt.Printf("Bigger %v < %v\n", prev, curr)
			}
			prev = curr
		}
	}

	fmt.Printf("The result is %v\n", count)
}

func Part1() {
	f, err := os.Open("./input/day01a.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	prev := math.MinInt32
	count := -1
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			curr, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("%v", err)
			}
			if curr > prev {
				count++
				fmt.Printf("Bigger %v < %v\n", prev, curr)
			}
			prev = curr
		}
	}

	fmt.Printf("The result is %v\n", count)
}
