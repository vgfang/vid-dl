package services

import (
	"fmt"
	"os/exec"
)

func GetFileOptions(url string) ([]byte, error) {
	cmd := exec.Command("../bin/yt-dlp", url, "-F")
	output, err := cmd.Output()
	fmt.Println(output)

	// we want to parse the output of the file options cmd
	// and return a JSON with relevant audio/video settings

	return output, err
}
