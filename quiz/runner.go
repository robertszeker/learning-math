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

func (r Runner) RunQuiz() {

	difficulty := r.config.Difficulty

	for {
		if quizTypes.RunSumQuiz(difficulty) {
			difficulty++
		} else {
			difficulty--
		}
		fmt.Printf("score: %v\n\n", difficulty)
	}
}
