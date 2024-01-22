package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetAnswer(quizString string) int {
	fmt.Print(quizString)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		//fmt.Println(err)
		return 0
	}
	input = strings.TrimSuffix(input, "\n")

	integer, err := strconv.Atoi(input)
	if err != nil {
		//fmt.Println(err)
		return 0
	}
	return integer
}

func EvaluateSolution(expectedAnswer int, answer int, quizString string) bool {

	success := true

	for expectedAnswer != answer {
		success = false
		fmt.Println("wrong, try again!")
		answer = GetAnswer(quizString)
	}

	fmt.Println("correct")

	return success
}
