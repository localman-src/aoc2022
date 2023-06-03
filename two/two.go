package two

import (
	"bufio"
	"os"
	"strings"
)

type RPSSimulator struct {
	opponent map[string]string
	player   map[string]string
	rules    map[string]string
	score    map[string][3]int
}

func (s *RPSSimulator) Result(first, second string) int {
	if s.rules[first] == second {
		return 2
	}

	if s.opponent[first] == s.player[second] {
		return 1
	}

	return 0
}

func DayTwoPartOneSolution() (int, error) {
	rpsSim := RPSSimulator{
		opponent: map[string]string{
			"A": "Rock",
			"B": "Paper",
			"C": "Scissors",
		},
		player: map[string]string{
			"X": "Rock",
			"Y": "Paper",
			"Z": "Scissors",
		},
		rules: map[string]string{
			"A": "Y",
			"B": "Z",
			"C": "X",
		},

		score: map[string][3]int{
			"X": {1, 4, 7},
			"Y": {2, 5, 8},
			"Z": {3, 6, 9},
		},
	}

	f, err := os.Open("./two/input")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		sum += rpsSim.score[round[1]][rpsSim.Result(round[0], round[1])]

	}
	err = scanner.Err()
	if err != nil {
		return 0, err
	}

	return sum, nil

}

func DayTwoPartTwoSolution() {

}
