package quizTypes

type Runnable interface {
	Run(difficulty int) (success bool)
}
