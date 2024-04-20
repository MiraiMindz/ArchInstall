package Select

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	prompt string
	// bottom  string
	options []string
	cursor  int
	answer  string
	style   lipgloss.Style
}

var hoverStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
var selectedStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
var optStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))

func GetAnswer(sel any) string {
	return sel.(model).answer
}
func Select(style lipgloss.Style, prompt string, options ...string) model {
	return model{
		prompt: prompt,
		// bottom:  bottom,
		options: options,
		style:   style,
	}
}

func (m model) Init() tea.Cmd { return nil }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "enter":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case " ":
			m.answer = m.options[m.cursor]
		}
	}

	return m, nil
}
func (m model) View() string {
	s := fmt.Sprintf("%s\n", m.prompt)

	for i, opt := range m.options {
		checked := " "
		cursor := " "
		if m.answer == opt {
			checked = "x"
		}
		if m.cursor == i {
			cursor = ">"
		}
		if m.cursor == i {
			s += hoverStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, opt))
			s += "\n"
		} else {
			if m.answer == opt {
				s += selectedStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, opt))
				s += "\n"
			} else {
				s += optStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, opt))
				s += "\n"
			}
		}
	}

	s += "\n"
	// s += fmt.Sprintf("\n%s\n", m.bottom)
	return m.style.Render(s)
}
