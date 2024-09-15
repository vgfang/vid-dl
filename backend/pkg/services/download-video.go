package services

import (
	"fmt"
	"os"
	"os/exec"
)

func DownloadVideo(url string, videoID string, audioID string) ([]byte, error) {
	cwd, _ := os.Getwd()
	exePath := cwd + "/bin/yt-dlp"
	outputPath := cwd + "/video-storage"

	cmd := exec.Command(
		exePath,
		url,
		"-f",
		videoID+"+"+audioID,
		"--merge-output-format",
		"mp4",
		"--paths",
		outputPath,
	)

	output, err := cmd.Output()
	fmt.Println(output)

	return output, err
}
