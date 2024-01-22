package quiz

import (
	"flag"
)

type Config struct {
	Debug      bool
	Difficulty int
	QuizType   string
}

func NewConfig() *Config {

	var debug = flag.Bool("debug", false, "Enable Debug logging")
	var difficulty = flag.Int("difficulty", 3, "the initial difficulty")
	var quizType = flag.String("type", "", "")

	flag.Parse()

	config := Config{
		Debug:      *debug,
		Difficulty: *difficulty,
		QuizType:   *quizType,
	}

	return &config
}
