package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const in = "../../input/day09a.txt"

func main() {
	f, err := os.Open(in)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input [][]int

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			parts := strings.Split(line, "")
			var row []int
			for _, x := range parts {
				n, e := strconv.Atoi(x)
				if e != nil {
					log.Fatalf("%v", e)
				}
				row = append(row, n)
			}
			input = append(input, row)
		}
	}
	var lows []int
	for i, r := range input {
		for j, _ := range r {
			if isLow(input, i, j) {
				lows = append(lows, countBasin(input, i, j))
			}
		}
	}
	sort.Ints(lows)

	fmt.Printf("The result is: %v\n", lows)
	fmt.Printf("The result is: %v\n", lows[len(lows)-1]*lows[len(lows)-2]*lows[len(lows)-3])
}

func isLow(grid [][]int, i, j int) bool {
	val := lookup(grid, i, j)
	return val < lookup(grid, i-1, j) &&
		val < lookup(grid, i+1, j) &&
		val < lookup(grid, i, j-1) &&
		val < lookup(grid, i, j+1)
}

func lookup(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return math.MaxInt
	}
	return grid[i][j]
}

func countBasin(grid [][]int, i, j int) int {
	t := cpy(grid)
	growBasin(t, i, j, math.MinInt)
	result := 0
	for _, r := range t {
		for _, v := range r {
			if v == math.MaxInt {
				result++
			}
		}
	}
	return result
}

func growBasin(grid [][]int, i, j, prev int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len((grid)[i]) {
		return
	}
	v := (grid)[i][j]
	if v == 9 || v == math.MaxInt || v < prev {
		return
	}

	(grid)[i][j] = math.MaxInt
	growBasin(grid, i-1, j, v)
	growBasin(grid, i+1, j, v)
	growBasin(grid, i, j-1, v)
	growBasin(grid, i, j+1, v)
}

func cpy(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
