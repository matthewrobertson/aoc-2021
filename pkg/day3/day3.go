package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inOne = "./input/day03a.txt"

func Part2() {
	f, err := os.Open(inOne)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			lines = append(lines, line)
		}
	}
	foo := filter(lines, mostCommon)
	bar := filter(lines, leastCommon)
	fmt.Printf("Depth %v = %v\n", foo, binToInt(foo))
	fmt.Printf("Horiz %v = %v\n", bar, binToInt(bar))
	fmt.Printf("Result %v\n", binToInt(bar)*binToInt(foo))
}

func Part1() {
	f, err := os.Open(inOne)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	depth := 0
	horiz := 0
	aim := 0
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			words := strings.Fields(line)
			num, err := strconv.Atoi(words[1])
			if err != nil {
				log.Fatalf("%v", err)
			}
			switch words[0] {
			case "down":
				aim = aim + num
			case "up":
				aim = aim - num
			case "forward":
				horiz = horiz + num
				depth = depth + (num * aim)
			default:
				log.Fatalf("Unkown thing %v", words[0])
			}
		}
	}

	fmt.Printf("Depth %v\n", depth)
	fmt.Printf("Horiz %v\n", horiz)
	fmt.Printf("Result %v\n", depth*horiz)
}

type common func([]string, int) rune

func filter(lines []string, cf common) string {
	ind := 0
	for len(lines) > 1 {
		x := cf(lines, ind)
		var newLines []string
		for _, s := range lines {
			if s[ind] == byte(x) {
				newLines = append(newLines, s)
			}
		}
		fmt.Printf("%v\n", newLines)
		lines = newLines
		ind++
	}
	return lines[0]
}

func mostCommon(lines []string, i int) rune {
	ones := 0
	zers := 0
	for _, s := range lines {
		if s[i] == '1' {
			ones++
		} else {
			zers++
		}
	}

	if ones >= zers {
		return '1'
	}
	return '0'
}

func leastCommon(lines []string, i int) rune {
	if mostCommon(lines, i) == '1' {
		return '0'
	}
	return '1'
}

func binToInt(in string) int {
	result := 0
	for i, x := range in {
		if x == '1' {
			result = result | 1<<(len(in)-i-1)
		}
	}
	return result
}
