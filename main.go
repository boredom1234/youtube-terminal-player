package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"youtube-terminal-player/ascii"
	"youtube-terminal-player/download"
	"youtube-terminal-player/frames"
	"youtube-terminal-player/player"
)

const (
	TempDir      = "tmp"
	VideoFile    = "tmp/video.mp4"
	FramePattern = "tmp/frame_%04d.png"
	FPS          = 60  // Increased frame rate
	Width        = 120 // Reduced resolution for better performance
	Height       = 67  // Maintains 1920x1080 aspect ratio
)

func main() {
	// Get YouTube URL from user input
	fmt.Print("Enter the Youtube Link: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	videoURL := scanner.Text()

	// Clear and recreate temp directory
	os.RemoveAll(TempDir)
	if err := os.MkdirAll(TempDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create temp directory: %v\n", err)
		return
	}

	// Step 1: Download video
	if err := download.DownloadVideo(videoURL, VideoFile); err != nil {
		fmt.Printf("Failed to download video: %v\n", err)
		return
	}

	// Step 2: Extract frames
	if err := frames.ExtractFrames(VideoFile, FramePattern, FPS, Width, Height); err != nil {
		fmt.Printf("Failed to extract frames: %v\n", err)
		return
	}

	// Step 3: Convert frames to ASCII
	frameFiles, _ := filepath.Glob("tmp/frame_*.png")
	asciiFrames := ascii.ProcessFrames(frameFiles, Width)

	// Step 4: Play frames in terminal
	player.PlayFrames(asciiFrames, FPS)
}
