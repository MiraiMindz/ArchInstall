package pager

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type dimensions struct {
	Width  int
	Height int
}

type model struct {
	content    string
	ready      bool
	viewport   viewport.Model
	dimensions dimensions
	title      string
}

const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.NormalBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

func Pager(ismarkdown bool, width, height int, title, content string) model {
	var c string

	if !ismarkdown {
		c = content
	} else {
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(width-2),
		)
		if err != nil {
			panic(err)
		}

		str, err := renderer.Render(content)
		if err != nil {
			panic(err)
		}
		c = str
	}

	return model{
		content: c,
		title:   title,
		dimensions: dimensions{
			Width:  width,
			Height: height,
		},
	}
}

func (m model) Init() tea.Cmd {
	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())
	verticalMarginHeight := headerHeight + footerHeight
	m.viewport = viewport.New(m.dimensions.Width-2, m.dimensions.Height-verticalMarginHeight)
	m.viewport.YPosition = headerHeight
	m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
	m.viewport.SetContent(m.content)
	m.viewport.YPosition = headerHeight
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
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "down", "j":
			m.viewport.LineDown(1)

		case "up", "k":
			m.viewport.LineUp(1)
		}

	}

	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())
	verticalMarginHeight := headerHeight + footerHeight

	if !m.ready {
		m.viewport = viewport.New(m.dimensions.Width-2, m.dimensions.Height-verticalMarginHeight)
		m.viewport.YPosition = headerHeight
		m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
		m.viewport.SetContent(m.content)
		m.viewport.YPosition = headerHeight
		m.ready = true
	} else {
		m.viewport.Width = m.dimensions.Width - 2
		m.viewport.Height = m.dimensions.Height - verticalMarginHeight
	}

	if useHighPerformanceRenderer {
		cmds = append(cmds, viewport.Sync(m.viewport))
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Press [TAB] or switch to this window to load it's contents"
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render(m.title)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
