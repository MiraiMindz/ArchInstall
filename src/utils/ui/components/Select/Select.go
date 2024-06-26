package Select

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type dim struct {
	width  int
	height int
}

type model struct {
	prompt string
	// bottom  string
	options    []string
	cursor     int
	answer     string
	style      lipgloss.Style
	viewport   viewport.Model
	dimensions dim
	ready      bool
}

var hoverStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
var selectedStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
var optStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))

func GetAnswer(sel any) string {
	return sel.(model).answer
}

func Select(style lipgloss.Style, width, height int, prompt string, options ...string) model {
	return model{
		prompt: prompt,
		// bottom:  bottom,
		options: options,
		style:   style,
		dimensions: dim{
			width:  width,
			height: height,
		},
	}
}

func (m model) Init() tea.Cmd {
	m.viewport = viewport.New(m.dimensions.width-1, m.dimensions.height-1)
	m.viewport.HighPerformanceRendering = false
	m.viewport.SetContent(m.renderList())
	m.ready = true
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
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
	if !m.ready {
		m.viewport = viewport.New(m.dimensions.width-1, m.dimensions.height-1)
		m.viewport.HighPerformanceRendering = false
		m.viewport.SetContent(m.renderList())
		m.ready = true
	} else {
		m.viewport.Width = m.dimensions.width - 2
		m.viewport.Height = m.dimensions.height
	}

	if msg.(tea.KeyMsg).String() != " " {
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.viewport.SetContent(m.renderList())

	return m, tea.Batch(cmds...)
}

func (m model) renderList() string {
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
	return m.style.Render(s)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Press the arrow keys to load it's contents"
	}
	return m.viewport.View()
}
