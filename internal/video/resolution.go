package video

import (
	"os/exec"
	"strconv"
	"strings"
)


func getVideoResolution(videoURL string) (int, int, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=p=0", videoURL)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0,0, err
	}

	resolution := strings.TrimSpace(string(output))


	parts := strings.Split(resolution, ",")

	if len(parts) != 2 {
		return 0,0, nil
	}

	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0,0, err
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0,0, err
	}
	return width, height, nil
}
