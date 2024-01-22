package quizTypes

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Words struct{}

const QuizTypeWords = "words"

var _ Runnable = &Words{}

func NewWordsQuiz() *Words {
	return &Words{}
}

func (w Words) Run(difficulty int) bool {
	quizWord := getRandomWord()
	quizString := quizWord[:len(quizWord)-1]
	answer := getAnswerString(quizString)
	expectedAnswer := quizWord[len(quizWord)-1:]

	return evaluateSolutionString(expectedAnswer, answer, quizString)
}

func getRandomWord() string {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Int() % len(wordList)
	return wordList[randomIndex]
}

func getAnswerString(quizString string) string {
	fmt.Print(quizString)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSuffix(input, "\n")
}

func evaluateSolutionString(expectedAnswer string, actualAnswer string, quizString string) bool {
	success := true

	for expectedAnswer != actualAnswer {
		success = false
		fmt.Println("wrong, try again!")
		actualAnswer = getAnswerString(quizString)
	}

	fmt.Println("correct")

	return success
}

// todo: move word list to config file
var wordList = []string{
	"Apfel",
	"Birne",
	"Daumen",
	"Elefant",
	"Flöte",
}
