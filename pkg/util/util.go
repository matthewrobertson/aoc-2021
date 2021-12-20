package util

import (
	"log"
	"math"
	"strconv"
)

func Sum(input []int) int {
	result := 0
	for _, v := range input {
		result += v
	}
	return result
}

func Max(input []int) int {
	result := math.MinInt
	for _, v := range input {
		if v > result {
			result = v
		}
	}
	return result
}

func ToIntArray(input []string) []int {
	var result []int
	for _, f := range input {
		i, e := strconv.Atoi(f)
		if e != nil {
			log.Fatalf("%v", e)
		}
		result = append(result, i)
	}
	return result
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
