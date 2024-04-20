package preinstall

import (
	"fmt"
	"os"

	"golang.org/x/term"

	sel "utils/ui/components/Select"
	spl "utils/ui/components/SplitScreen"
	exc "utils/ui/meta/Executor"

	"github.com/charmbracelet/lipgloss"
)

func PreInstall() {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))
	if e != nil {
		panic(e)
	}

	helpmenu := lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("8")).Render("[TAB]: Switch Focus\n[UP]: Move Cursor Up    [DOWN]: Move Cursor Down\n[SPACE]: Select Option  [ENTER]: Send Response")
	headertext := lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("13")).Render("[03/26]\tPre-Install\tSelect Window Manager")

	sw := (w / 2) - 4
	sh := h - 4 - lipgloss.Height(helpmenu) - lipgloss.Height(headertext)

	style := lipgloss.NewStyle().Height(sh).Width(sw).Align(lipgloss.Left, lipgloss.Top)

	sl := sel.Select(style, "Select Something", "i3WM", "BSPWM", "CAVA", "Kitty")
	m := exc.Executor(spl.SplitScreen(sl, sl, helpmenu, headertext))
	fmt.Println(sel.GetAnswer(spl.GetLeftSide(m)))
}
