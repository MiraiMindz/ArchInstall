package splitscreen

import (
	"golang.org/x/term"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	leftSide  tea.Model
	rightSide tea.Model
	selected  uint
	helpmenu  string
	header    string
}

func GetLeftSide(m any) any  { return m.(model).leftSide }
func GetRightSide(m any) any { return m.(model).rightSide }

func SplitScreen(leftSide, rightSide tea.Model, helpmenu, header string) model {
	return model{
		leftSide:  leftSide,
		rightSide: rightSide,
		selected:  0,
		helpmenu:  helpmenu,
		header:    header,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "tab":
			if m.selected == 0 {
				m.selected = 1
			} else {
				m.selected = 0
			}
		}

		switch m.selected {
		case 0:
			m.leftSide, cmd = m.leftSide.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.rightSide, cmd = m.rightSide.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))
	if e != nil {
		panic(e)
	}

	sw := (w / 2) - 2
	sh := h - 2 - lipgloss.Height(m.helpmenu) - lipgloss.Height(m.header)

	unselectedStyle := lipgloss.NewStyle().Height(sh).Width(sw).BorderStyle(lipgloss.HiddenBorder()).PaddingLeft(2).PaddingRight(2).PaddingTop(1).PaddingBottom(1)
	selectedStyle := lipgloss.NewStyle().Height(sh).Width(sw).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("69")).PaddingLeft(2).PaddingRight(2).PaddingTop(1).PaddingBottom(1)

	// unselectedStyle := lipgloss.NewStyle().Height(sh).Width(sw).Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.HiddenBorder()).Padding(2)
	// selectedStyle := lipgloss.NewStyle().Height(sh).Width(sw).Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("69")).Padding(2)
	if m.selected == 1 {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			m.header,
			lipgloss.JoinHorizontal(lipgloss.Top, unselectedStyle.Render(m.leftSide.View()), selectedStyle.Render(m.rightSide.View())),
			m.helpmenu,
		)
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.header,
		lipgloss.JoinHorizontal(lipgloss.Top, selectedStyle.Render(m.leftSide.View()), unselectedStyle.Render(m.rightSide.View())),
		m.helpmenu,
	)
}
