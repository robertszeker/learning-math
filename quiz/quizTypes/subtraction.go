package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"math/rand"
	"time"
)

const QuizTypeSub = "sub"

type Subtraction struct{}

func NewSubtractionQuiz() *Subtraction {
	return &Subtraction{}
}

func (s Subtraction) Run(max int) bool {

	var success bool

	rand.Seed(time.Now().UnixNano())
	minuend := rand.Intn(max-1) + 1
	subtrahend := rand.Intn(minuend)
	quizString := fmt.Sprint(minuend, " - ", subtrahend, " = ")

	answer := util.GetAnswer(quizString)
	expectedAnswer := minuend - subtrahend

	success = util.EvaluateSolution(expectedAnswer, answer, quizString)

	return success
}
