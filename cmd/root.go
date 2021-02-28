package cmd

import (
	"strconv"

	"github.com/mniak/trimv/internal"
	"github.com/mniak/trimv/pkg"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "trimv",
	Run: func(cmd *cobra.Command, args []string) {

		inputPath, err := cmd.Flags().GetString("input")
		internal.Handle(err)

		outputPath, err := cmd.Flags().GetString("output")
		internal.Handle(err)

		introDuration, err := cmd.Flags().GetFloat32("intro-duration")
		internal.Handle(err)

		outroDuration, err := cmd.Flags().GetFloat32("outro-duration")
		internal.Handle(err)

		// begin

		ffprobeResponse, err := pkg.Ffprobe(inputPath)
		internal.Handle(err)

		duration, err := strconv.ParseFloat(ffprobeResponse.Format.Duration, 32)
		internal.Handle(err)

		err = pkg.Trim(inputPath, outputPath, introDuration, float32(duration)-introDuration-outroDuration)
		internal.Handle(err)
	},
}
