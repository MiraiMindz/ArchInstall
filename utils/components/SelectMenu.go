package utils

// import (
// 	"fmt"
// 	"strings"

// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var SelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
// var UnselectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))

// type ItemElement struct {
// 	Text    string
// 	OnPress func() tea.Msg
// }

// type MenuElement struct {
// 	Options       []ItemElement
// 	SelectedIndex int
// }

// type SelectedOptionMsg struct{}

// func SelectMenu(options []ItemElement) MenuElement {
// 	return MenuElement{
// 		Options: options,
// 	}
// }

// func (m MenuElement) Init() tea.Cmd {
// 	return nil
// }

// func (m MenuElement) View() string {
// 	var options []string
// 	for i, o := range m.Options {
// 		if i == m.SelectedIndex {
// 			options = append(options, fmt.Sprintf("> %s", o.Text))
// 		} else {
// 			options = append(options, fmt.Sprintf("  %s", o.Text))
// 		}
// 	}

// 	retVal := "Press Enter/Return to select a value, use the arrow keys to move.\n"
// 	retVal += strings.Join(options, "\n")
// 	retVal += "\nPress Ctrl+C or q to exit.\n\n"
// 	return retVal
// }

// func (m MenuElement) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case SelectedOptionMsg:
// 		return m.toggleSelectedItem(), nil
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		case "down", "right", "up", "left":
// 			return m.moveCursor(msg), nil
// 		case "enter", "return":
// 			return m, m.Options[m.SelectedIndex].OnPress
// 		}
// 	}
// 	return m, nil
// }

// func (m MenuElement) moveCursor(msg tea.KeyMsg) MenuElement {
// 	switch msg.String() {
// 	case "up", "left":
// 		m.SelectedIndex--
// 	case "down", "right":
// 		m.SelectedIndex++
// 	default:
// 		// do nothing
// 	}

// 	optCount := len(m.Options)
// 	m.SelectedIndex = (m.SelectedIndex + optCount) % optCount
// 	return m
// }

// func (m MenuElement) toggleSelectedItem() tea.Model {
// 	selectedText := m.Options[m.SelectedIndex].Text
// 	if strings.Contains(selectedText, "[x]") {
// 		m.Options[m.SelectedIndex].Text = UnselectedStyle.Render(strings.Replace(selectedText, "[x]", "[ ]", -1))
// 	} else {
// 		m.Options[m.SelectedIndex].Text = SelectedStyle.Render(strings.Replace(selectedText, "[ ]", "[x]", -1))
// 	}
// 	return m
// }

///////////////////
///////////////////

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)



var docStyle = lipgloss.NewStyle().Margin(1, 2)

type MenuItem struct {
	TitleField, Desc string
}

func (i MenuItem) Title() string {
	return i.TitleField
}
func (i MenuItem) Description() string {
	return i.Desc
}
func (i MenuItem) FilterValue() string {
	return i.TitleField
}

type MenuElement struct {
	List   list.Model
	Choice MenuItem
	SubmitFunction func(args ...interface{}) []interface{}
}

func (m MenuElement) Init() tea.Cmd {
	return nil
}

func (m MenuElement) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			selectedItem := m.List.SelectedItem()
			m.Choice = selectedItem.(MenuItem)
			_ = m.SubmitFunction(m)
			return m, tea.Quit

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m MenuElement) View() string {
	return docStyle.Render(m.List.View())
}

func SelectMenu(title string, submitFunction func(args ...interface{}) []interface{}, items []list.Item) MenuElement {
	m := MenuElement{List: list.New(items, list.NewDefaultDelegate(), 0, 0), SubmitFunction: submitFunction}
	m.List.Title = title
	return m
}
