package preinstall

import (
	"fmt"
	"os"

	"golang.org/x/term"

	header "utils/ui/components/Header"
	helpmenu "utils/ui/components/HelpMenu"
	sel "utils/ui/components/Select"
	spl "utils/ui/components/SplitScreen"
	ti "utils/ui/components/TextInput"
	exc "utils/ui/meta/Executor"

	"github.com/charmbracelet/lipgloss"
)

func PreInstall() {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))
	if e != nil {
		panic(e)
	}

	hpt, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Cursor Up    [DOWN]: Move Cursor Down\n[SPACE]: Select Option  [ENTER]: Send Response")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := ti.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		ti.TextInput("Enter Something:", "Type Here", true, 256, 32, lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top)),
		sel.Select(lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top), "Select Something", "i3WM", "BSPWM", "CAVA", "Kitty"),
		hpt,
		hdt,
	))))

	fmt.Println(r)
}
