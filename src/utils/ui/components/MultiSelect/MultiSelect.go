package multiselect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/viewport"
)

type dim struct {
	width  int
	height int
}

type model struct {
	prompt     string
	style      lipgloss.Style
	cursor     int
	choices    []string
	selected   map[int]string
	viewport   viewport.Model
	dimensions dim
	ready      bool
}

var hoverStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
var selectedStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
var optStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))

func GetAnswer(sel any) []string {
	var sa []string
	for _, v := range sel.(model).selected {
		sa = append(sa, v)
	}
	return sa
}

func MultiSelect(style lipgloss.Style, width, height int, prompt string, options ...string) model {
	return model{
		prompt:   prompt,
		style:    style,
		choices:  options,
		selected: make(map[int]string),
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = m.choices[m.cursor]
			}
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

	for i, opt := range m.choices {
		checked := " "
		cursor := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		if m.cursor == i {
			cursor = ">"
		}
		if m.cursor == i {
			s += hoverStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, opt))
			s += "\n"
		} else {
			if _, ok := m.selected[i]; ok {
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
