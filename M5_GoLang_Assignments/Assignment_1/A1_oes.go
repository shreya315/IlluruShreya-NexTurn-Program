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

func takeQuiz(questions []Question) (int, error) {
	var score int
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		timer := time.NewTimer(10 * time.Second)

		answerCh := make(chan int)

		go func() {
			var userInput string
			fmt.Print("Enter your answer (1-4 or 'exit' to quit): ")
			fmt.Scanln(&userInput)

			if userInput == "exit" {
				answerCh <- -2
				return
			}

			userAnswer, err := strconv.Atoi(userInput)
			if err != nil || userAnswer < 1 || userAnswer > 4 {
				fmt.Println("Invalid input, please enter a number between 1 and 4.")
				answerCh <- -1
			} else {
				answerCh <- userAnswer
			}
		}()

		select {
		case userAnswer := <-answerCh:
			if userAnswer == -1 {
				i--
				continue
			} else if userAnswer == -2 {
				fmt.Println("You exited the quiz early.")
				return score, errors.New("quiz exited")
			} else if userAnswer == q.Answer {
				score++
			}
		case <-timer.C:
			fmt.Println("Time's up for this question!")
		}
	}
	return score, nil
}

func main() {

	questions := []Question{
		{"What is the capital of France?", [4]string{"Berlin", "Madrid", "Paris", "Rome"}, 3},
		{"Which language is used for Android development?", [4]string{"C", "Python", "Java", "Go"}, 3},
		{"Who is the founder of Microsoft?", [4]string{"Steve Jobs", "Elon Musk", "Bill Gates", "Mark Zuckerberg"}, 3},
	}

	score, err := takeQuiz(questions)
	if err != nil && err.Error() == "quiz exited" {
		return
	}

	fmt.Printf("\nYour final score is: %d/%d\n", score, len(questions))
	if score == len(questions) {
		fmt.Println("Excellent!")
	} else if score >= len(questions)/2 {
		fmt.Println("Good job!")
	} else {
		fmt.Println("Needs Improvement.")
	}
}
