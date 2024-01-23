package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
)

type Spelling struct {
	hideFunc func()
}

const QuizTypeSpelling = "spelling"

var _ Runnable = &Spelling{}

func NewSpellingQuiz(hideFunc func()) *Spelling {
	return &Spelling{
		hideFunc: hideFunc,
	}
}

func (s Spelling) Run(difficulty int) (success bool) {
	filterFunc := func(a string) bool {
		if difficulty < 10 {
			return len(a) <= 3
		}
		steps := 5
		return len(a) <= difficulty/steps+2
	}
	quizWord := util.GetRandomWord(filterFunc)

	question := fmt.Sprintf("%v\n", quizWord)
	answer := util.GetAnswerString(question, s.hideFunc)
	expectedAnswer := quizWord

	return util.EvaluateSolutionString(expectedAnswer, answer, question, s.hideFunc)
}
