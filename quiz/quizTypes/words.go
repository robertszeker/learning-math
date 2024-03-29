package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
)

type Words struct {
	filterWords string
}

const QuizTypeWords = "lastLetter"

var _ Runnable = &Words{}

func NewWordsQuiz(filterWords string) *Words {
	return &Words{
		filterWords: filterWords,
	}
}

func (w Words) Run(difficulty int) bool {

	filterFunc := func(a string) bool {
		return util.FilterWordsByLengthFunc(difficulty)(a) && util.FilterWordsBySubstringFunc(w.filterWords)(a)
	}

	quizWord := util.GetRandomWord(filterFunc)
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
