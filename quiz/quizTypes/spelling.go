package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"strings"
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
	wordLengthShorterThanDifficulty := func(a string) bool {
		if difficulty < 10 {
			return len(a) <= 3
		}
		steps := 5
		return len(a) <= difficulty/steps+2
	}

	filterWordsByString := func(a string) bool {
		if s.filterWords == "" {
			return true
		}

		return strings.Contains(a, s.filterWords)
	}

	filterFunc := func(a string) bool {
		return wordLengthShorterThanDifficulty(a) && filterWordsByString(a)
	}

	quizWord := util.GetRandomWord(filterFunc)

	question := fmt.Sprintf("%v\n", quizWord)
	answer := util.GetAnswerString(question, s.hideFunc)
	expectedAnswer := quizWord

	return util.EvaluateSolutionString(expectedAnswer, answer, question, s.hideFunc)
}
