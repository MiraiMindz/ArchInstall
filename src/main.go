package main

import (
	"fmt"
	"os"
	"strings"

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
		false,
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
		false,
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

func sectionC(w, h int) string {
	var (
		titleStyle = func() lipgloss.Style {
			b := lipgloss.NormalBorder()
			b.Right = "├"
			return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
		}()

		infoStyle = func() lipgloss.Style {
			b := lipgloss.NormalBorder()
			b.Left = "┤"
			return titleStyle.Copy().BorderStyle(b)
		}()
	)

	headerex := lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("13")).Render("This is an header example")
	helpmenuex := lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("8")).Render("[META] This is an example of help menu\n[TAB] The keys are specified between brackets\n[KEY] Example keymap")

	sw := (w / 2) - 2
	sh := h - 2 - lipgloss.Height(helpmenuex) - lipgloss.Height(headerex)

	header := lipgloss.JoinHorizontal(lipgloss.Center, titleStyle.Render("Title Example"), strings.Repeat("─", max(0, sw-(lipgloss.Width(titleStyle.Render("Title Example"))+4))))

	footer := lipgloss.JoinHorizontal(
		lipgloss.Center,
		strings.Repeat("─", max(0, sw-(lipgloss.Width(infoStyle.Render(fmt.Sprintf("%3.f%%", 98.32)))+4))),
		infoStyle.Render(fmt.Sprintf("%3.f%%", 98.32)))

	content := `This layout is not interactive.

This is an example of the Pager View using the <Viewport()> component, 
It can render both Markdown or Plain text and it's scrollable.
This side will be used to display Guides/Wikis/Informative content about 
the current step, so it's mainly used for reading things.

Because this is using the <Viewport()> component, it will not render it's contents 
at first glance, so you need to switch to this side, or trigger an event to render it.

Above It's the Pager Title the "Title Example" text, and below it's the position of the buffer
so, when it reachs 100% is at the end of the content and 0% at the beginning

You can switch between sides using the TAB key.

You can press enter to start the script, or press CTRL+C to exit.

`

	s := fmt.Sprintf("%s\n%s\n%s", header, content, footer)

	unselectedStyle := lipgloss.NewStyle().Height(sh).Width(sw).BorderStyle(lipgloss.HiddenBorder()).PaddingLeft(2).PaddingRight(2).PaddingTop(1).PaddingBottom(1)
	selectedStyle := lipgloss.NewStyle().Height(sh).Width(sw).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("69")).PaddingLeft(2).PaddingRight(2).PaddingTop(1).PaddingBottom(1)

	t := lipgloss.NewStyle().Render(`This layout is not interactive

The selected side is marked by the blue border.
This is the side where the input will be mostly handled, 
It's the "script" part of this installer, it will contain the selects, 
inputs, options, everything else that makes the script.

Above you will find the header, it marks the current step/section of the installer.
And below you will find the help menu, it's responsive, so, when you are focusing on the left side
it will show the keys for the left current component, and if you are focusing on the right side
it will show the keys for the right current component, it will mostly be the <Viewport()> component.

You can press enter to start the script, or press CTRL+C to exit.

`)

	f := lipgloss.JoinVertical(
		lipgloss.Left,
		headerex,
		lipgloss.JoinHorizontal(lipgloss.Top, selectedStyle.Render(t), unselectedStyle.Render(s)),
		helpmenuex,
	)

	r := te.GetAnswer(exc.Executor(te.TextElement(
		false,
		lipgloss.NewStyle(),
		f,
	)))

	if r == "" || r == "ctrl+c" {
		panic(fmt.Errorf("Wrong input %q\n", r))
	}

	return r

}

func main() {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))

	if e != nil {
		panic(e)
	}

	sectionA(w, h)
	sectionB(w, h)
	sectionC(w, h)
	pints.PreInstall()
}
