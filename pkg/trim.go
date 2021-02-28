package pkg

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func Trim(inputPath, outputPath string, start, duration time.Duration) (err error) {
	cmd := exec.Command("ffmpeg",
		"-ss", fmt.Sprintf("%f", start.Seconds()),
		"-i", inputPath,
		"-t", fmt.Sprintf("%f", duration.Seconds()),
		"-c", "copy",
		"-y",
		outputPath)

	var buffer bytes.Buffer
	cmd.Stdout = &buffer

	log.Printf("Trimming %s to %s-%s. Saving to %s.\n", inputPath, start, duration, outputPath)

	err = cmd.Run()
	return
}
