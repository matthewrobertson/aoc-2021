package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inOne = "./input/day02a.txt"

func Part1() {
	f, err := os.Open(inOne)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	depth := 0
	horiz := 0
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			words := strings.Fields(line)
			num, err := strconv.Atoi(words[1])
			if err != nil {
				log.Fatalf("%v", err)
			}
			switch words[0] {
			case "down":
				depth = depth + num
			case "up":
				depth = depth - num
			case "forward":
				horiz = horiz + num
			default:
				log.Fatalf("Unkown thing %v", words[0])
			}
		}
	}

	fmt.Printf("Depth %v\n", depth)
	fmt.Printf("Horiz %v\n", horiz)
	fmt.Printf("Result %v\n", depth*horiz)
}

func Part2() {
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
