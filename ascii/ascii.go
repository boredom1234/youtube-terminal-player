package ascii

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"
	"youtube-terminal-player/progress"

	"golang.org/x/image/draw"
)

const ASCIIChars = "@%#*+=-:. " // Characters for white areas

func ImageToASCII(img image.Image, width int) string {
	bounds := img.Bounds()
	aspectRatio := float64(bounds.Dy()) / float64(bounds.Dx())
	height := int(float64(width) * aspectRatio * 0.55) // Adjust for terminal font aspect ratio

	img = resize(img, width, height)
	grayImg := image.NewGray(img.Bounds())
	for y := 0; y < grayImg.Bounds().Dy(); y++ {
		for x := 0; x < grayImg.Bounds().Dx(); x++ {
			grayImg.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	asciiStr := ""
	for y := 0; y < grayImg.Bounds().Dy(); y++ {
		for x := 0; x < grayImg.Bounds().Dx(); x++ {
			pixel := grayImg.GrayAt(x, y).Y
			if pixel < 128 { // Black areas (low intensity)
				asciiStr += " " // Use space for black
			} else { // White areas (high intensity)
				asciiStr += string(ASCIIChars[pixel/51]) // Use ASCII characters for white
			}
		}
		asciiStr += "\n"
	}
	return asciiStr
}

func resize(img image.Image, width, height int) image.Image {
	resized := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)
	return resized
}

func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}
	return img, nil
}

func ProcessFrames(frameFiles []string, width int) []string {
	var wg sync.WaitGroup
	asciiFrames := make([]string, len(frameFiles))

	for i, frameFile := range frameFiles {
		wg.Add(1)
		go func(i int, frameFile string) {
			defer wg.Done()
			img, err := LoadImage(frameFile)
			if err != nil {
				fmt.Printf("Failed to load image: %v\n", err)
				return
			}
			asciiFrames[i] = ImageToASCII(img, width)
			progress.ShowProgress(i+1, len(frameFiles), "Converting to ASCII")
		}(i, frameFile)
	}

	wg.Wait()
	return asciiFrames
}
