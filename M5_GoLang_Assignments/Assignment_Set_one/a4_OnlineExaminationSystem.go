/*
Exercise 4: Online Examination System
Topics Covered: Go Loops, Go Break and Continue, Go Constants, Go Strings, Go
Functions, Go Errors
Case Study:
Develop an online examination system where users can take a quiz.
1. Question Bank: Define a slice of structs to store questions. Each question
should have a question string, options (array), and the correct answer.
2. Take Quiz: Use loops to iterate over questions and display them one by one.
Allow the user to select an answer by entering the option number.
o Use continue to skip invalid inputs and prompt the user again.
o Use break to exit the quiz early if the user enters a specific command
(e.g., "exit").
3. Score Calculation: After the quiz, calculate the user's score and display it. Use
conditions to classify performance (e.g., "Excellent", "Good", "Needs
Improvement").
4. Error Handling: Handle errors like invalid input during the quiz (e.g., entering a
non-integer value for an option).
Bonus:
• Add a timer for the quiz, limiting each question to a fixed amount of time.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

var questionBank = []Question{
	{
		Question: "Which keyword is used to declare a variable in Go?",
		Options:  [4]string{"1. var", "2. let", "3. declare", "4. dim"},
		Answer:   1,
	},
	{
		Question: "What is the default value of an uninitialized integer variable in Go?",
		Options:  [4]string{"1. 1", "2. 0", "3. -1", "4. nil"},
		Answer:   2,
	},
	{
		Question: "Which Go keyword is used to create a goroutine?",
		Options:  [4]string{"1. goroutine", "2. go", "3. async", "4. thread"},
		Answer:   2,
	},
	{
		Question: "What is the purpose of the `defer` keyword in Go?",
		Options: [4]string{
			"1. To delay the execution of a function until the surrounding function returns",
			"2. To pause the program",
			"3. To cancel a goroutine",
			"4. To handle errors",
		},
		Answer: 1,
	},
	{
		Question: "Which of the following is a valid Go data type?",
		Options:  [4]string{"1. slice", "2. tuple", "3. list", "4. map"},
		Answer:   4,
	},
}

const timeLimitSeconds = 10

func main() {
	fmt.Println("Welcome to the Online Examination System:\n")
	fmt.Println("Instructions:")
	fmt.Println("1. You have 10 seconds to answer each question.")
	fmt.Println("2. Type 'exit' anytime to end the quiz early.\n")

	score, err := takeQuiz()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nYour total score: %d/%d\n", score, len(questionBank))
	displayPerformance(score, len(questionBank))
}

func takeQuiz() (int, error) {
	score := 0

	for i, question := range questionBank {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Question)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		startTime := time.Now()

		answer, err := getUserAnswer(startTime)
		if err != nil {
			if err.Error() == "timeout" {
				fmt.Println("Time's up! Moving to the next question.")
				continue
			}
			if err.Error() == "exit" {
				fmt.Println("Exiting the quiz. Thank you for participating!")
				return score, nil
			}
			fmt.Println("Invalid input. Skipping question.")
			continue
		}

		if answer == question.Answer {
			score++
		}
	}

	return score, nil
}

func getUserAnswer(startTime time.Time) (int, error) {
	var input string
	fmt.Print("Enter your answer (1-4) or 'exit' to quit: ")
	fmt.Scanln(&input)

	elapsed := time.Since(startTime)
	if elapsed.Seconds() > float64(timeLimitSeconds) {
		return 0, errors.New("timeout")
	}

	if input == "exit" {
		return 0, errors.New("exit")
	}

	answer, err := strconv.Atoi(input)
	if err != nil || answer < 1 || answer > 4 {
		return 0, errors.New("invalid input")
	}

	return answer, nil
}

func displayPerformance(score, total int) {
	percentage := float64(score) / float64(total) * 100

	switch {
	case percentage >= 90:
		fmt.Println("Performance: Excellent!")
	case percentage >= 75:
		fmt.Println("Performance: Very Good.")
	case percentage >= 50:
		fmt.Println("Performance: Good.")
	default:
		fmt.Println("Performance: Needs Improvement.")
	}
}
