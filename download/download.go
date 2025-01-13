package download

import (
	"fmt"
	"io"
	"os"
	"youtube-terminal-player/progress"

	"github.com/kkdai/youtube/v2"
)

func DownloadVideo(url string, outputFile string) error {
	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		return fmt.Errorf("failed to get video info: %v", err)
	}

	fmt.Printf("Downloading: %s\n", video.Title)
	stream, size, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		return fmt.Errorf("failed to get video stream: %v", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Track download progress
	progressReader := &ProgressReader{
		Reader:     stream,
		TotalBytes: int(size),
		ProgressFunc: func(current, total int) {
			progress.ShowProgress(current, total, "Downloading")
		},
	}

	_, err = io.Copy(file, progressReader)
	if err != nil {
		return fmt.Errorf("failed to save video: %v", err)
	}

	fmt.Println("Download complete!")
	return nil
}

type ProgressReader struct {
	Reader       io.Reader
	TotalBytes   int
	CurrentBytes int
	ProgressFunc func(current, total int)
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.CurrentBytes += n
	if pr.ProgressFunc != nil {
		pr.ProgressFunc(pr.CurrentBytes, pr.TotalBytes)
	}
	return n, err
}
