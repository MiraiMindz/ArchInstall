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

type option struct {
	item        string
	description string
}

type model struct {
	prompt     string
	options    []option
	cursor     int
	answer     string
	style      lipgloss.Style
	viewport   viewport.Model
	dimensions dim
	ready      bool
}

var hoverStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Border(lipgloss.NormalBorder(), false, false, false, true).PaddingTop(1)
var selectedStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("12")).Border(lipgloss.NormalBorder(), false, false, false, true).PaddingLeft(2).PaddingTop(1)
var optStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).PaddingLeft(2)
var descStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).PaddingLeft(2)
var textSelectedStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("0"))

func Option(item, description string) option {
	return option{
		item:        item,
		description: description,
	}
}

func GetAnswer(sel any) string {
	return sel.(model).answer
}

func Select(style lipgloss.Style, width, height int, prompt string, options ...option) model {
	return model{
		prompt:  prompt,
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

func getDescriptionMaxHeight(options []option) (int, string) {
	maxCount := 0
	maxDescription := ""
	for _, opt := range options {
		count := lipgloss.Height(opt.description)
		if count > maxCount {
			maxCount = count
			maxDescription = opt.description
		}
	}
	return maxCount, maxDescription
}

func computeItemHeight(w int, options []option) (int, int) {
	var d, s, h, mdt string
	var mds, sz int

	mds, mdt = getDescriptionMaxHeight(options)
	h = textSelectedStyle.Width(w - 4).Render(fmt.Sprintf("[%s] %s", "x", "Test"))
	d = textSelectedStyle.Width(w - 4).Render(mdt)
	s += selectedStyle.Width(w - 4).Render(fmt.Sprintf("%s\n\n%s\n", h, d))
	s += "\n"

	if lipgloss.Height(s) == 0 {
		sz = lipgloss.Height(s) + 1
	} else {
		sz = lipgloss.Height(s)
	}

	return sz, mds
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
		op   []string
	)

	for _, v := range m.options {
		op = append(op, v.item)
	}

	cellH, _ := computeItemHeight(m.dimensions.width, m.options)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "enter":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				m.viewport.LineUp(cellH - 1)
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
				m.viewport.LineDown(cellH - 1)
			}
		case " ":
			m.answer = op[m.cursor]
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

	switch key := msg.(tea.KeyMsg).String(); key {
	case " ":
	case "j":
	case "k":
	case "up":
	case "down":
	default:
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.viewport.SetContent(m.renderList())

	return m, tea.Batch(cmds...)
}

func (m model) renderList() string {
	s := fmt.Sprintf("%s\n\n", m.prompt)
	var h, d string

	cellH, _ := computeItemHeight(m.dimensions.width, m.options)

	for i, opt := range m.options {
		checked := " "
		if m.answer == opt.item {
			checked = "x"
		}
		if m.cursor == i {
			h = optStyle.Width(m.dimensions.width - 4).Render(fmt.Sprintf("[%s] %s", checked, opt.item))
			d = descStyle.Width(m.dimensions.width - 4).Render(opt.description)

			s += hoverStyle.Height(cellH - 1).Width(m.dimensions.width - 4).Render(fmt.Sprintf("%s\n\n%s\n", h, d))
			s += "\n\n"
		} else {
			if m.answer == opt.item {
				h = textSelectedStyle.Width(m.dimensions.width - 4).Render(fmt.Sprintf("[%s] %s", checked, opt.item))
				d = textSelectedStyle.Width(m.dimensions.width - 4).Render(opt.description)

				s += selectedStyle.Height(cellH - 1).Width(m.dimensions.width - 4).Render(fmt.Sprintf("%s\n\n%s\n", h, d))
				s += "\n\n"
			} else {
				h = optStyle.Width(m.dimensions.width - 4).Render(fmt.Sprintf("[%s] %s", checked, opt.item))
				d = descStyle.Width(m.dimensions.width - 4).Render(opt.description)

				s += lipgloss.NewStyle().PaddingTop(1).Height(cellH - 1).Render(fmt.Sprintf("%s\n\n%s\n", h, d))
				s += "\n\n"
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
