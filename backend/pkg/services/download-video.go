package services

import (
	"fmt"
	"os"
	"os/exec"
)

func DownloadVideo(url string, videoID string, audioID string) ([]byte, error) {
	cwd, _ := os.Getwd()
	cmd := exec.Command(
		cwd+"/bin/yt-dlp",
		url,
		"-f",
		videoID+"+"+audioID,
		"--merge-output-format",
		"mp4",
	)

	output, err := cmd.Output()
	fmt.Println(output)

	return output, err
}
