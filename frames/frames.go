package frames

import (
	"fmt"
	"os/exec"
)

func ExtractFrames(videoFile string, outputPattern string, fps int, width int, height int) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", videoFile,
		"-vf", fmt.Sprintf("fps=%d,scale=%d:%d", fps, width, height),
		outputPattern,
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to extract frames: %v", err)
	}
	return nil
}
