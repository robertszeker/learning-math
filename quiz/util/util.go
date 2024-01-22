package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetAnswerInt(quizString string) int {
	fmt.Print(quizString)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0
	}
	input = strings.TrimSuffix(input, "\n")

	integer, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return integer
}

func EvaluateSolutionInt(expectedAnswer int, actualAnswer int, quizString string) bool {

	success := true

	for expectedAnswer != actualAnswer {
		success = false
		fmt.Println("wrong, try again!")
		actualAnswer = GetAnswerInt(quizString)
	}

	fmt.Println("correct")

	return success
}
