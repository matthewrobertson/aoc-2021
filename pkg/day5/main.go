package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const in = "../../input/day05a.txt"

func main() {
	f, err := os.Open(in)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make(map[string]int)

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			pts := points(line)
			fmt.Printf("%v: %v\n", line, pts)
			for _, s := range pts {
				if n, ok := grid[s]; ok {
					grid[s] = n + 1
				} else {
					grid[s] = 1
				}
			}
		}
	}

	result := 0
	for _, v := range grid {
		if v > 1 {
			result++
		}
	}
	fmt.Printf("The result is: %v", result)
}

type Point struct {
	x int
	y int
}

func points(line string) []string {
	coords := strings.Split(line, " -> ")
	if len(coords) != 2 {
		return []string{}
	}
	left := point(coords[0])
	right := point(coords[1])
	fmt.Printf("The points are: %v %v\n", left, right)
	var res []string
	if left.x != right.x && left.y != right.y {
		c := max(left.x, right.x) - min(left.x, right.x)
		negative := (left.x-right.x)*(left.y-right.y) < 0
		for offset := 0; offset <= c; offset++ {
			x := min(left.x, right.x) + offset
			var y int
			if negative {
				y = max(left.y, right.y) - offset
			} else {
				y = min(left.y, right.y) + offset
			}

			res = append(res, fmt.Sprintf("%v-%v", x, y))
		}
	} else {
		for x := min(left.x, right.x); x <= max(left.x, right.x); x++ {
			for y := min(left.y, right.y); y <= max(left.y, right.y); y++ {
				res = append(res, fmt.Sprintf("%v-%v", x, y))
			}
		}
	}
	return res
}

func point(p string) Point {
	coords := strings.Split(p, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatalf("%v", err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	return Point{x, y}
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
