package quiz

import (
	"fmt"
	"math/rand"
	"time"
)

func RunSumQuiz(max int) (success bool) {

	rand.Seed(time.Now().UnixNano())
	summand1 := rand.Intn(max - 1)
	summand2 := rand.Intn(max-summand1) + 1
	quizString := fmt.Sprint(summand1, " + ", summand2, " = ")

	answer := GetAnswer(quizString)
	expectedAnswer := summand2 + summand1

	success = EvaluateSolution(expectedAnswer, answer, quizString)

	return success
}
