package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/mniak/trimv/internal"
	"github.com/mniak/trimv/pkg"
	"github.com/spf13/cobra"
)

var AllCmd = &cobra.Command{
	Use:   "all <input-pattern>",
	Short: "Trim all videos matching pattern",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		inputPattern := args[0]

		outputFolder, err := cmd.Flags().GetString("output")
		internal.Handle(err)

		introDurationSeconds, err := cmd.Flags().GetFloat32("intro-duration")
		internal.Handle(err)
		introDuration := time.Duration(introDurationSeconds * float32(time.Second))
		fmt.Println("Intro Duration", introDuration)

		outroDurationSeconds, err := cmd.Flags().GetFloat32("outro-duration")
		internal.Handle(err)
		outroDuration := time.Duration(outroDurationSeconds * float32(time.Second))
		fmt.Println("Outro Duration", outroDuration)

		// begin

		matches, err := filepath.Glob(inputPattern)
		internal.Handle(err)

		if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
			os.Mkdir(outputFolder, os.ModeDir)
		}

		var wg sync.WaitGroup
		for _, inputPath := range matches {
			wg.Add(1)
			go func(inputPath string) {
				defer wg.Done()

				ffprobeResponse, err := pkg.Ffprobe(inputPath)
				internal.Handle(err)

				durationSeconds, err := strconv.ParseFloat(ffprobeResponse.Format.Duration, 32)
				internal.Handle(err)
				duration := time.Duration(durationSeconds * float64(time.Second))

				outputPath := path.Join(outputFolder, filepath.Base(inputPath))
				err = pkg.Trim(inputPath, outputPath, introDuration, duration-introDuration-outroDuration)
				internal.Handle(err)

			}(inputPath)
		}
		wg.Wait()
	},
}

func init() {
	RootCmd.AddCommand(AllCmd)

	AllCmd.Flags().StringP("output", "o", "", "Output folder")
	AllCmd.Flags().Float32P("intro-duration", "l", 15, "Specifies the intro duration")
	AllCmd.Flags().Float32P("outro-duration", "r", 15, "Specifies the outro duration")
}
