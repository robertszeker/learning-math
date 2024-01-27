package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
)

type Spelling struct {
	hideFunc    func()
	filterWords string
}

const QuizTypeSpelling = "spelling"

var _ Runnable = &Spelling{}

func NewSpellingQuiz(hideFunc func(), filterWords string) *Spelling {
	return &Spelling{
		hideFunc:    hideFunc,
		filterWords: filterWords,
	}
}

func (s Spelling) Run(difficulty int) (success bool) {

	filterFunc := func(a string) bool {
		return util.FilterWordsByLengthFunc(difficulty)(a) && util.FilterWordsBySubstringFunc(s.filterWords)(a)
	}

	quizWord := util.GetRandomWord(filterFunc)

	question := fmt.Sprintf("%v\n", quizWord)
	answer := util.GetAnswerString(question, s.hideFunc)
	expectedAnswer := quizWord

	return util.EvaluateSolutionString(expectedAnswer, answer, question, s.hideFunc)
}
