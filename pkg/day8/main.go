package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const in = "../../input/day08a.txt"

func main() {
	f, err := os.Open(in)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input [][]string
	var output [][]string

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			parts := strings.Split(line, " | ")
			input = append(input, SortedFields(parts[0]))
			output = append(output, SortedFields(parts[1]))
		}
	}

	result := 0
	for i, _ := range input {
		lookup := solve(input[i])
		// fmt.Printf("%v\n", lookup)
		x := toNum(output[i], lookup)
		fmt.Printf("%v\n", x)
		result += x
	}
	fmt.Printf("The result is: %v\n", result)
}

func SortedFields(s string) []string {
	var result []string
	words := strings.Fields(s)
	for _, w := range words {
		result = append(result, SortString(w))
	}
	return result

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func hasAll(super, sub string) bool {
	for _, x := range sub {
		if !strings.ContainsRune(super, x) {
			// fmt.Printf("Checking if %v in %v -> false\n", sub, super)
			return false
		}
	}
	// fmt.Printf("Checking if %v in %v -> true\n", sub, super)
	return true
}

func solve(combos []string) map[string]int {
	var one, two, three, four, five, six, seven, eight, nine, zero string
	var sixes, fives []string
	for _, x := range combos {
		if len(x) == 2 {
			one = x
		}
		if len(x) == 4 {
			four = x
		}
		if len(x) == 3 {
			seven = x
		}
		if len(x) == 7 {
			eight = x
		}
		if len(x) == 5 {
			fives = append(fives, x)
		}
		if len(x) == 6 {
			sixes = append(sixes, x)
		}
	}

	for _, x := range sixes {
		if hasAll(x, four) {
			nine = x
		} else if hasAll(x, one) {
			zero = x
		} else {
			six = x
		}
	}

	for _, x := range fives {
		if hasAll(x, one) {
			three = x
		} else if hasAll(six, x) {
			five = x
		} else {
			two = x
		}
	}

	return map[string]int{
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
		zero:  0,
	}
}

func toNum(output []string, lookup map[string]int) int {
	fmt.Printf("%v\n", output)
	// fmt.Printf("%v%v%v%v\n", output[0], output[1], output[2], output[3])
	fmt.Printf("%v\n", lookup)
	fmt.Printf("%v%v%v%v\n", lookup[output[0]], lookup[output[1]], lookup[output[2]], lookup[output[3]])
	return lookup[output[0]]*1000 + lookup[output[1]]*100 + lookup[output[2]]*10 + lookup[output[3]]
}
