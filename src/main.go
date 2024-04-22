package main

import (
	"fmt"
	"os"
	//"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	exc "utils/ui/meta/Executor"

	pints "sections/preinstall"

	te "utils/ui/components/TextElement"
)

func sectionA(w, h int) string {
	viewportStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Height(h).Width(w)

	headerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("4"))

	subHeaderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Border(lipgloss.NormalBorder()).PaddingTop(1).PaddingBottom(1).PaddingLeft(3).PaddingRight(3)
	headerText := `
    _             _       _     _                    ___           _        _ _           
   / \   _ __ ___| |__   | |   (_)_ __  _   ___  __ |_ _|_ __  ___| |_ ____| | | ___ _ __ 
  / _ \ | '__/ __| '_ \  | |   | | '_ \| | | \ \/ /  | || '_ \/ __| __/ _  | | |/ _ \ '__|
 / ___ \| | | (__| | | | | |___| | | | | |_| |>  <   | || | | \__ \ || (_| | | |  __/ |   
/_/   \_\_|  \___|_| |_| |_____|_|_| |_|\__,_/_/\_\ |___|_| |_|___/\__\__,_|_|_|\___|_|   
                                                                                          
               __        __    _ _   _               _          ____       
               \ \      / / __(_) |_| |_ ___ _ __   (_)_ __    / ___| ___  
                \ \ /\ / / '__| | __| __/ _ \ '_ \  | | '_ \  | |  _ / _ \ 
                 \ V  V /| |  | | |_| ||  __/ | | | | | | | | | |_| | (_) | 
                  \_/\_/ |_|  |_|\__|\__\___|_| |_| |_|_| |_|  \____|\___/ 


`

	subHeaderText := `
Some quick notes about this script:
	It doesn't supports Wayland (It's a planned feature).
	As of ArchISO, the default keyboard layout is en-US, so, be careful when typing if you use other layout (I plan to add a keyboard selecter).
	This was made for me to deploy my system easily, so there is a lot of personal taste here.
	Yet it's a very broad script, so you might be able to use it for bootstrapping your own system/rice.
	I tried to automate/simplify everything that I could from the installation guide.
	It uses the standard 1024B measurement instead of 1000B (this means GiB instead of GB)
	This script is divided in 6 sections:
		1. Startup: Only sets some basic options like keyboard layout, HDD/SDD, etc...
		2. Pre-Install: Setup the drive and pacstrap for installation.
		3. Base-Install: Installs and configure the system, installs base packages and creates users.
		4. User Config: User customizations and custom (AUR/NIX) packages.
		5. Post-Install: Enables services, clear ups the installation.
		6. Rice: applies a rice on the system.

All of this being said. I would like to highlight that this was a fun project to work on.
Considers following me on GitHub @MiraiMindz.

`
	cmodel := te.TextElement(
		lipgloss.NewStyle(),
		viewportStyle.Render(lipgloss.JoinVertical(lipgloss.Center, headerStyle.Render(headerText), subHeaderStyle.Render(subHeaderText), lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Height(5).Width(w).Render("PRESS ENTER TO CONTINUE"))))

	r := te.GetAnswer(exc.Executor(cmodel))

	if r == "" || r == "ctrl+c" {
		panic(fmt.Errorf("Wrong input %q\n", r))
	}

	return r

}

func sectionB(w, h int) string {
	viewportStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Height(h).Width(w)
	r := te.GetAnswer(exc.Executor(te.TextElement(
		viewportStyle,
		lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Top).Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.NewStyle().Bold(true).Render("DISCLAIMER\n"),
				lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Render(fmt.Sprintf("%s\n\n%s\n", lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1, 2).Render(
					`This script uses the Charmbracelet's lipgloss, bubbletea, bubbles and glamour TUI libraries, so it's constrained 
to the limits of the libraries. Nonetheless, it's pretty responsive, and the only drawback is in the use of the 
<Viewport()> component, which requires a manual trigger to render it's contents.

Next I'll present you with the layout of the script and some explanations about it.`),
					lipgloss.NewStyle().Background(lipgloss.Color("15")).Foreground(lipgloss.Color("0")).Padding(1, 2).Bold(true).Render("Press ANY key to continue"))))),
	)))

	if r == "" || r == "ctrl+c" {
		panic(fmt.Errorf("Wrong input %q\n", r))
	}

	return r
}

// func sectionC(w, h int) string {
// 	var (
// 		titleStyle = func() lipgloss.Style {
// 			b := lipgloss.NormalBorder()
// 			b.Right = "├"
// 			return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
// 		}()
//
// 		infoStyle = func() lipgloss.Style {
// 			b := lipgloss.NormalBorder()
// 			b.Left = "┤"
// 			return titleStyle.Copy().BorderStyle(b)
// 		}()
// 	)
//
// 	header := lipgloss.JoinHorizontal(lipgloss.Center, titleStyle.Render("Title Example"), strings.Repeat("─", max(0, w-lipgloss.Width(titleStyle.Render("Title Example")))))
//
// 	footer := lipgloss.JoinHorizontal(
// 		lipgloss.Center,
// 		strings.Repeat("─", max(0, w-lipgloss.Width(infoStyle.Render(fmt.Sprintf("%3.f%%", 98.32))))),
// 		infoStyle.Render(fmt.Sprintf("%3.f%%", 98.32)))
//
// 	s := fmt.Sprintf("%s\n%s\n%s", header, content, footer)
//
// }

func main() {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))

	if e != nil {
		panic(e)
	}

	sectionA(w, h)
	sectionB(w, h)
	pints.PreInstall()
}
