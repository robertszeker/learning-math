package quiz

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/quizTypes"
)

type Runner struct {
	config *Config
}

func NewRunner(config *Config) *Runner {
	return &Runner{
		config: config,
	}
}

func (r Runner) RunQuiz() error {

	difficulty := r.config.Difficulty
	q, err := r.determineQuizType()
	if err != nil {
		return fmt.Errorf("failed to determine quiz type: %w", err)
	}

	for {
		if q.Run(difficulty) {
			difficulty++
		} else {
			difficulty--
		}
		fmt.Printf("Punkte: %v\n\n", difficulty)
	}
}

func (r Runner) determineQuizType() (quizTypes.Runnable, error) {

	switch r.config.QuizType {
	case quizTypes.QuizTypeSum:
		return quizTypes.NewSumQuiz(), nil
	case quizTypes.QuizTypeSub:
		return quizTypes.NewSubtractionQuiz(), nil
	case quizTypes.QuizTypeWords:
		return quizTypes.NewWordsQuiz(), nil
	case quizTypes.QuizTypeSpelling:
		return quizTypes.NewSpellingQuiz(r.config.HideFunc, r.config.FilterWords), nil
	default:
		return nil, fmt.Errorf("quiz type unknown: %v", r.config.QuizType)
	}
}
