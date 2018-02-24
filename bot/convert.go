package bot

import (
	"fmt"
	"os/exec"
)

func Convert(title, url string) error {
	fileName := fmt.Sprintf("files/%s.mp3", title)
	ffmpegArgs := []string{
		"-i", url,
		"-headers", "User-Agent: Go-http-client/1.1",
		"-codec:a", "libmp3lame", "-qscale:a", "2", fileName,
	}
	cmd := exec.Command("ffmpeg", ffmpegArgs...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
