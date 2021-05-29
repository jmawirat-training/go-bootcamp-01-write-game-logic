package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurn = 5
)

func GetMax(a []int) int {
	max := 0
	for _, value := range a {
		if max < value {
			max = value
		}
	}
	return max
}

//DONE - Create Custom Win Message

func Messages(messageType string) string {
	message := []string{}
	if messageType == "win" {
		message = []string{
			"1! You Successfully Guess the Answer!",
			"2! You Successfully Guess the Answer!",
			"3! You Successfully Guess the Answer!",
		}
	} else if messageType == "loss" {
		message = []string{
			"1! You Failed Guess the Answer",
			"2! You Failed Guess the Answer",
			"3! You Failed Guess the Answer",
		}
	}

	rand.Seed(time.Now().UnixNano())
	return message[rand.Intn(len(message))]

}

func main() {

	//TODO - Allow User to Input 2 numbers

	userArg := os.Args[1:]

	if len(userArg) <= 0 || len(userArg) >= 3 {
		fmt.Println("[ERROR] - Input Number")
		return
	}

	userGuess := make([]int, 0)
	for _, value := range userArg {

		value, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("[ERROR] - User input unable to conver to integer")
			return
		}
		userGuess = append(userGuess, value)

	}

	//DONE - Create Seed for Random Number Generator
	rand.Seed(time.Now().UnixNano())

	//DONE - Allow User To Have nth Tries
	for n := 0; n <= maxTurn; n++ {

		//DONE - Generate Random Number
		correctAnswer := rand.Intn(GetMax(userGuess) + 1)

		//DONE - Special Bonus Message if the player wins on the first turn

		for _, value := range userGuess {

			if n == 0 && value == correctAnswer {
				fmt.Println("First Try Winner")
				return
			}

			//DONE - Check if User Input Matches the Random Number

			if value == correctAnswer {
				fmt.Println(Messages("win"))
				return
			}

		}

	}

	fmt.Println(Messages("loss"))

}
