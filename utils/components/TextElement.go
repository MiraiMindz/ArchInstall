package utils

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type textElement struct {
	text string
	style lipgloss.Style
}

func TextElement(text string, style lipgloss.Style) textElement {
	return textElement{text: text, style: style}
}

func (t textElement) Init() tea.Cmd {
	return nil
}

func (t textElement) View() string {
	retVal := t.style.Render(t.text)
	return retVal
}

func (t textElement) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return t, tea.Quit
		case "q":
			return t, tea.Quit
		}
	}
	return t, nil
}
