package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	exc "utils/ui/meta/Executor"

	pints "sections/preinstall"
)

type model struct {
	text   string
	answer string
}

func initModel(t string) model {
	return model{
		text: t,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc", "enter":
			m.answer = msg.String()
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.text
}

func main() {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))

	if e != nil {
		panic(e)
	}

	viewportStyle := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Height(h).Width(w)

	headerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("4"))

	subHeaderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Border(lipgloss.RoundedBorder()).PaddingTop(1).PaddingBottom(1).PaddingLeft(3).PaddingRight(3)
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
	cmodel := initModel(viewportStyle.Render(lipgloss.JoinVertical(lipgloss.Center), headerStyle.Render(headerText), subHeaderStyle.Render(subHeaderText), "\n\n\n\nPRESS ENTER TO CONTINUE"))

	result := exc.Executor(cmodel).(model).answer
	if result != "" {
		fmt.Println(result)
	}

	pints.PreInstall()
}
