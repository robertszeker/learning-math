package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"math/rand"
	"time"
)

func RunSubQuiz(max int) (success bool) {

	rand.Seed(time.Now().UnixNano())
	minuend := rand.Intn(max-1) + 1
	subtrahend := rand.Intn(minuend)
	quizString := fmt.Sprint(minuend, " - ", subtrahend, " = ")

	answer := util.GetAnswer(quizString)
	expectedAnswer := minuend - subtrahend

	success = util.EvaluateSolution(expectedAnswer, answer, quizString)

	return success
}
