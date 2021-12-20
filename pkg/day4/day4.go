package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inOne = "./input/day04b.txt"

type Pair struct {
	x, y int
}

type Board struct {
	matrix   [][]bool
	values   map[string]Pair
	isWinner bool
}

func newBoard(input []string) *Board {

	values := make(map[string]Pair)
	for i, s := range input {
		entries := strings.Fields(s)
		for j, v := range entries {
			if _, ok := values[v]; ok {
				log.Fatalf("matrix has duplicate entries %v", v)
			}
			values[v] = Pair{
				x: i,
				y: j,
			}
		}
	}
	return &Board{
		values: values,
		matrix: [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
	}
}

func (board *Board) call(input string) bool {
	if pair, ok := board.values[input]; ok {
		delete(board.values, input)
		board.matrix[pair.x][pair.y] = true
		return board.check(pair.x, pair.y)
	}
	return false
}

func (board *Board) score() int {
	sum := 0
	for k, _ := range board.values {
		val, err := strconv.Atoi(k)
		if err != nil {
			log.Fatalf("Something didn't conver %v", k)
		}
		sum += val
	}
	return sum
}

func (board *Board) check(x, y int) bool {
	// fmt.Printf("checking for winner %v\n", board)
	resultx := true
	resulty := true
	for i := 0; i < 5; i++ {
		resultx = resultx && board.matrix[i][y]
		resulty = resulty && board.matrix[x][i]
	}
	return resulty || resultx
}

func (b Board) String() string {
	var m []string
	for _, row := range b.matrix {
		m = append(m, fmt.Sprintf("%v", row))
	}
	return fmt.Sprintf("\nWinner:%v\n%v\n\n%v", b.isWinner, b.values, strings.Join(m, "\n"))
}

func Part1() {
	f, err := os.Open(inOne)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input []string
	var boards []*Board

	var lines []string
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			fmt.Printf("XCANNER %v\n", line)
			lines = append(lines, line)
		} else {
			fmt.Printf("SCANNER %v\n", line)
			if len(lines) == 1 {
				fmt.Printf("Creating input %v\n", lines)
				input = strings.Split(lines[0], ",")
			} else {
				boards = append(boards, newBoard(lines))
			}
			lines = []string{}
		}
	}

loop:
	for _, i := range input {
		for _, board := range boards {
			fmt.Printf("Calling %v\n", i)
			if board.call(i) {
				score := board.score()
				fmt.Printf("We have a winner %v\n", score)
				fmt.Printf("We have a winner %v\n", board)
				val, err := strconv.Atoi(i)
				if err != nil {
					log.Fatalf("Something didn't conver %v", i)
				}
				fmt.Printf("The answer is %v\n", score*val)
				break loop
			}
		}
	}
}

func Part2() {
	f, err := os.Open(inOne)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input []string
	var boards []*Board

	var lines []string
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			fmt.Printf("XCANNER %v\n", line)
			lines = append(lines, line)
		} else {
			fmt.Printf("SCANNER %v\n", line)
			if len(lines) == 1 {
				fmt.Printf("Creating input %v\n", lines)
				input = strings.Split(lines[0], ",")
			} else {
				boards = append(boards, newBoard(lines))
			}
			lines = []string{}
		}
	}

	for _, i := range input {
		for _, board := range boards {
			fmt.Printf("Calling %v\n", i)
			if board.isWinner {
				continue
			}
			if board.call(i) {
				score := board.score()
				fmt.Printf("We have a winner %v\n", score)
				fmt.Printf("We have a winner %v\n", board)
				val, err := strconv.Atoi(i)
				if err != nil {
					log.Fatalf("Something didn't conver %v", i)
				}
				fmt.Printf("The answer is %v\n", score*val)

				board.isWinner = true
			}
		}
	}
}
