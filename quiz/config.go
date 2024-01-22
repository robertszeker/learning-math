package quiz

import (
	"flag"
)

type Config struct {
	Debug      bool
	Difficulty int
}

func NewConfig() *Config {

	var debug = flag.Bool("debug", false, "Enable Debug logging")
	var difficulty = flag.Int("difficulty", 3, "the initial difficulty")

	flag.Parse()

	config := Config{
		Debug:      *debug,
		Difficulty: *difficulty,
	}

	return &config
}
