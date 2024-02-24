package utils

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

