package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func trim(inputPath, outputPath string, start, duration float32) (err error) {
	cmd := exec.Command("ffmpeg",
		"-ss", fmt.Sprintf("%f", start),
		"-i", inputPath,
		"-t", fmt.Sprintf("%f", duration),
		"-c", "copy",
		outputPath)

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err = cmd.Run()
	return
}
