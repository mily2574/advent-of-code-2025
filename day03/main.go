package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Not more two batteries, now twelve batteries
// Num of 12 digits is too big for int (32 bit), best solution is to use int64
func part2(input string) int64 {
	banks := fromInputToBanks(input)
	sumMaxJoltageTotal := int64(0)
	for _, bank := range banks {
		maxJoltage := findMaxJoltageOf12Digits(bank)
		max, _ := strconv.ParseInt(maxJoltage, 10, 64)
		sumMaxJoltageTotal += max
	}
	return sumMaxJoltageTotal
}

// Monotic stack
func findMaxJoltageOf12Digits(bank string) string {
	lenMax := 12
	stack := make([]rune, 0, lenMax)
	numToRemove := len(bank) - lenMax

	for _, digit := range bank {
		for numToRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1] //remove last digit
			numToRemove--
		}
		stack = append(stack, digit)
	}
	return string(stack[:lenMax])
}

func part1(input string) int {
	banks := fromInputToBanks(input)
	sumMaxJoltageTotal := 0
	for _, bank := range banks {
		maxJoltageNum, err := strconv.Atoi(findMaxJoltageOf2Digits(bank))
		if err != nil {
			panic(err)
		}
		sumMaxJoltageTotal += maxJoltageNum
	}
	return sumMaxJoltageTotal
}

// Brute Force
func findMaxJoltageOf2Digits(bank string) string {
	maxJoltage := ""
	joltage := ""
	l := len(bank)
	for i := 0; i < l; i++ {
		firstValue := rune(bank[i])
		for j := i + 1; j < l; j++ {
			nextValue := rune(bank[j])
			joltage = string(firstValue) + string(nextValue)
			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}
	//fmt.Println(maxJoltage)
	return maxJoltage
}

// Each line of digits in the input corresponds to a single bank of batteries
func fromInputToBanks(input string) []string {
	return strings.Fields(string(input))
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution part1: %d\n", part1(string(input)))
	fmt.Printf("Solution part2: %d\n", part2(string(input)))
}
