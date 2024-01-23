package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"math/rand"
	"time"
)

const QuizTypeSub = "sub"

type Subtraction struct{}

var _ Runnable = &Subtraction{}

func NewSubtractionQuiz() *Subtraction {
	return &Subtraction{}
}

func (s Subtraction) Run(max int) bool {

	var success bool

	rand.Seed(time.Now().UnixNano())
	minuend := rand.Intn(max-1) + 1
	subtrahend := rand.Intn(minuend)
	question := fmt.Sprint(minuend, " - ", subtrahend, " = ")

	answer := util.GetAnswerInt(question)
	expectedAnswer := minuend - subtrahend

	success = util.EvaluateSolutionInt(expectedAnswer, answer, question)

	return success
}
