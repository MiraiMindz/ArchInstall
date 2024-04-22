package textelement

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	text       string
	answer     string
	style      lipgloss.Style
	ismarkdown bool
}

func GetAnswer(mod any) string {
	return mod.(model).answer
}

func TextElement(ismarkdown bool, style lipgloss.Style, t string) model {
	return model{
		text:       t,
		style:      style,
		ismarkdown: ismarkdown,
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
	if m.ismarkdown {
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(lipgloss.Width(m.text)+2),
		)
		if err != nil {
			panic(err)
		}

		str, err := renderer.Render(m.text)
		if err != nil {
			panic(err)
		}

		return m.style.Render(str)
	}

	return m.style.Render(m.text)
}
