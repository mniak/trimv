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

		introDuration, err := cmd.Flags().GetFloat32("intro-duration")
		handle(err)

		outroDuration, err := cmd.Flags().GetFloat32("outro-duration")
		handle(err)

		// begin

		ffprobeResponse, err := ffprobe(inputPath)
		handle(err)

		duration, err := strconv.ParseFloat(ffprobeResponse.Format.Duration, 32)
		handle(err)

		err = trim(inputPath, outputPath, introDuration, float32(duration)-introDuration-outroDuration)
		handle(err)
	},
}
