package textelement

import (
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	text   string
	answer string
	style  lipgloss.Style
}

func GetAnswer(mod any) string {
	return mod.(model).answer
}

func TextElement(style lipgloss.Style, t string) model {
	return model{
		text:  t,
		style: style,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc", "enter":
			m.answer = msg.String()
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.style.Render(m.text)
}
