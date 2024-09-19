package services

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

func DownloadVideo(url string, videoFormatId string, audioFormatId string) ([]byte, error) {
	uuidGen := uuid.New()
	cwd, _ := os.Getwd()
	exePath := cwd + "/bin/yt-dlp"
	outputPath := cwd + "/video-storage"
	outputFilename := "%(title)s [%(id)s]" + uuidGen.String() + ".%(ext)s"

	fmt.Println("Download Starting for: " + url)

	// First, get the expected filename
	filenameCmd := exec.Command(
		exePath,
		url,
		"-f",
		videoFormatId+"+"+audioFormatId,
		"--merge-output-format",
		"mp4",
		"-o",
		outputFilename,
		"--get-filename",
	)

	filename, err := filenameCmd.Output()
	if err != nil {
		return nil, err
	}

	filenameStr := string(filename) + uuidGen.String()
	fmt.Println("Filename: ", filenameStr)

	// Then, start the download
	cmd := exec.Command(
		exePath,
		url,
		"-f",
		videoFormatId+"+"+audioFormatId,
		"--merge-output-format",
		"mp4",
		"--paths",
		outputPath,
		"--output",
		filenameStr,
	)

	output, err := cmd.Output()

	return output, err
}
