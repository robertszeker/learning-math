package quiz

import (
	"fmt"
	"math/rand"
	"time"
)

func RunSubQuiz(max int) (success bool) {

	rand.Seed(time.Now().UnixNano())
	minuend := rand.Intn(max-1) + 1
	subtrahend := rand.Intn(minuend)
	quizString := fmt.Sprint(minuend, " - ", subtrahend, " = ")

	answer := GetAnswer(quizString)
	expectedAnswer := minuend - subtrahend

	success = EvaluateSolution(expectedAnswer, answer, quizString)

	return success
}
