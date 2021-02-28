package cmd

import (
	"strconv"
	"time"

	"github.com/mniak/trimv/internal"
	"github.com/mniak/trimv/pkg"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "trimv",
	Short: "Trim video",
	Run: func(cmd *cobra.Command, args []string) {

		inputPath, err := cmd.Flags().GetString("input")
		internal.Handle(err)

		outputPath, err := cmd.Flags().GetString("output")
		internal.Handle(err)

		introDurationSeconds, err := cmd.Flags().GetFloat32("intro-duration")
		internal.Handle(err)
		introDuration := time.Duration(introDurationSeconds * float32(time.Second))

		outroDurationSeconds, err := cmd.Flags().GetFloat32("outro-duration")
		internal.Handle(err)
		outroDuration := time.Duration(outroDurationSeconds * float32(time.Second))

		// begin

		ffprobeResponse, err := pkg.Ffprobe(inputPath)
		internal.Handle(err)

		durationSeconds, err := strconv.ParseFloat(ffprobeResponse.Format.Duration, 32)
		internal.Handle(err)
		duration := time.Duration(durationSeconds * float64(time.Second))

		err = pkg.Trim(inputPath, outputPath, introDuration, duration-introDuration-outroDuration)
		internal.Handle(err)
	},
}

func init() {
	RootCmd.Flags().StringP("input", "i", "", "Input file path")
	RootCmd.MarkFlagRequired("input")

	RootCmd.Flags().StringP("output", "o", "", "Output file path")
	RootCmd.MarkFlagRequired("output")

	RootCmd.Flags().Float32P("intro-duration", "l", 15, "Specifies the intro duration")
	RootCmd.Flags().Float32P("outro-duration", "r", 15, "Specifies the outro duration")
}
