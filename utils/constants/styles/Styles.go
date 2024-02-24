package utils

import (
	"strings"
	functions "utils/functions"

	"github.com/charmbracelet/lipgloss"
)

var DefaultTextStyle = lipgloss.NewStyle().Foreground(DefaultTextColor)
var HeaderTextStyle = lipgloss.NewStyle().Foreground(DefaultCyanColor)

func AlignTextWithStyle(alignment string, style lipgloss.Style) lipgloss.Style {
	w, _, _ := functions.GetTerminalSize()
	var position lipgloss.Position
	switch strings.ToLower(alignment) {
	case "center":
		position = lipgloss.Center
	case "left":
		position = lipgloss.Left
	case "right":
		position = lipgloss.Right
	default:
		position = lipgloss.Left
	}
	newStyle := style.Width(w).Align(position)
	return newStyle
}