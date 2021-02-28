package pkg

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func Ffprobe(path string) (response FfprobeResponse, err error) {
	cmd := exec.Command("ffprobe", "-i", path, "-print_format", "json", "-show_format")

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err = cmd.Run()
	if err != nil {
		return
	}
	err = json.Unmarshal(buffer.Bytes(), &response)
	return
}
