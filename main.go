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
	config := quiz.NewConfig()
	runner := quiz.NewRunner(config)

	return runner.RunQuiz()
}
