package quiz

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	Debug       bool
	Difficulty  int
	QuizType    string
	HideFunc    func()
	FilterWords string
}

// todo: add config validation
// todo: add config "all" replacing difficulty

func NewConfig() (*Config, error) {

	var debug = flag.Bool("debug", false, "Enable Debug logging")
	var difficulty = flag.Int("difficulty", 3, "the initial difficulty")
	var quizType = flag.String("type", "", "can be one of: sum, sub, spelling, letter\nsum - train addition\nsub - train subtraction\nspelling - train spelling complete word\nletter - train guessing word by providing last letter")
	var hideType = flag.String("hide", "", "if given, question will be hidden either after hitting enter or after duration\ncan be one of: enter, [0-9]+[smh]")
	var filterWords = flag.String("filter", "", "filter words by given string")

	flag.Parse()

	hideFunc, err := getHideFunc(hideType)
	if err != nil {
		return nil, fmt.Errorf("failed to get hide func: %w", err)
	}

	config := Config{
		Debug:       *debug,
		Difficulty:  *difficulty,
		QuizType:    *quizType,
		HideFunc:    hideFunc,
		FilterWords: *filterWords,
	}

	return &config, nil
}

// todo: move hideFunc to... well... not here

func getHideFunc(hideType *string) (func(), error) {
	switch {
	case emptyPointer(hideType):
		return func() {}, nil
	case *hideType == "enter":
		return func() {
			waitForPressingEnter()
			fmt.Print("\033[H\033[2J")
		}, nil
	case isDuration(*hideType):
		duration, _ := time.ParseDuration(*hideType)
		return func() {
			time.Sleep(duration)
			fmt.Print("\033[H\033[2J")
		}, nil
	default:
		return func() {}, fmt.Errorf("failed to determine hide function")
	}
}

func waitForPressingEnter() {
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	reader := bufio.NewReader(os.Stdin)
	var b = make([]byte, 1)
	reader.Read(b)
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func emptyPointer(hideType *string) bool {
	return hideType == nil || *hideType == ""
}

func isDuration(duration string) bool {
	_, err := time.ParseDuration(duration)
	return err == nil
}
