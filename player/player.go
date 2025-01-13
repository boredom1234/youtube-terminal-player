package player

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

const (
	ClearScreen = "\033[H\033[2J" // ANSI escape code to clear the screen
	HideCursor  = "\033[?25l"     // ANSI escape code to hide the cursor
	ShowCursor  = "\033[?25h"     // ANSI escape code to show the cursor
)

func PlayFrames(asciiFrames []string, fps int) {
	delay := time.Duration(1000/fps) * time.Millisecond
	fmt.Print(HideCursor) // Hide the cursor during playback

	// Start playing audio in background
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-an", "tmp/video.mp4")
		cmd.Run()
	}()

	// Display frames
	for _, frame := range asciiFrames {
		fmt.Print(ClearScreen) // Clear the screen
		fmt.Print(frame)       // Display the frame
		time.Sleep(delay)      // Wait for the next frame
	}

	fmt.Print(ShowCursor) // Show the cursor after playback
	wg.Wait()             // Wait for audio to finish
}
