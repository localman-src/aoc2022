package one

import (
	"bufio"
	"fmt"
	"os"
)

func DayOneSolution() (int, error) {
	f, err := os.Open("./input")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("Blank Line")
		}
	}

	return 0, nil
}
