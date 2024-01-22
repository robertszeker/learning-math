package quizTypes

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz/util"
	"math/rand"
	"time"
)

func RunSumQuiz(max int) (success bool) {

	rand.Seed(time.Now().UnixNano())
	summand1 := rand.Intn(max - 1)
	summand2 := rand.Intn(max-summand1) + 1
	quizString := fmt.Sprint(summand1, " + ", summand2, " = ")

	answer := util.GetAnswer(quizString)
	expectedAnswer := summand2 + summand1

	success = util.EvaluateSolution(expectedAnswer, answer, quizString)

	return success
}
