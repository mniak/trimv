package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "trimv",
	Run: func(cmd *cobra.Command, args []string) {

		inputPath, err := cmd.Flags().GetString("input")
		handle(err)

		outputPath, err := cmd.Flags().GetString("output")
		handle(err)

		ffprobeResponse, err := ffprobe(inputPath)
		handle(err)

		duration, err := strconv.ParseFloat(ffprobeResponse.Format.Duration, 32)
		handle(err)

		err = trim(inputPath, outputPath, INTRO_SIZE, duration-(INTRO_SIZE*2))
		handle(err)
	},
}
