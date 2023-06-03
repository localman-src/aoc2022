package one

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func DayOnePartOneSolution() (int, error) {
	f, err := os.Open("./one/input")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var calories []int
	total := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			calories = append(calories, total)
			total = 0
			continue
		}
		tally, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		total += tally
	}

	err = scanner.Err()
	if err != nil {
		return 0, err
	}
	max := 0
	for _, elf := range calories {
		if elf > max {
			max = elf
		}
	}

	return max, nil
}

func DayOnePartTwoSolution() (int, error) {
	f, err := os.Open("./one/input")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var calories []int
	total := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			calories = append(calories, total)
			total = 0
			continue
		}
		tally, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		total += tally
	}

	err = scanner.Err()
	if err != nil {
		return 0, err
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	topThree := calories[0] + calories[1] + calories[2]

	return topThree, nil
}
