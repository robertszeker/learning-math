package main

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz"
)

func main() {
	if err := mainInternal(); err != nil {
		fmt.Println("failed with error:", err.Error())
	}
}

func mainInternal() error {
	config, err := quiz.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	runner := quiz.NewRunner(config)

	return runner.RunQuiz()
}
