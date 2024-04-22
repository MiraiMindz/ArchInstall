package textinput

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput textinput.Model
	prompt    string
	err       error
	style     lipgloss.Style
}

type errMsg error

func GetAnswer(ti any) string {
	return ti.(model).textInput.Value()
}

func TextInput(prompt, placeholder string, hidden bool, charlimit, width int, style lipgloss.Style) model {
	texinp := textinput.New()
	texinp.Placeholder = placeholder
	texinp.CharLimit = charlimit
	texinp.Width = width
	if hidden {
		texinp.EchoMode = textinput.EchoPassword
		texinp.EchoCharacter = '•'
	}
	texinp.Focus()

	return model{
		textInput: texinp,
		prompt:    prompt,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.style.Render(fmt.Sprintf(
		"%s\n\n%s\n\n",
		m.prompt,
		m.textInput.View(),
	))
}
