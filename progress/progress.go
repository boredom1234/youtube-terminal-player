package progress

import (
	"fmt"
	"strings"
)

func ShowProgress(current, total int, message string) {
	width := 50
	percent := float64(current) / float64(total)
	filled := int(float64(width) * percent)

	// Use more visually appealing characters
	leftBracket := "【"
	rightBracket := "】"
	filledChar := "█"
	emptyChar := "░"

	// Build the progress bar
	bar := strings.Repeat(filledChar, filled) + strings.Repeat(emptyChar, width-filled)

	// Add color and style
	coloredMessage := fmt.Sprintf("\033[1;36m%s\033[0m", message)         // Cyan, bold
	coloredBar := fmt.Sprintf("\033[1;32m%s\033[0m", bar)                 // Green, bold
	coloredPercent := fmt.Sprintf("\033[1;33m%.1f%%\033[0m", percent*100) // Yellow, bold

	// Print the formatted progress bar
	fmt.Printf("\r%s %s%s%s %s",
		coloredMessage,
		leftBracket,
		coloredBar,
		rightBracket,
		coloredPercent,
	)

	if current == total {
		fmt.Println() // Move to the next line after completion
	}
}
