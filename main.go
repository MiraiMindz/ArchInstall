package main

/* Notes
Create a common config file (JSON) to store common variables for the script files.
Compile the files separated and uses a shell script to launch then accordingly.

Split the installation in multiple sections:
	[ ] - Pre-Setup (Sets the variables, the files, etc..)
	[ ] - Setup (Pre Chroot stuff)
	[ ] - Base Install
	[ ] - User Configuration
	[ ] - Post Install
	[ ] - Rice System
Pre-Setup Steps:
	[ ] -
	[ ] -
	[ ] -
	[ ] -
Setup Steps:
	[ ] - Set the HDD
	[ ] - Load keys
	[ ] - Update system clock
	[ ] - Generate FSTab
	[ ] - Check FSTab
	[ ] - Install base packages
	[ ] - Chroot
Base Install Steps:
	[ ] - Update ZoneInfo
	[ ] - Sync hardware clock
	[ ] - Generate locales
	[ ] - Save Locales
	[ ] - Save keyboard layout
	[ ] - Set hostname
	[ ] - Hosts file and custom hosts
	[ ] - Root password
	[ ] - Processor micro-code
	[ ] - Download bootloader and more
	[ ] - Install bootloader
	[ ] - Create Dotfiles Install Script
User Configuration Steps:
	[ ] -
	[ ] -
	[ ] -
	[ ] -
Post Install Steps:
	[ ] -
	[ ] -
	[ ] -
	[ ] -
Rice System Steps:
	[ ] -
	[ ] -
	[ ] -
	[ ] -
*/

// IMPORTS
import (
	"ArchInstall/helpers"
	"ArchInstall/sections"
	"fmt"
	"github.com/fatih/color"
)

// HELPER FUNCTIONS

func welcomeScreenPrint() {
	cyan := color.New(color.FgCyan).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	black := color.New(color.FgBlack).SprintFunc()

	//+ This is a big mess but it looks beautiful in the terminal
	coloredArchLinux1 := blue("    _             _       _     _                  ")
	coloredArchLinux2 := blue("   / \\   _ __ ___| |__   | |   (_)_ __  _   ___  __")
	coloredArchLinux3 := blue("  / _ \\ | '__/ __| '_ \\  | |   | | '_ \\| | | \\ \\/ /")
	coloredArchLinux4 := blue(" / ___ \\| | | (__| | | | | |___| | | | | |_| |>  < ")
	coloredArchLinux5 := blue("/_/   \\_\\_|  \\___|_| |_| |_____|_|_| |_|\\__,_/_/\\_\\")
	ghostWrittenInGo1 := "__        __    _ _   _               _          ____       "
	ghostWrittenInGo2 := "\\ \\      / / __(_) |_| |_ ___ _ __   (_)_ __    / ___| ___  "
	ghostWrittenInGo3 := " \\ \\ /\\ / / '__| | __| __/ _ \\ '_ \\  | | '_ \\  | |  _ / _ \\ "
	ghostWrittenInGo4 := "  \\ V  V /| |  | | |_| ||  __/ | | | | | | | | | |_| | (_) |"
	ghostWrittenInGo5 := "   \\_/\\_/ |_|  |_|\\__|\\__\\___|_| |_| |_|_| |_|  \\____|\\___/ "
	coloredWrittenInGo1 := fmt.Sprintf("__        __    _ _   _               _       %s", cyan("   ____       "))
	coloredWrittenInGo2 := fmt.Sprintf("\\ \\      / / __(_) |_| |_ ___ _ __   (_)_ __  %s", cyan("  / ___| ___  "))
	coloredWrittenInGo3 := fmt.Sprintf(" \\ \\ /\\ / / '__| | __| __/ _ \\ '_ \\  | | '_ \\ %s", cyan(" | |  _ / _ \\ "))
	coloredWrittenInGo4 := fmt.Sprintf("  \\ V  V /| |  | | |_| ||  __/ | | | | | | | |%s", cyan(" | |_| | (_) |"))
	coloredWrittenInGo5 := fmt.Sprintf("   \\_/\\_/ |_|  |_|\\__|\\__\\___|_| |_| |_|_| |_|%s", cyan("  \\____|\\___/ "))

	ghostArchlinuxInstaller1 := "    _             _       _     _                    ___           _        _ _           "
	ghostArchlinuxInstaller2 := "   / \\   _ __ ___| |__   | |   (_)_ __  _   ___  __ |_ _|_ __  ___| |_ __ _| | | ___ _ __ "
	ghostArchlinuxInstaller3 := "  / _ \\ | '__/ __| '_ \\  | |   | | '_ \\| | | \\ \\/ /  | || '_ \\/ __| __/ _` | | |/ _ \\ '__|"
	ghostArchlinuxInstaller4 := " / ___ \\| | | (__| | | | | |___| | | | | |_| |>  <   | || | | \\__ \\ || (_| | | |  __/ |   "
	ghostArchlinuxInstaller5 := "/_/   \\_\\_|  \\___|_| |_| |_____|_|_| |_|\\__,_/_/\\_\\ |___|_| |_|___/\\__\\__,_|_|_|\\___|_|   "
	coloredArchLinuxInstaller1 := fmt.Sprintf("%s  ___           _        _ _           ", coloredArchLinux1)
	coloredArchLinuxInstaller2 := fmt.Sprintf("%s |_ _|_ __  ___| |_ __ _| | | ___ _ __ ", coloredArchLinux2)
	coloredArchLinuxInstaller3 := fmt.Sprintf("%s  | || '_ \\/ __| __/ _` | | |/ _ \\ '__|", coloredArchLinux3)
	coloredArchLinuxInstaller4 := fmt.Sprintf("%s  | || | | \\__ \\ || (_| | | |  __/ |   ", coloredArchLinux4)
	coloredArchLinuxInstaller5 := fmt.Sprintf("%s |___|_| |_|___/\\__\\__,_|_|_|\\___|_|   ", coloredArchLinux5)
	fmt.Println(helpers.CenterSprint(coloredArchLinuxInstaller1, ghostArchlinuxInstaller1))
	fmt.Println(helpers.CenterSprint(coloredArchLinuxInstaller2, ghostArchlinuxInstaller2))
	fmt.Println(helpers.CenterSprint(coloredArchLinuxInstaller3, ghostArchlinuxInstaller3))
	fmt.Println(helpers.CenterSprint(coloredArchLinuxInstaller4, ghostArchlinuxInstaller4))
	fmt.Println(helpers.CenterSprint(coloredArchLinuxInstaller5, ghostArchlinuxInstaller5))

	fmt.Println(helpers.CenterSprint(coloredWrittenInGo1, ghostWrittenInGo1))
	fmt.Println(helpers.CenterSprint(coloredWrittenInGo2, ghostWrittenInGo2))
	fmt.Println(helpers.CenterSprint(coloredWrittenInGo3, ghostWrittenInGo3))
	fmt.Println(helpers.CenterSprint(coloredWrittenInGo4, ghostWrittenInGo4))
	fmt.Println(helpers.CenterSprint(coloredWrittenInGo5, ghostWrittenInGo5))

	fmt.Println(yellow("Some quick notes about this script:"))
	fmt.Println(black("\tI doesn't supports wayland."))
	fmt.Println(black("\tAs of ArchISO, the default keyboard layout is en-US, so, be careful when typing if you use other layout."))
	fmt.Println(black("\tThis was made for me to deploy my sistem easily, so there is a lot of personal taste here."))
	fmt.Println(black("\tYet it's a very broad script, so you might be able to use it for bootstraping your own system/rice."))
	fmt.Println(black("\tI tried to automate/simplify everything that I could from the installation guide."))
	fmt.Println(black("\tThis script is divided in 5 sections:"))
	fmt.Println(black("\t\t1. Startup: Only sets some basic options like keyboard layout, HDD/SDD, etc..."))
	fmt.Println(black("\t\t2. Pre-Install: Setup the drive and pacstrap for installation."))
	fmt.Println(black("\t\t3. Base-Install: Installs and configure the system, installs base packages and creates users."))
	fmt.Println(black("\t\t4. User Config: User customizations and custom (AUR/NIX) packages."))
	fmt.Println(black("\t\t5. Post-Install: Enables services, clear ups the installation."))

	fmt.Println("All of this I would only highlight that this was a fun project to work on.")
	fmt.Print("Considers following me on GitHub @MiraiMindz.\n\n")
}

// MAIN FUNC
func main() {
	welcomeScreenPrint()
	if helpers.YesNo("Can we proceed to the installation?") {
		sections.Startupp()
		sections.PreInstall()
		sections.BaseInstall()
		sections.UserConfig()
		sections.PostInstall()
	}
}
