package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"math/rand"
	"time"
)

const QuizTypeSum = "sum"

type Sum struct{}

var _ Runnable = &Sum{}

func NewSumQuiz() *Sum {
	return &Sum{}
}

func (s Sum) Run(maxSum int) bool {

	var success bool

	rand.Seed(time.Now().UnixNano())
	summand1 := rand.Intn(maxSum - 1)
	summand2 := rand.Intn(maxSum-summand1) + 1
	question := fmt.Sprint(summand1, " + ", summand2, " = ")

	answer := util.GetAnswerInt(question)
	expectedAnswer := summand2 + summand1

	success = util.EvaluateSolutionInt(expectedAnswer, answer, question)

	return success
}
