/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/robertszeker/math-problems/quiz"
	"github.com/spf13/cobra"
)

var defaultMax int = 10
var max int
var score int

var loop bool = false

// SumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if loop {
			for {
				success := quiz.RunSumQuiz(max)
				if success {
					max++
				} else {
					max--
				}
				fmt.Println("score:", max, "\n")
			}
		} else {
			quiz.RunSubQuiz(max)
		}
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
	sumCmd.Flags().IntVar(&max, "max", defaultMax, "maximum value for sum")
	sumCmd.Flags().BoolVar(&loop, "loop", loop, "run continuously")
}
