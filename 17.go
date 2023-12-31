package main

import (
	"strconv"
	"strings"
)

func parseInput17(input string) (left, top, right, bottom int) {
	input = input[13:]
	parts := apply(strings.Split(input, ", "), func(value string) []int {
		value = strings.Split(value, "=")[1]
		return apply(strings.Split(value, ".."), func(value string) int {
			return unwrap(strconv.Atoi(value))
		})
	})
	return parts[0][0], parts[1][1], parts[0][1], parts[1][0]
}

func sum1ToN(n int) int {
	return n * (n + 1) / 2
}

func getX(x_vel int, t int) int {
	end_vel := max(0, x_vel-t)
	return sum1ToN(x_vel) - sum1ToN(end_vel)
}

func solution17Part1(input string) int {
	left, top, right, bottom := parseInput17(input)

	for y_vel := -bottom - 1; y_vel >= 0; y_vel-- {
		orig_y_vel := y_vel
		// Skip past Λ arc to the second point where y == 0
		t := y_vel*2 + 1
		y_vel = -y_vel - 1

		y := 0
		valid_ts := []int{}
		for y >= bottom {
			if y <= top {
				valid_ts = append(valid_ts, t)
			}

			y += y_vel
			y_vel -= 1
			t += 1
		}

		for _, t := range valid_ts {
			for i := 0; i <= right; i++ {
				if x := getX(i, t); x >= left && x <= right {
					return sum1ToN(orig_y_vel)
				}
			}
		}
	}

	panic("Target unreachable")
}

func solution17Part2(input string) int {
	left, top, right, bottom := parseInput17(input)

	valid_count := 0
	counted := makeSet[[2]int](0)
	for orig_y_vel := absInt(bottom) - 1; orig_y_vel >= bottom; orig_y_vel-- {
		// Skip past Λ arc to the second point where y == 0
		t := 0
		y_vel := orig_y_vel

		y := 0
		valid_ts := []int{}
		for y >= bottom {
			if y <= top {
				valid_ts = append(valid_ts, t)
			}

			y += y_vel
			y_vel -= 1
			t += 1
		}

		for _, t := range valid_ts {
			for x_vel := 0; x_vel <= right; x_vel++ {
				if counted.contains([2]int{x_vel, orig_y_vel}) {
					continue
				}

				if x := getX(x_vel, t); x >= left && x <= right {
					counted.add([2]int{x_vel, orig_y_vel})
					valid_count += 1
				}
			}
		}
	}

	return valid_count
}
