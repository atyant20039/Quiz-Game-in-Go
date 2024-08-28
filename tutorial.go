package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Question struct {
	Text string `json:"text"`
	Options []string `json:"options"`
	Answer uint8 `json:"answer"`
}

func printEffect(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

func main() {
	FAST_TEXT_PRINT := 50 * time.Millisecond
	MID_TEXT_PRINT := 100 * time.Millisecond
	SLOW_TEXT_PRINT := 500 * time.Millisecond

	file, err := os.Open("questions.json")
	if err != nil {
		printEffect("Internal Error occured. Try again later...", FAST_TEXT_PRINT)
		return
	}
	defer file.Close()

	var questions []Question
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&questions)
	if err != nil {
		printEffect("Internal Error occured. Try again later...", FAST_TEXT_PRINT)
		return
	}

	printEffect("Welcome to ", MID_TEXT_PRINT)
	printEffect("...\n", SLOW_TEXT_PRINT)

	for _, char := range "THE QUIZ SHOW" {
		fmt.Printf("%c\t", char)
		time.Sleep(MID_TEXT_PRINT)
	}

	printEffect("\nEnter your name here: ", FAST_TEXT_PRINT)

	var name string
	fmt.Scan(&name)

	printEffect(fmt.Sprintf("Hi %s, what is your age? ", name), FAST_TEXT_PRINT)

	var age uint8
	fmt.Scan(&age)

	if (age < 18) {
		printEffect("Sorry, you are too young to take this quiz!", MID_TEXT_PRINT)
		return
	}

	printEffect("Alrighty, Let's get started then...\n", FAST_TEXT_PRINT)
	printEffect("Here comes the first questions...\n", FAST_TEXT_PRINT)

	var score uint8


	for index, quest := range questions {
		printEffect(fmt.Sprintf("Question %d:\n", index+1), FAST_TEXT_PRINT)
		printEffect(quest.Text, MID_TEXT_PRINT)
		fmt.Println()
		for no, option := range quest.Options {
			time.Sleep(SLOW_TEXT_PRINT)
			printEffect(fmt.Sprintf("Option %d: ", no+1), FAST_TEXT_PRINT)
			time.Sleep(MID_TEXT_PRINT)
			printEffect(option, FAST_TEXT_PRINT)
			fmt.Println()
		}

		fmt.Printf("Enter your answer here: ")
		var answer uint8
		fmt.Scan(&answer)

		for answer < 1 || answer > 4 {
			fmt.Printf("Invalid Response. Enter again: ")
			fmt.Scan(&answer)
		}

		printEffect("Your answer is", FAST_TEXT_PRINT)
		printEffect("...", SLOW_TEXT_PRINT)

		if answer == quest.Answer {
			score++
			printEffect("correct. Congratulations!", FAST_TEXT_PRINT)
		} else {
			printEffect("wrong. :(", FAST_TEXT_PRINT)
			printEffect(fmt.Sprintf("\nThe correct answer is Option %d\n", quest.Answer), FAST_TEXT_PRINT)
		}

		if index != len(questions)-1 {
			printEffect("\nLet's move on to next question...\n", FAST_TEXT_PRINT)
		} else {
			printEffect("\nThis is the end of the quiz.\n", FAST_TEXT_PRINT)
		}
	}

	printEffect("Your total score is: ", FAST_TEXT_PRINT)
	time.Sleep(SLOW_TEXT_PRINT)
	printEffect(fmt.Sprintf("%d\n", score), SLOW_TEXT_PRINT)

	if int(score) > (len(questions)/2)-1 {
		printEffect("Well Done! You did a great job.", FAST_TEXT_PRINT)
	} else {
		printEffect("Your score is too low! Better luck next time!", FAST_TEXT_PRINT)
	}
}