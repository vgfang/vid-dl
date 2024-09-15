package services

import (
	"fmt"
	"os"
	"os/exec"
)

func GetFileOptions(url string) ([]byte, error) {
	cwd, _ := os.Getwd()
	cmd := exec.Command(cwd+"/bin/yt-dlp", url, "-F")
	output, err := cmd.Output()
	fmt.Println(output)

	// we want to parse the output of the file options cmd
	// and return a JSON with relevant audio/video settings

	return output, err
}
