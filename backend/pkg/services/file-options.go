package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Format struct {
	ID         string
	Ext        string
	Resolution string
	Filesize   string
	TBR        string
	Proto      string
}

func GetFileOptions(url string) ([]Format, error) {
	cwd, _ := os.Getwd()
	cmd := exec.Command(cwd+"/bin/yt-dlp", url, "-F")
	output, err := cmd.Output()

	// we want to parse the output of the file options cmd
	// and return a JSON with relevant audio/video settings
	var formats []Format
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "---") || strings.HasPrefix(line, "ID") || strings.HasPrefix(line, "[") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		formatted := Format{
			ID:         fields[0],
			Ext:        fields[1],
			Resolution: fields[2],
			Filesize:   fields[3],
			TBR:        fields[4],
			Proto:      fields[5],
		}

		formats = append(formats, formatted)
	}

	for _, form := range formats {
		fmt.Println(form, form.Ext)
	}

	return formats, err
}
