package main

import (
	"fmt"

	"github.com/localman-src/aoc2022/one"
	"github.com/localman-src/aoc2022/two"
)

func main() {
	dayOnePartOneAnswer, err := one.DayOnePartOneSolution()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Day 1 Part One Answer: %d\n", dayOnePartOneAnswer)

	dayOnePartTwoAnswer, err := one.DayOnePartTwoSolution()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Day 1 Part Two Answer: %d\n", dayOnePartTwoAnswer)

	dayTwoPartOneAnswer, err := two.DayTwoPartOneSolution()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Day 2 Part One Answer: %d\n", dayTwoPartOneAnswer)
}
