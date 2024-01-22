package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"math/rand"
	"time"
)

const QuizTypeSum = "sum"

type Sum struct{}

func NewSumQuiz() *Sum {
	return &Sum{}
}

func (s Sum) Run(maxSum int) bool {

	var success bool

	rand.Seed(time.Now().UnixNano())
	summand1 := rand.Intn(maxSum - 1)
	summand2 := rand.Intn(maxSum-summand1) + 1
	quizString := fmt.Sprint(summand1, " + ", summand2, " = ")

	answer := util.GetAnswer(quizString)
	expectedAnswer := summand2 + summand1

	success = util.EvaluateSolution(expectedAnswer, answer, quizString)

	return success
}
