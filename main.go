/*******************************************************************************
Notes:
Create a common config file to store common variables for the script files.
Compile the files separated and uses a shell script to launch then accordingly.

Planned Features:
Use the i18n and i10n pkgs to implement internationalization.
Implements a verbose/explanatory script log like the NIX installer.

THINGS THAT I *NEED* TO DO:
Implement check/search criteria in packages to prevent the user to install
packages that are not available in the Arch Linux main repos.

Split the installation in multiple sections:
	[ ] - Startup
	[ ] - Pre-Install
	[ ] - Base-Install
	[ ] - User Configuration
	[ ] - Post-Install
	[ ] - Rice System
Startup Steps:
	[ ] - Create Config File
	[ ] - Get User Password
	[ ] - Get TimeZone
	[ ] - Get Keyboard Layout
	[ ] - Get Disk Device
	[ ] - Get Disk Size
	[ ] - Set Mount Options
	[ ] - Get User Name
	[ ] - Get Host Name
	[ ] - Get AUR Helper or NIX Package Manager
Pre-Install Steps:
	[ ] - Get Country ISO
	[ ] - Synchronize Hardware Clock
	[ ] - Set Pacman
	[ ] - Download Essential Packages
	[ ] - Format Disk
	[ ] - Get BOOT System
	[ ] - Get SWAP Partition
	[ ] - Partition Disk
	[ ] - Mount Partitions
	[ ] - Install Base System
	[ ] - Generate File System Table (FSTAB)
	[ ] - Install GRUB
Base-Install Steps:
	[ ] - Set Network
	[ ] - Config Pacman (again)
	[ ] - Set MAKEPKG Config
	[ ] - Set Language
	[ ] - Set Locale
	[ ] - Set Time Zone
	[ ] - Set NOPASSWD SUDO (temporary)
	[ ] - Config Pacman (again)
	[ ] - Installing Setup Packages
	[ ] - Install Microcode
	[ ] - Install Graphics Cards (NVIDIA|AMD|INTEGRATED|INTEL)
	[ ] - Set User
	[ ] - Set Hosts file
User Configuration Steps:
	[ ] - Creates .cache
	[ ] - Sets SHELL
	[ ] - Installs AUR/NIX
	[ ] - Install Packages
Post-Install Steps:
	[ ] - Customizes GRUB
	[ ] - Installs Display Manager
	[ ] - Customizes Display Manager
	[ ] - Enable Services
	[ ] - Customizing Plymounth
	[ ] - Remove NOPASSWD SUDO
	[ ] - Set Default SUDO
Rice System Steps:
	[ ] - import and install user dotfiles (version systems like git/svn)
	[ ] - import and install user files from hard-drive
	[ ] -
	[ ] -
*******************************************************************************/

package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	utils "utils/components"

// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// )

// // Automates the process of running shell commands
// func RunShellCommand(testMode, returnOutput bool, name string, args ...string) string {
// 	if testMode {
// 		x := []string{name}
// 		x = append(x, args...)
// 		fmt.Println(x)
// 		return ""
// 	} else {
// 		cmd := exec.Command(name, args...)
// 		cmd.Stdin = os.Stdin
// 		if returnOutput {
// 			out, err := cmd.Output()
// 			Check(err)
// 			return string(out)
// 		} else {
// 			cmd.Stdout = os.Stdout
// 			cmd.Run()
// 			return ""
// 		}
// 	}
// }

// // Checks a error and panics
// func Check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func main() {
// 	itemsCount := 64
// 	opts := make([]list.Item, itemsCount)
// 	text := "item"

// 	for i := 0; i < itemsCount; i++ {
// 		item := utils.MenuItem{
// 			TitleField: fmt.Sprintf("%s %d", text, i),
// 			Desc:       fmt.Sprintf("desc %s %d", text, i),
// 		}
// 		opts[i] = item
// 	}

// 	submitFunction := func(args ...interface{}) []interface{} {
// 		model := args[0].(utils.MenuElement)
// 		RunShellCommand(false, false, "touch", model.Choice.TitleField)
// 		return nil
// 	}

// 	p := tea.NewProgram(utils.SelectMenu("Test", submitFunction, opts), tea.WithAltScreen())
// 	m, err := p.Run()
// 	if err != nil {
// 		panic(err)
// 	}

// 	nullItem := utils.MenuItem{TitleField: "", Desc: ""}
// 	if m, ok := m.(utils.MenuElement); ok && m.Choice != nullItem {
// 		fmt.Printf("\n---\nYou chose %s!\n", m.Choice.TitleField)
// 	}
// }





////////////////////////////////////////
////////////////////////////////////////

import (
	components "utils/components"
	//constants "utils/constants"
	styles "utils/constants/styles"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(components.TextElement("Hello World!", styles.DefaultTextStyle), tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}