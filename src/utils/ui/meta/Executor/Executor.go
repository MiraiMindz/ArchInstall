package executor

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Executor(model any) any {
	p := tea.NewProgram(model.(tea.Model), tea.WithAltScreen())

	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return m
}
