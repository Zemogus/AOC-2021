package main

import (
	"sort"
	"strings"
)

func solution9Part1(input string) int {
	heights := parseHeights(input)
	sinks := findSinks(heights)

	sum := aggregate(sinks, 0, func(agg int, p Point, _ int) int {
		return agg + heights[p.Y][p.X]
	})
	return sum + len(sinks)
}

func parseHeights(input string) [][]int {
	heights := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		heights = append(heights, apply([]rune(line), func(char rune) int {
			return int(char) - int('0')
		}))
	}
	return heights
}

func findSinks(heights [][]int) []Point {
	sinks := make([]Point, 0)
	for i, row := range heights {
		for j, h := range row {
			if i > 0 && h >= heights[i-1][j] {
				continue
			}
			if j > 0 && h >= heights[i][j-1] {
				continue
			}
			if i < len(heights)-1 && h >= heights[i+1][j] {
				continue
			}
			if j < len(row)-1 && h >= heights[i][j+1] {
				continue
			}

			sinks = append(sinks, Point{j, i})
		}
	}

	return sinks
}

func solution9Part2(input string) int {
	heights := parseHeights(input)
	sinks := findSinks(heights)

	basin_sizes := apply(sinks, func(s Point) int {
		return getBasinSize(heights, s, makeSet[Point](0))
	})
	// Descending sort
	sort.Slice(basin_sizes, func(i, j int) bool {
		return basin_sizes[i] > basin_sizes[j]
	})

	return basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
}

func getBasinSize(heights [][]int, seed Point, seen Set[Point]) int {
	x, y := seed.X, seed.Y
	h := heights[y][x]
	seen.add(seed)

	agg := 0
	if y > 0 {
		new_p := Point{x, y - 1}
		new_h := heights[new_p.Y][new_p.X]
		if new_h < 9 && h < new_h && !seen.contains(new_p) {
			agg += getBasinSize(heights, new_p, seen)
		}
	}
	if x > 0 {
		new_p := Point{x - 1, y}
		new_h := heights[new_p.Y][new_p.X]
		if new_h < 9 && h < new_h && !seen.contains(new_p) {
			agg += getBasinSize(heights, new_p, seen)
		}
	}
	if y < len(heights)-1 {
		new_p := Point{x, y + 1}
		new_h := heights[new_p.Y][new_p.X]
		if new_h < 9 && h < new_h && !seen.contains(new_p) {
			agg += getBasinSize(heights, new_p, seen)
		}
	}
	if x < len(heights[y])-1 {
		new_p := Point{x + 1, y}
		new_h := heights[new_p.Y][new_p.X]
		if new_h < 9 && h < new_h && !seen.contains(new_p) {
			agg += getBasinSize(heights, new_p, seen)
		}
	}

	return agg + 1
}
