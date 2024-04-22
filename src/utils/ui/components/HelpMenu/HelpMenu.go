package helpmenu

import (
	"github.com/charmbracelet/lipgloss"
)

func HelpMenu(content string) (string, int) {
	text := lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("8")).Render(content)
	size := lipgloss.Height(text)
	return text, size
}
