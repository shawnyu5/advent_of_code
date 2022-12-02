package day_2

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// map of RPS : points earned
var shapes = make(map[string]int)

// map of the outcome of a round : points earned
var roundOutcome = make(map[string]int)

// representation of opponent input : RPC
var opponent = make(map[string]string)

// representation of our input : RPC
var mine = make(map[string]string)

func init() {
	shapes["rock"] = 1
	shapes["paper"] = 2
	shapes["scissors"] = 3

	roundOutcome["lost"] = 0
	roundOutcome["draw"] = 3
	roundOutcome["win"] = 6

	opponent["a"] = "rock"
	opponent["b"] = "paper"
	opponent["c"] = "scissors"

	mine["x"] = "rock"
	mine["y"] = "paper"
	mine["z"] = "scissors"
}

func Day2() {

	myTotalPoints := 0
	f := readFile()
	scanner := bufio.NewScanner(strings.NewReader(f))

	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		opponentInput, ourInput := parseInput(scanner.Text())
		winner := determineWinner(opponentInput, ourInput)

		switch winner {
		case "our":
			myTotalPoints += calculatePoints("win", ourInput)
		case "tie":
			myTotalPoints += calculatePoints("draw", "")
		}
	}
	fmt.Printf("Day2 myTotalPoints: %v\n", myTotalPoints) // __AUTO_GENERATED_PRINT_VAR__
}

// readFile reads a file and return the contents
func readFile() string {
	f, err := ioutil.ReadFile("./day_2/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}

// parseInput parses a single line of user input
// input : a single line of input
// return: tuple representation of the 2 values in the string
func parseInput(input string) (string, string) {
	opponentInput := strings.ToLower(string(input[0]))
	oursInput := strings.ToLower(string(input[2]))
	return opponent[opponentInput], mine[oursInput]
}

// determineWinner determines who won the RPS round
// opponentInput: opponent input in text form
// ourInput     : our input in text form
// return       : "opponent" for opponent win. "our" for our win, "draw" for tie. And the number of points associated with the win
func determineWinner(opponentInput string, ourInput string) string {
	const opponentWin = "opponent"
	const ourWin = "our"
	const tie = "draw"
	if opponentInput == "rock" && ourInput == "rock" {
		return tie
	} else if opponentInput == "rock" && ourInput == "scissors" {
		return opponentWin
	} else if opponentInput == "rock" && ourInput == "paper" {
		return ourWin
	} else if opponentInput == "paper" && ourInput == "paper" {
		return tie
	} else if opponentInput == "paper" && ourInput == "scissors" {
		return ourWin
	} else if opponentInput == "paper" && ourInput == "rock" {
		return opponentWin
	} else if opponentInput == "scissors" && ourInput == "rock" {
		return ourWin
	} else if opponentInput == "scissors" && ourInput == "paper" {
		return opponentWin
	} else if opponentInput == "scissors" && ourInput == "scissors" {
		return tie
	}

	// this should never happen
	panic("error input...")
}

// calculatePoints determine the points based on the result of the game
// result: can be either "win", "lost", or "draw"
// choice: the RPS that was chosen. Can either be "rock", "paper", or "scissors"
// return: number of points gained.
func calculatePoints(result string, choice string) int {
	if result != "win" && result != "lost" && result != "draw" {
		panic("Wrong result type")
	}
	return roundOutcome[result] + shapes[choice]

}
