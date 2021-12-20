package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const in = "../../input/day06a.txt"

func main() {
	f, err := os.Open(in)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var buffer [9]int

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			fish := strings.Split(line, ",")
			for _, f := range fish {
				i, e := strconv.Atoi(f)
				if e != nil {
					log.Fatalf("%v", e)
				}
				buffer[i] = buffer[i] + 1
			}
		}
	}

	for day := 0; day < 256; day++ {
		zeros := buffer[0]
		for i := 1; i < len(buffer); i++ {
			buffer[i-1] = buffer[i]
		}
		buffer[6] += zeros
		buffer[8] = zeros
	}

	result := 0
	for _, v := range buffer {
		result += v
	}

	// fmt.Printf("The result is: %v", buffer)
	fmt.Printf("The result is: %v\n", result)
}
