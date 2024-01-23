package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
)

type Words struct{}

const QuizTypeWords = "letter"

var _ Runnable = &Words{}

func NewWordsQuiz() *Words {
	return &Words{}
}

func (w Words) Run(difficulty int) bool {
	quizWord := util.GetRandomWord(func(a string) bool { return true })
	wordWithoutLastLetter := quizWord[:len(quizWord)-1]
	answerLetter := util.GetAnswerString(wordWithoutLastLetter, func() {})
	answerWord := wordWithoutLastLetter + answerLetter

	return evaluateSolutionString(answerWord, wordWithoutLastLetter)
}

func evaluateSolutionString(answerWord string, question string) bool {
	success := true

	for util.WordList.Contains(answerWord) == false {
		success = false
		fmt.Println("Falsch! Versuche es erneut.")
		answerLetter := util.GetAnswerString(question, func() {})
		answerWord = question + answerLetter
	}

	fmt.Println("Richtig!")

	return success
}
