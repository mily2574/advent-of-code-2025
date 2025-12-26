package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parts(input string) (sum1 int, sum2 int) {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	for _, r := range ranges {
		bounds := strings.Split(r, "-")

		firstID, _ := strconv.Atoi(bounds[0])
		lastID, _ := strconv.Atoi(bounds[1])

		for n := firstID; n <= lastID; n++ {
			s := strconv.Itoa(n)
			if isInvalidIDPart1(s) {
				sum1 += n
			}
			if isInvalidIDPart2(s) {
				sum2 += n

			}
		}
	}
	return sum1, sum2
}

func isInvalidIDPart2(s string) bool {
	l := len(s)
	//All possible lengths of the basic pattern
	for k := 1; k <= l/2; k++ {
		//k is the length of sequence possible, at least 2 repetitions are needed
		if l%k != 0 {
			continue
		}

		pattern := s[:k]
		//Check of all repetitions
		valid := true
		for i := k; i < l; i += k {
			if s[i:i+k] != pattern {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}

func isInvalidIDPart1(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	sol1, sol2 := parts(string(input))
	fmt.Printf("Solution part1: %d\n", sol1)
	fmt.Printf("Solution part2: %d\n", sol2)
}
