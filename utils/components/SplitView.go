package utils

import (
	utils "utils/functions"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	leftState = iota
	rightState
)

type stateType uint
type SplitViewModel struct {
	leftModel, rightModel tea.Model
	// bottomMessage string
	state stateType
}

func SplitView(leftSide, rightSide tea.Model) SplitViewModel {
	return SplitViewModel{
		state: leftState,
		leftModel: leftSide,
		rightModel: rightSide,
		// bottomMessage: botMsg,
	}
}

func (m SplitViewModel) Init() tea.Cmd {
	return nil
}

func (m SplitViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == leftState {
				m.state = rightState
			} else {
				m.state = leftState
			}
		}
		switch m.state {
		// update whichever model is focused
		case leftState:
			m.leftModel, cmd = m.leftModel.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.rightModel, cmd = m.rightModel.Update(msg)
			cmds = append(cmds, cmd)
		}
	default:
		if m.state == leftState {
			m.leftModel, cmd = m.leftModel.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			m.rightModel, cmd = m.rightModel.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m SplitViewModel) View() string {
	var s string
	w, h, e := utils.GetTerminalSize()
	if e != nil {
		panic(e)
	}
	// bottomMessageLines := utils.CountLines(m.bottomMessage, w) + 1
	// viewHeight := h - bottomMessageLines
	viewHeight := h
	lwidth, rwidth := utils.DivideTerminalWidth(w)
	leftStyle := lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderRight(true).Width(lwidth).Height(viewHeight).PaddingLeft(1).PaddingRight(1)
	rightStyle := lipgloss.NewStyle().Width(rwidth).Height(viewHeight).PaddingLeft(1).PaddingRight(1)
	focusedLeft := leftStyle.Copy().Foreground(lipgloss.Color("15")).Background(lipgloss.Color("8"))
	focusedRight := rightStyle.Copy().Foreground(lipgloss.Color("15")).Background(lipgloss.Color("8"))
	unfocusedLeft := leftStyle.Copy().Foreground(lipgloss.Color("8")).Background(lipgloss.Color("0"))
	unfocusedRight := rightStyle.Copy().Foreground(lipgloss.Color("8")).Background(lipgloss.Color("0"))


	if m.state == leftState {
		s += lipgloss.JoinHorizontal(lipgloss.Top, focusedLeft.Render(m.leftModel.View()), unfocusedRight.Render(m.rightModel.View()))
	} else {
		s += lipgloss.JoinHorizontal(lipgloss.Top, unfocusedLeft.Render(m.leftModel.View()), focusedRight.Render(m.rightModel.View()))
	}

	//views = append(views, leftStyle.Render(m.leftModel.View()))
	//views = append(views, rightStyle.Render(m.rightModel.View()))


	return s
	// return lipgloss.JoinHorizontal(lipgloss.Top, views...) + "\n\n" + m.bottomMessage
	
}