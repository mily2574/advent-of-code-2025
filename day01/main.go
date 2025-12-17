package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func caseR(start int, num int) (new_start int) {
	new_start = (start + num) % 100
	return
}

func caseL(start int, num int) (new_start int) {
	new_start = start - num
	new_start %= 100
	if new_start < 0 {
		new_start += 100
	}
	return
}

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	start := 50
	countZero := 0
	for scanner.Scan() {
		rotations := scanner.Text()
		if len(rotations) < 2 {
			continue
		}
		num, err := strconv.Atoi(rotations[1:])
		if err != nil {
			continue
		}
		dir := rotations[0]

		if dir == 'R' {
			start = caseR(start, num)
		} else if dir == 'L' {
			start = caseL(start, num)
		}

		if (start % 100) == 0 {
			countZero++
		}
	}
	return countZero
}

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	start := 50
	countTotal := 0
	for scanner.Scan() {
		rotations := scanner.Text()
		if len(rotations) < 2 {
			continue
		}
		num, err := strconv.Atoi(rotations[1:])
		if err != nil {
			continue
		}
		dir := rotations[0]

		//100 clicks = 1 full rotation that passes through 0
		countTotal += num / 100
		rest := num % 100

		if dir == 'R' {
			if start != 0 && rest >= 100-start { //start + k click = 100 --> passes through 0
				countTotal++
			}
			start = caseR(start, num)
		} else if dir == 'L' {
			if start != 0 && rest >= start { //start - k click = 0
				countTotal++
			}
			start = caseL(start, num)
		} else {
			fmt.Println("Invalid rotations")
			break
		}

		/*	We can use the iteration with a FOR:
			for i := 0; i < num; i++ {
				if dir == 'R' {
					start = (start + 1) % 100
				} else if dir == 'L' {
					start = (start - 1 + 100) % 100
				} else {
					fmt.Println("Invalid rotations")
					break
				}
				if start == 0 {
					countTotal++
				}
			}
		*/
	}
	return countTotal
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution part1: %d\n", part1(string(input)))
	fmt.Printf("Solution part2: %d\n", part2(string(input)))
}
