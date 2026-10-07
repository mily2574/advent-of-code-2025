package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Read input and creation of grid ([][]byte)
func readGrid(input string) [][]byte {
	var grid [][]byte
	scanner := bufio.NewScanner((strings.NewReader(input)))
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	return grid
}

// Count how many '@' symbols are in 8 adjacent cells to (r,c)
func countNeighbors(grid [][]byte, r, c int) int {
	count := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue //same cell no neighbor
			}

			nr := r + dr
			nc := c + dc

			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[nr]) {
				if grid[nr][nc] == '@' {
					count++
				}
			}
		}
	}
	return count
}

func part1(grid [][]byte) int {
	rolls := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '@' {
				if countNeighbors(grid, r, c) < 4 {
					rolls++
				}
			}
		}
	}
	return rolls
}

type Point struct {
	r, c int
}

func part2(grid [][]byte) int {
	total := 0
	for {
		var toRemove []Point
		for r := range grid {
			for c := range grid {
				if grid[r][c] == '@' {
					if countNeighbors(grid, r, c) < 4 {
						toRemove = append(toRemove, Point{r, c})
					}
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}
		for _, p := range toRemove {
			grid[p.r][p.c] = '.' //change from '@' to '.'
		}
		total += len(toRemove)
	}
	return total
}
func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	grid1 := readGrid(string(input))
	fmt.Println("Accessible rolls:", part1(grid1))
	grid2 := readGrid(string(input))
	fmt.Println("Total rolls removed:", part2(grid2))
}
