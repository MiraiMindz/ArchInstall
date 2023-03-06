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
)

// HELPER FUNCTIONS

func welcomeScreenPrint() {

	helpers.PrintAsciiArt()

	fmt.Println(helpers.PrintYellow("Some quick notes about this script:"))
	fmt.Println(helpers.PrintBlack("\tIt doesn't supports Wayland."))
	fmt.Println(helpers.PrintBlack("\tAs of ArchISO, the default keyboard layout is en-US, so, be careful when typing if you use other layout."))
	fmt.Println(helpers.PrintBlack("\tThis was made for me to deploy my system easily, so there is a lot of personal taste here."))
	fmt.Println(helpers.PrintBlack("\tYet it's a very broad script, so you might be able to use it for bootstrapping your own system/rice."))
	fmt.Println(helpers.PrintBlack("\tI tried to automate/simplify everything that I could from the installation guide."))
	fmt.Println(helpers.PrintBlack("\tIt uses the standard 1024B measurement instead of 1000B (this means GiB instead of GB)"))
	fmt.Println(helpers.PrintBlack("\tThis script is divided in 5 sections:"))
	fmt.Println(helpers.PrintBlack("\t\t1. Startup: Only sets some basic options like keyboard layout, HDD/SDD, etc..."))
	fmt.Println(helpers.PrintBlack("\t\t2. Pre-Install: Setup the drive and pacstrap for installation."))
	fmt.Println(helpers.PrintBlack("\t\t3. Base-Install: Installs and configure the system, installs base packages and creates users."))
	fmt.Println(helpers.PrintBlack("\t\t4. User Config: User customizations and custom (AUR/NIX) packages."))
	fmt.Println(helpers.PrintBlack("\t\t5. Post-Install: Enables services, clear ups the installation."))

	fmt.Println("All of this being said. I would like to highlight that this was a fun project to work on.")
	fmt.Print("Considers following me on GitHub @MiraiMindz.\n\n")
}

// MAIN FUNC
func main() {
	welcomeScreenPrint()

	//z := "300M"
	//helpers.ParseSizeString(z)
	if helpers.YesNo("Can we proceed to the installation?") {
		sections.Startupp()
		sections.PreInstall()
		//sections.BaseInstall()
		//sections.UserConfig()
		//sections.PostInstall()
	}
}
