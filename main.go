package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sumOfSquares(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	n := numbers[0]
	// Skip negative numbers
	if n < 0 {
		return sumOfSquares(numbers[1:])
	}
	return n*n + sumOfSquares(numbers[1:])
}

func readInts(scanner *bufio.Scanner, X int) []int {
	if !scanner.Scan() {
		log.Fatal("failed to scan input")
	}
	inputs := strings.Fields(scanner.Text())
	integers := make([]int, X)
	if len(inputs) != X {
		log.Fatal("wrong number of inputs")
	}
	for i, input := range inputs {
		val, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("failed to convert", "err", err)
		}
		if val < -100 || val > 100 {
			log.Fatal("integer out of range: -100 <= Yn <= 100", "err", err)
		}
		integers[i] = val
	}
	return integers
}

func main() {
	// Creating a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("failed to scan the number of test cases")
	}
	// Number of test cases (1 <= N <= 100)
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("failed to convert", "err", err)
	}
	if N < 1 || N > 100 {
		log.Fatal("cases should be between 1 and 100", "err", err)
	}
	// Slice of resulting squares
	results := make([]int, N)
	var processCase func(int)
	processCase = func(i int) {
		if i == N {
			return
		}
		if !scanner.Scan() {
			log.Fatal("failed to scan the number of integers")
		}
		// Number of intengers for this case (0 < X <= 100)
		X, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("failed to convert", "err", err)
		}
		if X < 1 || X > 100 {
			log.Fatal("the number of integers should be between 1 and 100", "err", err)
		}
		// The integers themselves (-100 <= Yn <= 100)
		Y := readInts(scanner, X)
		if len(Y) != X {
			log.Fatal("wrong number of inputs", "err", err)
		}
		// Getting a result
		results[i] = sumOfSquares(Y)
		processCase(i + 1)
	}
	processCase(0)
	for _, result := range results {
		fmt.Println(result)
	}
}
