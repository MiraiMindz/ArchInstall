package utils

import "strings"

func CountLines(s string, width int) int {
	// Split the string into words
	words := strings.Fields(s)

	// Initialize variables to keep track of current line width and line count
	currentWidth := 0
	lineCount := 1

	// Iterate over each word
	for _, word := range words {
		// Calculate the width of the current word including spaces
		wordWidth := len(word) + 1 // Add 1 for the space

		// If adding the current word width exceeds the specified width, start a new line
		if currentWidth+wordWidth > width {
			lineCount++
			currentWidth = 0
		}

		// Add the current word width to the current line width
		currentWidth += wordWidth
	}

	// Return the total line count
	return lineCount
}
