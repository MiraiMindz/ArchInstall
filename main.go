package main

import (
	//"fmt"
	components "utils/components"
	//constants "utils/constants"
	styles "utils/constants/styles"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// style := styles.AlignTextWithStyle("center", styles.HeaderTextStyle)
	lside := components.TextElement("LEFT SIDE", styles.DefaultTextStyle)
	rside := components.TextElement("RIGHT SIDE", styles.DefaultTextStyle)
	p := tea.NewProgram(components.SplitView(lside, rside), tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
