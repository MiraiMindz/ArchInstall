package header

import "github.com/charmbracelet/lipgloss"

func Header(content string) (string, int) {
	text := lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("13")).Render(content)
	size := lipgloss.Height(text)
	return text, size
}
